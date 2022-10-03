package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

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

func main() {
	options := MustParseFlags()
	if options.Version {
		fmt.Printf("%s\n", version)
		fmt.Printf("Git Commit: %s\n", commit)
		os.Exit(0)
	}
	ctx := context.Background()
	_, err := imds.NewClient(ctx, options.MetadataIP, options.MetadataIPMode)
	if err != nil {
		log.Fatalln(err)
	}
}

func MustParseFlags() Options {
	options := Options{}
	root := flag.NewFlagSet(path.Base(os.Args[0]), flag.ExitOnError)
	root.StringVar(&options.MetadataIP, "metadata-ip", envutil.WithDefault("METADATA_IP", "http://169.254.169.254"), "The IP address of the instance metadata service")
	root.StringVar(&options.MetadataIPMode, "metadata-ip-mode", envutil.WithDefault("METADATA_IP_MODE", "ipv4"), "IP mode (ipv4 or ipv6) to access the instance metadata service")
	root.BoolVar(&options.Version, "version", false, "version information")
	root.Parse(os.Args[1:])
	return options
}
