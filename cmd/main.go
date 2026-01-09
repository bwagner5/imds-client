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

	"github.com/bwagner5/imds/pkg/imds"
	"github.com/bwagner5/imds/pkg/tui"
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

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	rootCmd := &cobra.Command{
		Use:   "imds [path]",
		Short: "EC2 Instance Metadata Service CLI",
		Long:  "A CLI for accessing EC2 Instance Metadata Service (IMDS) information.\nRun without arguments to launch interactive TUI explorer.",
		Example: `  imds                    # Launch interactive TUI
  imds instance-id        # Get specific value
  imds placement/region   # Get nested value
  imds -r                 # Tree view of all keys
  imds --dump             # Dump all keys with values
  imds spot --dump        # Dump specific path`,
		Args: cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Version {
				fmt.Printf("Version: %s\nCommit: %s\n", version, commit)
				return nil
			}
			return run(cmd.Context(), args)
		},
	}

	rootCmd.Flags().StringVarP(&opts.Endpoint, "endpoint", "e", envOr("IMDS_ENDPOINT", imds.DefaultEndpoint), "IMDS endpoint")
	rootCmd.Flags().BoolVarP(&opts.Recursive, "recursive", "r", false, "List paths recursively (tree, keys only)")
	rootCmd.Flags().BoolVarP(&opts.Dump, "dump", "d", false, "Dump all paths with values")
	rootCmd.Flags().BoolVarP(&opts.JSON, "json", "j", false, "Output as JSON")
	rootCmd.Flags().BoolVarP(&opts.Watch, "watch", "w", false, "Watch for changes")
	rootCmd.Flags().BoolVar(&opts.Version, "version", false, "Show version")

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string) error {
	client, err := imds.NewClient(ctx, opts.Endpoint)
	if err != nil {
		return fmt.Errorf("creating client: %w", err)
	}

	path := strings.Join(args, "/")

	// JSON flag always dumps all data as JSON
	if opts.JSON {
		data := client.GetAll(ctx, imds.NormalizePath(path))
		enc, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println(string(enc))
		return nil
	}

	// Launch TUI if no args and no output flags
	if path == "" && !opts.Dump && !opts.Recursive && !opts.Watch {
		return tui.Run(ctx, client)
	}

	if opts.Watch {
		return watch(ctx, client, imds.NormalizePath(path))
	}

	if opts.Dump || opts.Recursive {
		return dumpOrTree(ctx, client, imds.NormalizePath(path))
	}

	return query(ctx, client, path)
}

func query(ctx context.Context, client *imds.Client, path string) error {
	// Smart lookup for simple keys (no slashes)
	if !strings.Contains(path, "/") {
		if found := client.FindKey(ctx, path); found != "" {
			resp, err := client.Get(ctx, found)
			if err == nil && !imds.IsDirectory(resp) {
				return output(resp)
			}
		}
	}

	resp, err := client.Get(ctx, imds.NormalizePath(path))
	if err != nil {
		keys := client.AllKeys(ctx)
		similar := imds.FindSimilar(path, keys, 5)
		if len(similar) > 0 {
			fmt.Fprintf(os.Stderr, "Key %q not found. Did you mean:\n", path)
			for _, s := range similar {
				fmt.Fprintf(os.Stderr, "  - %s\n", s)
			}
			return nil
		}
		return fmt.Errorf("key %q not found", path)
	}
	return output(resp)
}

func output(resp []byte) error {
	var js any
	if json.Unmarshal(resp, &js) == nil {
		enc, _ := json.MarshalIndent(js, "", "  ")
		fmt.Println(string(enc))
		return nil
	}
	fmt.Println(strings.TrimSpace(string(resp)))
	return nil
}

func dumpOrTree(ctx context.Context, client *imds.Client, path string) error {
	data := client.GetAll(ctx, path)

	if opts.Dump {
		printDump(data, 0)
	} else {
		printTree(data, 0)
	}
	return nil
}

func watch(ctx context.Context, client *imds.Client, path string) error {
	for data := range client.Watch(ctx, path) {
		enc, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println(string(enc))
	}
	return nil
}

func printTree(data any, depth int) {
	m, ok := data.(map[string]any)
	if !ok {
		return
	}
	indent := strings.Repeat("  ", depth)
	for key, val := range m {
		if _, isMap := val.(map[string]any); isMap {
			fmt.Printf("%s%s/\n", indent, key)
			printTree(val, depth+1)
		} else {
			fmt.Printf("%s%s\n", indent, key)
		}
	}
}

func printDump(data any, depth int) {
	m, ok := data.(map[string]any)
	if !ok {
		return
	}
	indent := strings.Repeat("  ", depth)
	for key, val := range m {
		switch v := val.(type) {
		case map[string]any:
			fmt.Printf("%s%s/\n", indent, key)
			printDump(v, depth+1)
		case []any:
			fmt.Printf("%s%s:\n", indent, key)
			printList(v, depth+1)
		case string:
			if js := tryParseJSON(v); js != nil {
				fmt.Printf("%s%s:\n", indent, key)
				printValue(js, depth+1)
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

func printList(items []any, depth int) {
	indent := strings.Repeat("  ", depth)
	for _, item := range items {
		if m, ok := item.(map[string]any); ok {
			fmt.Printf("%s-\n", indent)
			printValue(m, depth+1)
		} else {
			fmt.Printf("%s- %v\n", indent, item)
		}
	}
}

func printValue(data any, depth int) {
	indent := strings.Repeat("  ", depth)
	switch v := data.(type) {
	case map[string]any:
		for key, val := range v {
			switch inner := val.(type) {
			case map[string]any:
				fmt.Printf("%s%s/\n", indent, key)
				printValue(inner, depth+1)
			case []any:
				fmt.Printf("%s%s:\n", indent, key)
				printList(inner, depth+1)
			default:
				fmt.Printf("%s%s: %v\n", indent, key, val)
			}
		}
	case []any:
		printList(v, depth)
	default:
		fmt.Printf("%s%v\n", indent, v)
	}
}

func tryParseJSON(s string) any {
	var v any
	if json.Unmarshal([]byte(s), &v) == nil {
		switch v.(type) {
		case map[string]any, []any:
			return v
		}
	}
	return nil
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
