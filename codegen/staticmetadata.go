package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"reflect"
	"strings"
)

type staticPaths struct {
	ramDiskID          string   `imds:"path=ramdisk-id"`
	reservationID      string   `imds:"path=reservation-id"`
	securityGroups     []string `imds:"path=security-groups"`
	availabilityZone   string   `imds:"path=placement/availability-zone"`
	availabilityZoneID string   `imds:"path=placement/availability-zone-id"`
	groupName          string   `imds:"path=placement/group-name"`
	hostID             string   `imds:"path=placement/host-id"`
	partitionNumber    int      `imds:"path=placement/partition-number"`
	region             int      `imds:"path=placement/region"`
	productCodes       []string `imds:"path=product-codes"`
	publicHostname     string   `imds:"path=public-hostname"`
	publicIPv4         string   `imds:"path=public-ipv4"`
	localHostname      string   `imds:"path=local-hostname"`
	localIPv4          string   `imds:"path=local-ipv4"`
	mac                string   `imds:"path=mac"`
	instanceAction     string   `imds:"path=instance-action"`
	instanceID         string   `imds:"path=instance-id"`
	instanceLifecycle  string   `imds:"path=instance-life-cycle"`
	instanceType       string   `imds:"path=instance-type"`
	kernelID           string   `imds:"path=kernel-id"`
	amiID              string   `imds:"path=ami-id"`
	amiLaunchIndex     int      `imds:"path=ami-launch-index"`
	amiManifestPath    string   `imds:"path=ami-manifest-path"`
}

func main() {
	src := &bytes.Buffer{}
	fmt.Fprintln(src, "package imds")
	t := reflect.TypeOf(staticPaths{})
	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		// Get the field tag value
		tag := field.Tag.Get("imds")
		tagProps := strings.Split(tag, "=")
		tagPropsMap := map[string]string{}
		for i := 0; i < len(tagProps)-1; i += 2 {
			tagPropsMap[tagProps[i]] = tagProps[i+1]
		}
		path, ok := tagPropsMap["path"]
		if !ok {
			panic(fmt.Sprintf("field %s has no path tag", field.Name))
		}
		switch field.Type.Name() {
		case "string":
			fmt.Fprintln(src, GenMetadataStringFunc(field.Name, path))
		case "int":
			fmt.Fprintln(src, GenMetadataIntFunc(field.Name, path))
		case "[]string":
			fmt.Fprintln(src, GenMetadataStringSliceFunc(field.Name, path))
		default:
		}
	}

	formatted, err := format.Source(src.Bytes())
	if err != nil {
		log.Fatalf("formatting generated source, %s", err)
	}

	fmt.Println(string(formatted))

	if err := os.WriteFile(flag.Arg(0), formatted, 0644); err != nil {
		log.Fatalf("writing output, %s", err)
	}
}

func GenMetadataStringSliceFunc(fieldName string, path string) string {
	methodFieldName := fmt.Sprint(strings.ToUpper(string(fieldName[0])), fieldName[1:])
	method := &bytes.Buffer{}
	fmt.Fprintf(method, "func (i IMDS) Get%s() []string {\n", methodFieldName)
	fmt.Fprintf(method, "%s, err := i.GetMetadata(ctx, \"%s\")\n", fieldName, path)
	fmt.Fprintln(method, `if err  != nil { return "", err }`)
	fmt.Fprintf(method, "return strings.Split(%s, \"\n\"), nil\n", fieldName)
	fmt.Fprintln(method, "}")
	return method.String()
}

func GenMetadataIntFunc(fieldName string, path string) string {
	methodFieldName := fmt.Sprint(strings.ToUpper(string(fieldName[0])), fieldName[1:])
	method := &bytes.Buffer{}
	fmt.Fprintf(method, "func (i IMDS) Get%s() string {\n", methodFieldName)
	fmt.Fprintf(method, "%s, err := i.GetMetadata(ctx, \"%s\")\n", fieldName, path)
	fmt.Fprintln(method, `if err  != nil { return 0, err }`)
	fmt.Fprintf(method, "%sNum, err := strconv.Atoi(%s)\n", fieldName, fieldName)
	fmt.Fprintf(method, "if err  != nil { return 0, fmt.Errorf(\"unable to convert %s of %%s to integer: %%w\", %s, err) }\n", path, fieldName)
	fmt.Fprintf(method, "return %sNum, nil\n", fieldName)
	fmt.Fprintln(method, "}")
	return method.String()
}

func GenMetadataStringFunc(fieldName string, path string) string {
	methodFieldName := fmt.Sprint(strings.ToUpper(string(fieldName[0])), fieldName[1:])
	method := &bytes.Buffer{}
	fmt.Fprintf(method, "func (i IMDS) Get%s() string {\n", methodFieldName)
	fmt.Fprintf(method, "%s, err := i.GetMetadata(ctx, \"%s\")\n", fieldName, path)
	fmt.Fprintln(method, `if err  != nil { return "", err }`)
	fmt.Fprintf(method, "return %s, nil\n", fieldName)
	fmt.Fprintln(method, "}")
	return method.String()
}
