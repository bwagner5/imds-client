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
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/bwagner5/imds-client/pkg/imds"
)

var (
	version string
	commit  string
)

type Options struct {
	Endpoint  string
	Recursive bool
	Dump      bool
	JSON      bool
	Watch     bool
	Version   bool
}

var opts = &Options{}

var rootCmd = &cobra.Command{
	Use:   "imds [path]",
	Short: "EC2 Instance Metadata Service CLI",
	Long:  "A CLI for accessing EC2 Instance Metadata Service (IMDS) information",
	Example: `  imds instance-id
  imds placement/host-id
  imds placement host-id
  imds list
  imds list dynamic
  imds list --recursive`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if opts.Version {
			fmt.Printf("Version: %s\nCommit: %s\n", version, commit)
			return
		}
		
		ctx := cmd.Context()
		client, err := imds.NewClient(ctx, opts.Endpoint)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating IMDS client: %v\n", err)
			os.Exit(1)
		}

		if len(args) == 0 {
			handleDumpAll(ctx, client)
			return
		}

		if args[0] == "list" {
			handleList(ctx, client, args[1:])
			return
		}

		path := strings.Join(args, "/")
		path = strings.ReplaceAll(path, " ", "/")
		handleQuery(ctx, client, path)
	},
}

func handleDumpAll(ctx context.Context, client *imds.Client) {
	if opts.Dump {
		data := client.GetRecurse(ctx, "")
		if opts.JSON {
			js, _ := json.MarshalIndent(data, "", "  ")
			fmt.Println(string(js))
			return
		}
		printDump(data, 0)
		return
	}

	if opts.Recursive {
		data := client.GetRecurse(ctx, "")
		if len(data) == 0 {
			fmt.Println("meta-data/")
			fmt.Println("dynamic/")
			fmt.Println("user-data")
			return
		}
		printTree(data, 0)
		return
	}

	if opts.JSON {
		data := client.GetRecurse(ctx, "")
		js, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println(string(js))
		return
	}

	fmt.Println("meta-data/")
	fmt.Println("dynamic/")
	fmt.Println("user-data")
}

func handleList(ctx context.Context, client *imds.Client, args []string) {
	if len(args) == 0 {
		if opts.Recursive {
			data := client.GetRecurse(ctx, "")
			printTree(data, 0)
			return
		}
		fmt.Println("meta-data/")
		fmt.Println("dynamic/")
		fmt.Println("user-data")
		return
	}

	path := args[0]
	if opts.Recursive {
		data := client.GetRecurse(ctx, path)
		if len(data) == 0 {
			fmt.Printf("No data available for %s\n", path)
			return
		}
		printTree(data, 0)
		return
	}

	// List specific category
	switch path {
	case "dynamic":
		printPaths(imds.DynamicDocs()["dynamic"])
	case "meta-data":
		printPaths(imds.MetadataDocs()["meta-data"])
	case "user-data":
		fmt.Println("user-data")
	default:
		resp, err := client.Get(ctx, ensurePrefix(path))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing %s: %v\n", path, err)
			os.Exit(1)
		}

		lines := strings.Split(strings.TrimSpace(string(resp)), "\n")
		for _, line := range lines {
			if line != "" {
				fmt.Println(line)
			}
		}
	}
}

func handleQuery(ctx context.Context, client *imds.Client, path string) {
	if opts.Watch {
		watchPath(ctx, client, ensurePrefix(path))
		return
	}

	if opts.Dump || opts.Recursive {
		data := client.GetRecurse(ctx, ensurePrefix(path))
		if opts.JSON {
			js, _ := json.MarshalIndent(data, "", "  ")
			fmt.Println(string(js))
			return
		}
		if opts.Dump {
			printDump(data, 0)
		} else {
			printTree(data, 0)
		}
		return
	}

	originalPath := path

	// For simple keys (no slashes), always try smart lookup first
	if !strings.Contains(originalPath, "/") {
		if foundPath := client.FindKey(ctx, originalPath); foundPath != "" {
			if smartResp, smartErr := client.Get(ctx, foundPath); smartErr == nil {
				// Check if the smart lookup result is a terminal value (not a directory)
				smartContent := strings.TrimSpace(string(smartResp))
				if !strings.Contains(smartContent, "\n") && !strings.HasSuffix(smartContent, "/") {
					// This is a terminal value, use it
					outputResponse(smartResp, foundPath)
					return
				}
			}
		}
	}

	// Fallback to exact path lookup with prefix
	path = ensurePrefix(path)
	resp, err := client.Get(ctx, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error querying %s: %v\n", path, err)
		os.Exit(1)
	}

	outputResponse(resp, path)
}

func outputResponse(resp []byte, path string) {
	// Try to parse as JSON first
	var jsonData interface{}
	if json.Unmarshal(resp, &jsonData) == nil {
		if opts.JSON {
			js, _ := json.MarshalIndent(jsonData, "", "  ")
			fmt.Println(string(js))
		} else {
			js, _ := json.MarshalIndent(jsonData, "", "  ")
			fmt.Println(string(js))
		}
		return
	}

	// Check if it's a directory listing
	content := strings.TrimSpace(string(resp))
	if strings.Contains(content, "\n") && !strings.Contains(content, " ") {
		// Looks like a directory listing
		lines := strings.Split(content, "\n")
		for _, line := range lines {
			if line != "" {
				fmt.Println(line)
			}
		}
		return
	}

	// Plain text response
	fmt.Print(content)
	if !strings.HasSuffix(content, "\n") {
		fmt.Println()
	}
}

