/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package imds

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/bwagner5/imds-client/pkg/docs"
	"github.com/olekukonko/tablewriter"
	"github.com/samber/lo"
)

const (
	spotITNPath     = "spot/termination-time"
	scheduledEvents = "events/maintenance/scheduled"
)

type Client struct {
	*imds.Client
}

type ScheduledEventDetail struct {
	NotBefore   string `json:"NotBefore"`
	Code        string `json:"Code"`
	Description string `json:"Description"`
	EventID     string `json:"EventId"`
	NotAfter    string `json:"NotAfter"`
	State       string `json:"State"`
}

type InstanceAction struct {
	Action string `json:"action"`
	Time   string `json:"time"`
}

type RebalanceRecommendation struct {
	NoticeTime string `json:"noticeTime"`
}

func NewClient(ctx context.Context, endpoint string) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, withIMDSEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: imds.NewFromConfig(cfg),
	}, nil
}

func withIMDSEndpoint(imdsEndpoint string) func(*config.LoadOptions) error {
	return func(lo *config.LoadOptions) error {
		lo.EC2IMDSEndpoint = imdsEndpoint
		if net.ParseIP(imdsEndpoint).To4() == nil {
			lo.EC2IMDSEndpointMode = imds.EndpointModeStateIPv6
		} else {
			lo.EC2IMDSEndpointMode = imds.EndpointModeStateIPv4
		}
		return nil
	}
}

type lookup struct {
	path     string
	terminal bool
}

func (i Client) WatchRecurse(ctx context.Context, startingPath string) <-chan map[string]any {
	outChan := make(chan map[string]any, 10)
	go func(ctx context.Context, outC chan map[string]any) {
		var prev map[string]any
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		initial := true
		watchFn := func() {
			if out := i.GetRecurse(ctx, startingPath); !reflect.DeepEqual(prev, out) {
				select {
				case outC <- out:
					prev = out
				default: // full channel, so take one off and put on the latest
					<-outC
					outC <- out
					prev = out
				}
			}
		}
		for {
			if initial {
				initial = false
				watchFn()
			}
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				watchFn()
			}
		}
	}(ctx, outChan)
	return outChan
}

func (i Client) GetRecurse(ctx context.Context, startingPath string) map[string]any {
	errs := map[string]bool{}
	startingPath = strings.Trim(startingPath, "/")
	var paths []lookup
	all := map[string]any{}
	if startingPath == "" {
		paths = append(paths, lookup{path: "dynamic"}, lookup{path: "meta-data"}, lookup{path: "user-data"})
	} else {
		paths = append(paths, lookup{path: startingPath})
	}
	for len(paths) > 0 {
		p := paths[0]
		paths = paths[1:]
		resp, err := i.Get(ctx, p.path)
		if err != nil && !errs[p.path] {
			// if the path cannot be found, then we probably didn't identify that it was terminal, so add back to the stack
			// for re-evaluation as terminal
			tokens := strings.Split(p.path, "/")
			paths = append(paths, lookup{path: strings.Join(tokens[0:len(tokens)-1], "/"), terminal: true})
			errs[p.path] = true
			continue
		} else if err != nil {
			continue
		} else if _, ok := asJSON(resp); ok {
			p.terminal = true
		}

		if strings.HasPrefix(p.path, "user-data") {
			all["user-data"] = string(resp)
		} else if p.terminal {
			m := mapAt(all, p)
			lastToken := lastToken(p.path)
			if lastToken == "pkcs7" || lastToken == "signature" || lastToken == "rsa2048" {
				m[lastToken] = string(resp)
			} else if jsMap, ok := asJSON(resp); ok {
				m[lastToken] = jsMap
			} else if lines := strings.Split(string(resp), "\n"); len(lines) > 1 || strings.HasSuffix(lastToken, "s") {
				m[lastToken] = lines
			} else {
				m[lastToken] = string(resp)
			}
		} else {
			paths = append(paths, getLookups(p.path, resp)...)
		}
	}
	return all
}

