package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"reflect"
	"strings"
)

type metadata struct {
	ramDiskID                       string   `imds:"path=meta-data/ramdisk-id"`
	reservationID                   string   `imds:"path=meta-data/reservation-id"`
	securityGroups                  []string `imds:"path=meta-data/security-groups"`
	availabilityZone                string   `imds:"path=meta-data/placement/availability-zone"`
	availabilityZoneID              string   `imds:"path=meta-data/placement/availability-zone-id"`
	groupName                       string   `imds:"path=meta-data/placement/group-name"`
	hostID                          string   `imds:"path=meta-data/placement/host-id"`
	partitionNumber                 int      `imds:"path=meta-data/placement/partition-number"`
	region                          string   `imds:"path=meta-data/placement/region"`
	productCodes                    []string `imds:"path=meta-data/product-codes"`
	publicHostname                  string   `imds:"path=meta-data/public-hostname"`
	publicIPv4                      string   `imds:"path=meta-data/public-ipv4"`
	localHostname                   string   `imds:"path=meta-data/local-hostname"`
	localIPv4                       string   `imds:"path=meta-data/local-ipv4"`
	mac                             string   `imds:"path=meta-data/mac"`
	instanceAction                  string   `imds:"path=meta-data/instance-action"`
	instanceID                      string   `imds:"path=meta-data/instance-id"`
	instanceLifecycle               string   `imds:"path=meta-data/instance-life-cycle"`
	instanceType                    string   `imds:"path=meta-data/instance-type"`
	kernelID                        string   `imds:"path=meta-data/kernel-id"`
	amiID                           string   `imds:"path=meta-data/ami-id"`
	amiLaunchIndex                  int      `imds:"path=meta-data/ami-launch-index"`
	amiManifestPath                 string   `imds:"path=meta-data/ami-manifest-path"`
	ancestorAMIIDs                  []string `imds:"path=meta-data/ancestor-ami-ids"`
	autoscalingTargetLifecycleState string   `imds:"path=meta-data/autoscaling/target-lifecycle-state"`
	blockDeviceMappingAMI           string   `imds:"path=meta-data/block-device-mapping/ami"`
	blockDeviceMappingRoot          []string `imds:"path=meta-data/block-device-mapping/root"`
	eventsMaintenanceHistory        string   `imds:"path=meta-data/events/maintenance/history"`
	eventsMaintenanceScheduled      string   `imds:"path=meta-data/events/maintenance/scheduled"`
	eventsRecommendationsRebalance  string   `imds:"path=meta-data/events/recommendations/rebalance"`
	iamInfo                         string   `imds:"path=meta-data/iam/info"`
}

func main() {
	src := &bytes.Buffer{}
	fmt.Fprintln(src, "package imds")
	fmt.Fprintln(src, "// DO NOT EDIT")
	fmt.Fprintln(src, "// THIS FILE IS AUTO GENERATED")
	fmt.Fprintln(src, `import (
		"context"
		"fmt"
		"strconv"
		)`)

	fmt.Fprintln(src, genStructs())
	// fmt.Fprintln(src, genKongCMD())
	t := reflect.TypeOf(metadata{})
	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		for _, must := range []bool{true, false} {
			for _, context := range []bool{true, false} {

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
					fmt.Fprintln(src, GenMetadataStringFunc(field.Name, path, must, context))
					fmt.Fprintln(src, GenMetadataStringFunc(field.Name, strings.Replace(path, "meta-data/", "", 1), must, context))
				case "int":
					fmt.Fprintln(src, GenMetadataIntFunc(field.Name, path, must, context))
					fmt.Fprintln(src, GenMetadataIntFunc(field.Name, strings.Replace(path, "meta-data/", "", 1), must, context))
				case "[]string":
					fmt.Fprintln(src, GenMetadataStringSliceFunc(field.Name, path, must, context))
					fmt.Fprintln(src, GenMetadataStringSliceFunc(field.Name, strings.Replace(path, "meta-data/", "", 1), must, context))
				default:
				}
			}
		}
	}

	formatted, err := format.Source(src.Bytes())
	if err != nil {
		log.Fatalf("formatting generated source, %s", err)
	}

	fmt.Println(string(formatted))
}

func genStructs() string {
	t := reflect.TypeOf(metadata{})
	s := &bytes.Buffer{}
	fmt.Fprintln(s, "type metadata struct {")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Fprintf(s, "%s %s `%s`\n", field.Name, field.Type, field.Tag)
	}
	fmt.Fprint(s, "}")
	return s.String()
}