func watchPath(ctx context.Context, client *imds.Client, path string) {
	watchChan := client.WatchRecurse(ctx, path)
	for {
		select {
		case <-ctx.Done():
			return
		case update := <-watchChan:
			js, _ := json.MarshalIndent(update, "", "  ")
			fmt.Println(string(js))
		}
	}
}

func normalizePath(args []string) string {
	path := strings.Join(args, "/")
	path = strings.ReplaceAll(path, " ", "/")
	return ensurePrefix(path)
}

func ensurePrefix(path string) string {
	path = strings.Trim(path, "/")
	if path == "" {
		return ""
	}
	
	// Add appropriate prefix if not present
	if !strings.HasPrefix(path, "meta-data/") && 
	   !strings.HasPrefix(path, "dynamic/") && 
	   !strings.HasPrefix(path, "user-data") {
		// Default to meta-data for common queries
		return "meta-data/" + path
	}
	return path
}

func printPaths(data interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key := range v {
			fmt.Printf("%s/\n", key)
		}
	case string:
		fmt.Println(v)
	}
}

func printTree(data interface{}, depth int) {
	m, ok := data.(map[string]interface{})
	if !ok {
		return
	}
	indent := strings.Repeat("  ", depth)
	for key, value := range m {
		if _, isMap := value.(map[string]interface{}); isMap {
			fmt.Printf("%s%s/\n", indent, key)
			printTree(value, depth+1)
		} else {
			fmt.Printf("%s%s\n", indent, key)
		}
	}
}

func printDump(data interface{}, depth int) {
	m, ok := data.(map[string]interface{})
	if !ok {
		return
	}
	indent := strings.Repeat("  ", depth)
	for key, value := range m {
		switch v := value.(type) {
		case map[string]interface{}:
			fmt.Printf("%s%s/\n", indent, key)
			printDump(v, depth+1)
		case []interface{}:
			fmt.Printf("%s%s:\n", indent, key)
			for _, item := range v {
				if m, ok := item.(map[string]interface{}); ok {
					fmt.Printf("%s  -\n", indent)
					printDumpValue(m, depth+2)
				} else {
					fmt.Printf("%s  - %v\n", indent, item)
				}
			}
		case string:
			// Try to parse as JSON
			var js interface{}
			if json.Unmarshal([]byte(v), &js) == nil {
				fmt.Printf("%s%s:\n", indent, key)
				printDumpValue(js, depth+1)
			} else if strings.Contains(v, "\n") {
				fmt.Printf("%s%s: |\n", indent, key)
				for _, line := range strings.Split(v, "\n") {
					fmt.Printf("%s  %s\n", indent, line)
				}
			} else {
				fmt.Printf("%s%s: %s\n", indent, key, v)
			}
		default:
			fmt.Printf("%s%s: %v\n", indent, key, v)
		}
	}
}

func printDumpValue(data interface{}, depth int) {
	indent := strings.Repeat("  ", depth)
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			switch val := value.(type) {
			case map[string]interface{}:
				fmt.Printf("%s%s/\n", indent, key)
				printDumpValue(val, depth+1)
			case []interface{}:
				fmt.Printf("%s%s:\n", indent, key)
				for _, item := range val {
					if m, ok := item.(map[string]interface{}); ok {
						fmt.Printf("%s  -\n", indent)
						printDumpValue(m, depth+2)
					} else {
						fmt.Printf("%s  - %v\n", indent, item)
					}
				}
			default:
				fmt.Printf("%s%s: %v\n", indent, key, value)
			}
		}
	case []interface{}:
		for _, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				fmt.Printf("%s-\n", indent)
				printDumpValue(m, depth+1)
			} else {
				fmt.Printf("%s- %v\n", indent, item)
			}
		}
	default:
		fmt.Printf("%s%v\n", indent, v)
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	rootCmd.PersistentFlags().StringVarP(&opts.Endpoint, "endpoint", "e", 
		withDefault("IMDS_ENDPOINT", "http://169.254.169.254"), "IMDS endpoint")
	rootCmd.PersistentFlags().BoolVarP(&opts.Recursive, "recursive", "r", false, 
		"List all paths recursively (tree structure, keys only)")
	rootCmd.PersistentFlags().BoolVarP(&opts.Dump, "dump", "d", false, 
		"Dump all paths recursively with values")
	rootCmd.PersistentFlags().BoolVarP(&opts.JSON, "json", "j", false, 
		"Output in JSON format")
	rootCmd.PersistentFlags().BoolVarP(&opts.Watch, "watch", "w", false, 
		"Watch for changes")
	rootCmd.PersistentFlags().BoolVar(&opts.Version, "version", false, 
		"Show version information")

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func withDefault(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
