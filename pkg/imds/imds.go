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
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
)

const DefaultEndpoint = "http://169.254.169.254"

// Client wraps the AWS IMDS client with additional functionality.
type Client struct {
	*imds.Client
}

// NewClient creates a new IMDS client with the specified endpoint.
func NewClient(ctx context.Context, endpoint string) (*Client, error) {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}
	cfg, err := config.LoadDefaultConfig(ctx, withIMDSEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	return &Client{Client: imds.NewFromConfig(cfg)}, nil
}

func withIMDSEndpoint(endpoint string) func(*config.LoadOptions) error {
	return func(lo *config.LoadOptions) error {
		lo.EC2IMDSEndpoint = endpoint
		if net.ParseIP(endpoint).To4() == nil {
			lo.EC2IMDSEndpointMode = imds.EndpointModeStateIPv6
		} else {
			lo.EC2IMDSEndpointMode = imds.EndpointModeStateIPv4
		}
		return nil
	}
}

// Get retrieves raw data from the specified IMDS path.
func (c *Client) Get(ctx context.Context, path string) ([]byte, error) {
	path = strings.Trim(path, "/")
	if path == "" {
		return []byte("meta-data/\ndynamic/\nuser-data"), nil
	}

	switch {
	case strings.HasPrefix(path, "dynamic"):
		subPath := strings.TrimPrefix(strings.TrimPrefix(path, "dynamic"), "/")
		resp, err := c.Client.GetDynamicData(ctx, &imds.GetDynamicDataInput{Path: subPath})
		if err != nil {
			return nil, err
		}
		return io.ReadAll(resp.Content)
	case strings.HasPrefix(path, "meta-data"):
		subPath := strings.TrimPrefix(strings.TrimPrefix(path, "meta-data"), "/")
		resp, err := c.Client.GetMetadata(ctx, &imds.GetMetadataInput{Path: subPath})
		if err != nil {
			return nil, err
		}
		return io.ReadAll(resp.Content)
	case strings.HasPrefix(path, "user-data"):
		resp, err := c.Client.GetUserData(ctx, &imds.GetUserDataInput{})
		if err != nil {
			return nil, err
		}
		return io.ReadAll(resp.Content)
	default:
		return nil, fmt.Errorf("unsupported path: %s", path)
	}
}

// GetAll recursively retrieves all IMDS data from the specified path.
// If path is empty, retrieves all data from all categories.
func (c *Client) GetAll(ctx context.Context, path string) map[string]any {
	path = strings.Trim(path, "/")
	result := map[string]any{}
	errs := map[string]bool{}

	type item struct {
		path     string
		terminal bool
	}

	var queue []item
	if path == "" {
		queue = []item{{path: "dynamic"}, {path: "meta-data"}, {path: "user-data"}}
	} else {
		queue = []item{{path: path}}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		resp, err := c.Get(ctx, cur.path)
		if err != nil {
			if !errs[cur.path] {
				// Retry parent as terminal
				tokens := strings.Split(cur.path, "/")
				if len(tokens) > 1 {
					queue = append(queue, item{path: strings.Join(tokens[:len(tokens)-1], "/"), terminal: true})
				}
				errs[cur.path] = true
			}
			continue
		}

		if _, ok := parseJSON(resp); ok {
			cur.terminal = true
		}

		if strings.HasPrefix(cur.path, "user-data") {
			result["user-data"] = string(resp)
			continue
		}

		if cur.terminal {
			c.setValueAt(result, cur.path, resp)
		} else {
			respStr := string(resp)
			for _, line := range strings.Split(strings.Trim(respStr, "\n"), "\n") {
				if line == "" {
					continue
				}
				isDir := strings.HasSuffix(line, "/")
				cleanLine := strings.TrimSuffix(line, "/")
				queue = append(queue, item{
					path:     cur.path + "/" + cleanLine,
					terminal: !isDir,
				})
			}
		}
	}
	return result
}

func (c *Client) setValueAt(root map[string]any, path string, resp []byte) {
	tokens := strings.Split(path, "/")
	key := tokens[len(tokens)-1]

	// Navigate to parent
	m := root
	for _, t := range tokens[:len(tokens)-1] {
		if m[t] == nil {
			m[t] = make(map[string]any)
		}
		if next, ok := m[t].(map[string]any); ok {
			m = next
		} else {
			m[t] = make(map[string]any)
			m = m[t].(map[string]any)
		}
	}

	// Set value
	switch {
	case key == "pkcs7" || key == "signature" || key == "rsa2048":
		m[key] = string(resp)
	default:
		if js, ok := parseJSON(resp); ok {
			m[key] = js
		} else if strings.Contains(string(resp), "\n") {
			m[key] = strings.Split(string(resp), "\n")
		} else {
			m[key] = string(resp)
		}
	}
}

