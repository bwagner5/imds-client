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
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/bwagner5/imds-client/pkg/imds"
)

var (
	version string
	commit  string
)

type Options struct {
	MetadataIP string
	Version    bool
}

// Examples:
// > imds metadata region
// > imds metadata placement availability-zone

func main() {
	root := flag.NewFlagSet(path.Base(os.Args[0]), flag.ExitOnError)
	root.Usage = HelpFunc(root)
	options := MustParseFlags(root)
	if options.Version {
		fmt.Printf("%s\n", version)
		fmt.Printf("Git Commit: %s\n", commit)
		os.Exit(0)
	}
	ctx := context.Background()
	// fmt.Println(options.MetadataIP)
	imdsClient, err := imds.NewClient(ctx, options.MetadataIP)
	if err != nil {
		log.Fatalln(err)
	}
	switch root.Arg(0) {
	case "metadata", "md":
		path := strings.Join(root.Args()[1:], "/")
		resp, err := imdsClient.GetMetadata(ctx, path)
		if err != nil {
			fmt.Printf("oops %s: %v\n", path, err)
			os.Exit(1)
		}
		fmt.Println(resp)
	case "dynamic", "dyn":
		path := strings.Join(root.Args()[1:], "/")
		resp, err := imdsClient.GetDynamicData(ctx, path)
		if err != nil {
			fmt.Printf("oops %s: %v\n", path, err)
			os.Exit(1)
		}
		fmt.Println(resp)
	case "userdata", "ud":
		resp, err := imdsClient.GetUserdata(ctx)
		if err != nil {
			fmt.Printf("oops: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(resp)
	default:
		js, err := json.MarshalIndent(imdsClient.GetAll(ctx, root.Args()[0]), "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(js))
	}
}

func MustParseFlags(f *flag.FlagSet) Options {
	options := Options{}
	f.StringVar(&options.MetadataIP, "metadata-ip", WithDefault("METADATA_IP", "http://169.254.169.254"), "The IP address of the instance metadata service")
	f.BoolVar(&options.Version, "version", false, "version information")
	if err := f.Parse(os.Args[1:]); err != nil {
		panic(fmt.Sprintf("Unable to parse arguments: %v", err))
	}
	return options
}

func HelpFunc(f *flag.FlagSet) func() {
	return func() {
		fmt.Printf("Usage for %s:\n\n", os.Args[0])
		fmt.Println("   metadata | md")
		fmt.Println("      Retrieve instance metadata")
		fmt.Println("   dynamic | dyn")
		fmt.Println("      Retrieve dynamic data")
		fmt.Println("   userdata | ud")
		fmt.Println("      Retrieve userdata")
		fmt.Println("")
		fmt.Println(" Flags:")
		f.VisitAll(func(fl *flag.Flag) {
			fmt.Printf("   --%s\n", fl.Name)
			fmt.Printf("      %s\n", fl.Usage)
		})
	}
}

func WithDefault(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}
