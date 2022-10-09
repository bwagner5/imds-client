package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/bwagner5/imds-client/pkg/imds"
	"github.com/jaypipes/envutil"
)

var (
	version string
	commit  string
)

type Options struct {
	MetadataIP     string
	MetadataIPMode string
	Version        bool
}

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
	imdsClient, err := imds.NewClient(ctx, options.MetadataIP, options.MetadataIPMode)
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
	}
}

func MustParseFlags(f *flag.FlagSet) Options {
	options := Options{}
	f.StringVar(&options.MetadataIP, "metadata-ip", envutil.WithDefault("METADATA_IP", "http://169.254.169.254"), "The IP address of the instance metadata service")
	f.StringVar(&options.MetadataIPMode, "metadata-ip-mode", envutil.WithDefault("METADATA_IP_MODE", "ipv4"), "IP mode (ipv4 or ipv6) to access the instance metadata service")
	f.BoolVar(&options.Version, "version", false, "version information")
	f.Parse(os.Args[1:])
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