// FindKey searches for a key name and returns its full path.
func (c *Client) FindKey(ctx context.Context, key string) string {
	for _, base := range []string{"meta-data", "dynamic"} {
		data := c.GetAll(ctx, base)
		if baseData, ok := data[base].(map[string]any); ok {
			if path := findKeyIn(baseData, "", key); path != "" {
				return base + "/" + path
			}
		}
	}
	return ""
}

func findKeyIn(data any, prefix, key string) string {
	m, ok := data.(map[string]any)
	if !ok {
		return ""
	}

	var best string
	maxDepth := -1

	for k, v := range m {
		path := k
		if prefix != "" {
			path = prefix + "/" + k
		}
		if k == key && strings.Count(path, "/") > maxDepth {
			best = path
			maxDepth = strings.Count(path, "/")
		}
		if found := findKeyIn(v, path, key); found != "" && strings.Count(found, "/") > maxDepth {
			best = found
			maxDepth = strings.Count(found, "/")
		}
	}
	return best
}

// Watch monitors the specified path for changes and sends updates to the returned channel.
func (c *Client) Watch(ctx context.Context, path string) <-chan map[string]any {
	ch := make(chan map[string]any, 10)
	go func() {
		defer close(ch)
		var prev map[string]any
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		check := func() {
			data := c.GetAll(ctx, path)
			if !reflect.DeepEqual(prev, data) {
				select {
				case ch <- data:
				default:
					<-ch
					ch <- data
				}
				prev = data
			}
		}

		check() // Initial check
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				check()
			}
		}
	}()
	return ch
}

// NormalizePath adds the appropriate prefix (meta-data/) if not present.
func NormalizePath(path string) string {
	path = strings.Trim(path, "/")
	if path == "" {
		return ""
	}
	if strings.HasPrefix(path, "meta-data/") ||
		strings.HasPrefix(path, "dynamic/") ||
		strings.HasPrefix(path, "user-data") {
		return path
	}
	return "meta-data/" + path
}

// IsDirectory returns true if the response looks like a directory listing.
func IsDirectory(resp []byte) bool {
	content := strings.TrimSpace(string(resp))
	return strings.Contains(content, "\n") && !strings.Contains(content, " ")
}

func parseJSON(data []byte) (any, bool) {
	var v any
	if json.Unmarshal(data, &v) != nil {
		return nil, false
	}
	switch v.(type) {
	case map[string]any, []any:
		return v, true
	}
	return nil, false
}

// AllKeys returns all leaf key names from the IMDS data.
func (c *Client) AllKeys(ctx context.Context) []string {
	data := c.GetAll(ctx, "")
	var keys []string
	collectKeys(data, &keys)
	return keys
}

func collectKeys(data any, keys *[]string) {
	m, ok := data.(map[string]any)
	if !ok {
		return
	}
	for k, v := range m {
		if _, isMap := v.(map[string]any); isMap {
			collectKeys(v, keys)
		} else {
			*keys = append(*keys, k)
		}
	}
}

// FindSimilar returns keys similar to the query, sorted by similarity.
func FindSimilar(query string, keys []string, maxResults int) []string {
	type scored struct {
		key   string
		score int
	}
	var matches []scored
	query = strings.ToLower(query)

	for _, k := range keys {
		kLower := strings.ToLower(k)
		score := 0

		// Exact substring match
		if strings.Contains(kLower, query) || strings.Contains(query, kLower) {
			score += 100
		}

		// Levenshtein-like: count matching characters
		matching := countMatching(query, kLower)
		score += matching * 2

		// Prefix match bonus
		prefixLen := len(query)
		if len(kLower) < prefixLen {
			prefixLen = len(kLower)
		}
		if prefixLen > 0 && strings.HasPrefix(kLower, query[:prefixLen]) {
			score += 50
		}

		// Only include if reasonably similar
		threshold := len(query) * 2
		if score >= threshold {
			matches = append(matches, scored{k, score})
		}
	}

	// Sort by score descending
	for i := range matches {
		for j := i + 1; j < len(matches); j++ {
			if matches[j].score > matches[i].score {
				matches[i], matches[j] = matches[j], matches[i]
			}
		}
	}

	var result []string
	seen := make(map[string]bool)
	for i := 0; i < len(matches) && len(result) < maxResults; i++ {
		if !seen[matches[i].key] {
			result = append(result, matches[i].key)
			seen[matches[i].key] = true
		}
	}
	return result
}

func countMatching(a, b string) int {
	count := 0
	used := make([]bool, len(b))
	for _, c := range a {
		for i, bc := range b {
			if !used[i] && c == bc {
				count++
				used[i] = true
				break
			}
		}
	}
	return count
}
