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

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/bwagner5/imds-client/pkg/imds"
	"github.com/spf13/cobra"
)

var (
	version string
	commit  string
)

type Options struct {
	Endpoint string
	Recurse  bool
	Watch    bool
	Version  bool
	Help     bool
}

var (
	opts = &Options{}
)

// Examples:
// > imds metadata region
// > imds metadata placement availability-zone

var rootCmd = &cobra.Command{
	Use: "imds [path]",
	Example: `  imds meta-data/region
  imds /meta-data/network --recurse`,
	Args:  cobra.RangeArgs(0, 100),
	Short: "IMDS is a CLI for interacting with the Amazon EC2 Instance Metadata Service (IMDS)",
	Long: `IMDS is a CLI for interacting with the Amazon EC2 Instance Metadata Service (IMDS). 
	https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html`,
	Run: func(cmd *cobra.Command, args []string) {
		if opts.Version {
			fmt.Printf("Version: %s\n", version)
			fmt.Printf("Commit: %s\n", commit)
			return
		}
		path := strings.Join(args, "/")
		queryIMDS(cmd.Context(), path)
	},
}

var mdCmd = &cobra.Command{
	Use: "meta-data [path]",
	Example: `  imds meta-data region
  imds md network --recurse`,
	Aliases: []string{"md"},
	GroupID: "1",
	Short:   "Retrieve meta-data information",
	Run: func(cmd *cobra.Command, args []string) {
		path := strings.Join(args, "/")
		if strings.HasPrefix(path, "/") {
			path = strings.Replace(path, "/", "", 1)
		}
		queryIMDS(cmd.Context(), fmt.Sprintf("/meta-data/%s", path))
	},
}
var dynCmd = &cobra.Command{
	Use: "dynamic",
	Example: `  imds dynamic/instance-identity
  imds dyn --recurse`,
	Aliases: []string{"dyn"},
	GroupID: "1",
	Short:   "Retrieve dynamic data",
	Run: func(cmd *cobra.Command, args []string) {
		path := strings.Join(args, "/")
		if strings.HasPrefix(path, "/") {
			path = strings.Replace(path, "/", "", 1)
		}
		queryIMDS(cmd.Context(), fmt.Sprintf("/dynamic/%s", path))
	},
}
var udCmd = &cobra.Command{
	Use:     "user-data",
	Example: `  imds ud`,
	Aliases: []string{"ud"},
	GroupID: "1",
	Short:   "Retrieve user-data information",
	Run: func(cmd *cobra.Command, args []string) {
		path := strings.Join(args, "/")
		if strings.HasPrefix(path, "/") {
			path = strings.Replace(path, "/", "", 1)
		}
		queryIMDS(cmd.Context(), fmt.Sprintf("/user-data/%s", path))
	},
}

func queryIMDS(ctx context.Context, path string) {
	imdsClient, err := imds.NewClient(ctx, opts.Endpoint)
	if err != nil {
		fmt.Printf("Unable to create IMDS client with endpoint %s: %v", opts.Endpoint, err)
		os.Exit(1)
	}
	if opts.Recurse {
		if opts.Watch {
			watchChan := imdsClient.WatchRecurse(ctx, path)
			for {
				select {
				case <-ctx.Done():
					return
				case update := <-watchChan:
					js, err := json.MarshalIndent(update, "", "    ")
					if err != nil {
						fmt.Printf("Unable to recurse starting with path %s: %v", path, err)
						os.Exit(1)
					}
					fmt.Println(string(js))
				}
			}
		}
		js, err := json.MarshalIndent(imdsClient.GetRecurse(ctx, path), "", "    ")
		if err != nil {
			fmt.Printf("Unable to recurse starting with path %s: %v", path, err)
			os.Exit(1)
		}
		fmt.Println(string(js))
		return
	}
	out, err := imdsClient.Get(ctx, path)
	if err != nil {
		fmt.Printf("Unable to retrieve path \"%s\": %v", path, err)
	}
	var jsMap map[string]interface{}
	if err := json.Unmarshal(out, &jsMap); err == nil {
		js, err := json.MarshalIndent(jsMap, "", "    ")
		if err != nil {
			fmt.Printf("Unable to pretty print json for path %s: %v", path, err)
			os.Exit(1)
		}
		fmt.Println(string(js))
		return
	}
	fmt.Println(string(out))
}

func main() {
	ctx := context.Background()
	rootCmd.PersistentFlags().BoolVarP(&opts.Watch, "watch", "w", false, "Watch an IMDS path and print changes to stdout")
	rootCmd.PersistentFlags().BoolVarP(&opts.Recurse, "recurse", "r", false, "Recurse down IMDS paths and return all sub-paths as a JSON doc")
	rootCmd.PersistentFlags().StringVarP(&opts.Endpoint, "endpoint", "e", WithDefault("ENDPOINT", "http://169.254.169.254:80"), "The endpoint to use to connect to IMDS")
	rootCmd.PersistentFlags().BoolVar(&opts.Version, "version", false, "version information")
	rootCmd.AddGroup(&cobra.Group{ID: "1", Title: "Query Groups"})
	rootCmd.AddCommand(mdCmd, dynCmd, udCmd)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func WithDefault(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}