func (i Client) Get(ctx context.Context, path string) ([]byte, error) {
	path = strings.Trim(path, "/")
	switch {
	case strings.HasPrefix(path, "dynamic"):
		out, err := i.Client.GetDynamicData(ctx, &imds.GetDynamicDataInput{Path: strings.Replace(path, "dynamic", "", 1)})
		if err != nil {
			return nil, err
		}
		resp, err := io.ReadAll(out.Content)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case strings.HasPrefix(path, "meta-data"):
		out, err := i.Client.GetMetadata(ctx, &imds.GetMetadataInput{Path: strings.Replace(path, "meta-data", "", 1)})
		if err != nil {
			return nil, err
		}
		resp, err := io.ReadAll(out.Content)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case strings.HasPrefix(path, "user-data"):
		out, err := i.Client.GetUserData(ctx, &imds.GetUserDataInput{})
		if err != nil {
			return nil, err
		}
		resp, err := io.ReadAll(out.Content)
		if err != nil {
			return nil, err
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("unsupported path")
	}
}

func AllDocs() map[string]any {
	return lo.Assign(UserdataDocs(), DynamicDocs(), MetadataDocs())
}

func UserdataDocs() map[string]any {
	return map[string]any{"user-data": ""}
}

func MetadataDocs() map[string]any {
	nestedMap := map[string]any{}
	for _, entry := range docs.InstanceMetadataCategoryEntries {
		curr := mapAt(nestedMap, lookup{path: entry.Category, terminal: true})
		curr[lastToken(entry.Category)] = entry.Description
	}
	return map[string]any{"meta-data": nestedMap}
}

func DynamicDocs() map[string]any {
	nestedMap := map[string]any{}
	for _, entry := range docs.DynamicCategoryEntries {
		curr := mapAt(nestedMap, lookup{path: entry.Category, terminal: true})
		curr[lastToken(entry.Category)] = entry.Description
	}
	return map[string]any{"dynamic": nestedMap}
}

// TODO: create file tree table with help messages...
func TablePrintMetadata(indentation int) string {
	table := tablewriter.NewWriter(os.Stdout)
	t := reflect.TypeOf(docs.InstanceMetadataCategoryEntries[0])
	var headers []string
	for i := 0; i < t.NumField(); i++ {
		headers = append(headers, t.Field(i).Name)
	}
	table.SetHeader(headers)
	// data := [][]string{}
	// for _, entry := range docs.InstanceMetadataCategoryEntries {
	// }
	table.Render() // Send output
	return ""
}

func getLookups(path string, resp []byte) []lookup {
	var lookups []lookup
	for _, line := range strings.Split(strings.Trim(string(resp), "\n"), "\n") {
		if strings.HasSuffix(line, "/") {
			lookups = append(lookups, lookup{path: fmt.Sprintf("%s/%s", strings.Trim(path, "/"), line), terminal: false})
		} else {
			lookups = append(lookups, lookup{path: fmt.Sprintf("%s/%s", strings.Trim(path, "/"), line), terminal: true})
		}
	}
	return lookups
}

func mapAt(all map[string]any, l lookup) map[string]any {
	tokens := strings.Split(l.path, "/")
	for i, p := range tokens {
		if l.terminal && i == len(tokens)-1 {
			return all
		} else if all[p] == nil {
			all[p] = make(map[string]any)
			all = all[p].(map[string]any)
		} else {
			all = all[p].(map[string]any)
		}
	}
	return all
}

func lastToken(path string) string {
	tokens := strings.Split(path, "/")
	return tokens[len(tokens)-1]
}

func asJSON(str []byte) (map[string]any, bool) {
	var jsMap map[string]any
	if err := json.Unmarshal(str, &jsMap); err != nil {
		return nil, false
	}
	return jsMap, true
}