func genKongCMD() string {
	t := reflect.TypeOf(metadata{})
	s := &bytes.Buffer{}
	fmt.Fprintln(s, "type MetadataCmd struct {")
	fmt.Fprintln(s, "Path string `arg:\"\" name:\"path\" help:\"Metadata path to retrieve\" type:\"path\"`")
	fmt.Fprintln(s, "IMDSClient *IMDS")
	fmt.Fprintln(s, "}")
	fmt.Fprintln(s, `func (c MetadataCmd) Run(ctx context.Context) error {
		resp, err := c.imdsClient.GetMetadata(ctx, c.Path)
		if err != nil {
			return err
		}
		fmt.Print(resp)
		return nil
		}`)
	fmt.Fprintln(s, "func (c MetadataCmd) Help() string {")
	fmt.Fprintf(s, "return `")
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
		fmt.Fprintln(s, strings.Replace(path, "meta-data/", "", 1))
	}
	fmt.Fprintf(s, "`")
	fmt.Fprintf(s, "}")
	return s.String()
}

func funcSignature(fieldName string, returnType string, must bool, context bool) string {
	methodFieldName := fmt.Sprint(strings.ToUpper(string(fieldName[0])), fieldName[1:])
	prefix := ""
	suffix := ""
	inputArgs := ""
	returnArgs := fmt.Sprintf("(%s, error)", returnType)
	if must {
		prefix = "Must"
		returnArgs = returnType
	}
	if context {
		suffix = "WithContext"
		inputArgs = "ctx context.Context"
	}
	name := fmt.Sprintf("%sGet%s%s", prefix, methodFieldName, suffix)
	method := &bytes.Buffer{}
	fmt.Fprintf(method, "func (i IMDS) %s(%s) %s {", name, inputArgs, returnArgs)
	if !context {
		fmt.Fprint(method, "\nctx := context.Background()")
	}
	return method.String()
}

func GenMetadataStringSliceFunc(fieldName string, path string, must bool, context bool) string {
	method := &bytes.Buffer{}
	fmt.Fprintln(method, funcSignature(fieldName, "[]string", must, context))
	fmt.Fprintf(method, "%s, err := i.GetMetadata(ctx, \"%s\")\n", fieldName, path)
	if must {
		fmt.Fprintln(method, "if err  != nil { ")
		fmt.Fprintf(method, "panic(fmt.Sprintf(\"unable to fetch %s: %%v\", err))\n", fieldName)
		fmt.Fprintln(method, "}")
		fmt.Fprintf(method, "return strings.Split(%s, \"\n\")\n", fieldName)
	} else {
		fmt.Fprintln(method, `if err != nil { return "", err }`)
		fmt.Fprintf(method, "return strings.Split(%s, \"\n\"), nil\n", fieldName)
	}
	fmt.Fprintln(method, "}")
	return method.String()
}

func GenMetadataIntFunc(fieldName string, path string, must bool, context bool) string {
	method := &bytes.Buffer{}
	fmt.Fprintln(method, funcSignature(fieldName, "int", must, context))
	fmt.Fprintf(method, "%s, err := i.GetMetadata(ctx, \"%s\")\n", fieldName, path)
	if must {
		fmt.Fprintln(method, "if err  != nil { ")
		fmt.Fprintf(method, "panic(fmt.Sprintf(\"unable to fetch %s: %%v\", err))\n", fieldName)
		fmt.Fprintln(method, "}")
		fmt.Fprintf(method, "%sNum, err := strconv.Atoi(%s)\n", fieldName, fieldName)
		fmt.Fprintln(method, "if err  != nil { ")
		fmt.Fprintf(method, "panic(fmt.Sprintf(\"unable to convert %s of %%s to integer: %%v\", %s, err))\n", path, fieldName)
		fmt.Fprintln(method, "}")
		fmt.Fprintf(method, "return %sNum\n", fieldName)
	} else {
		fmt.Fprintln(method, `if err != nil { return 0, err }`)
		fmt.Fprintf(method, "%sNum, err := strconv.Atoi(%s)\n", fieldName, fieldName)
		fmt.Fprintf(method, "if err  != nil { return 0, fmt.Errorf(\"unable to convert %s of %%s to integer: %%w\", %s, err) }\n", path, fieldName)
		fmt.Fprintf(method, "return %sNum, nil\n", fieldName)
	}
	fmt.Fprintln(method, "}")
	return method.String()
}

func GenMetadataStringFunc(fieldName string, path string, must bool, context bool) string {
	method := &bytes.Buffer{}
	fmt.Fprintln(method, funcSignature(fieldName, "string", must, context))
	fmt.Fprintf(method, "%s, err := i.GetMetadata(ctx, \"%s\")\n", fieldName, path)
	if must {
		fmt.Fprintln(method, "if err  != nil { ")
		fmt.Fprintf(method, "panic(fmt.Sprintf(\"unable to fetch %s: %%v\", err))\n", fieldName)
		fmt.Fprintln(method, "}")
		fmt.Fprintf(method, "return %s", fieldName)
	} else {
		fmt.Fprintln(method, `if err != nil { return "", err }`)
		fmt.Fprintf(method, "return %s, nil\n", fieldName)
	}
	fmt.Fprintln(method, "}")
	return method.String()
}
