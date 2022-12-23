package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var (
	//| Category | Description | Version when category was released |
	mdTableRe     = regexp.MustCompile(`(\|.*\|.*(\n)?)+`)
	headerDivider = regexp.MustCompile(`(\|?[ ]*[\-]+[ ]*\|)+`)
	header        = `/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/`
)

func main() {
	doc, err := retrieveDoc("https://raw.githubusercontent.com/awsdocs/amazon-ec2-user-guide/master/doc_source/instancedata-data-categories.md")
	if err != nil {
		panic(err)
	}
	mdTables := mdTableRe.FindAll(doc, -1)
	if len(mdTables) != 2 {
		panic(fmt.Sprintf("Should have found 2 markdown tables, but found %d", len(mdTables)))
	}
	metadataTable := mdTables[0]
	dynamicTable := mdTables[1]
	src := &bytes.Buffer{}
	fmt.Fprintln(src, header)
	fmt.Fprintln(src, "package doc")
	fmt.Fprintln(src, "// DO NOT EDIT")
	fmt.Fprintln(src, "// THIS FILE IS AUTO GENERATED")
	fmt.Fprintln(src, generateStruct("InstanceMetadataCategory", metadataTable))
	fmt.Fprintln(src, generateStruct("DynamicCategory", dynamicTable))
	fmt.Fprintln(src, generatePopulatedSliceVar("InstanceMetadataCategory", metadataTable))
	fmt.Fprintln(src, generatePopulatedSliceVar("DynamicCategory", dynamicTable))

	formatted, err := format.Source(src.Bytes())
	if err != nil {
		log.Fatalf("formatting generated source, %s", err)
	}
	fmt.Println(string(formatted))
}

func retrieveDoc(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func generateStruct(name string, table []byte) string {
	src := &bytes.Buffer{}
	fmt.Fprintf(src, "type %s struct {\n", name)
	for _, header := range getHeaders(table) {
		fmt.Fprintf(src, "%s string\n", header)
	}
	fmt.Fprintln(src, "}")
	return src.String()
}

func generatePopulatedSliceVar(name string, table []byte) string {
	src := &bytes.Buffer{}
	headers := getHeaders(table)
	fmt.Fprintf(src, "var %sEntries = []%s{\n", name, name)
	for _, row := range getRows(table) {
		cols := getCols(row)
		if len(cols) == 0 {
			continue
		}
		if len(headers) != len(cols) {
			panic(fmt.Sprintf("The number of headers (%d) does not match the number of columns (%d): \n Headers: %s \n Columns: %s\n",
				len(headers), len(cols), strings.Join(headers, "|"), strings.Join(cols, "|")))
		}
		fmt.Fprintln(src, "{")
		for i, header := range headers {
			fmt.Fprintf(src, "%s: \"%s\",\n", header, cols[i])
		}
		fmt.Fprintln(src, "},")
	}
	fmt.Fprintln(src, "}")
	return src.String()
}

func getHeaders(table []byte) []string {
	rows := strings.Split(string(table), "\n")
	if len(rows) == 0 {
		panic("No rows found...")
	}
	var headers []string
	for _, header := range strings.Split(rows[0], "|") {
		trimmedHeader := strings.TrimSpace(header)
		if trimmedHeader == "" {
			continue
		}
		headers = append(headers, strings.Split(trimmedHeader, " ")[0])
	}
	return headers
}

func getRows(table []byte) []string {
	var rows []string
	var afterHeader bool
	for _, r := range strings.Split(string(table), "\n") {
		if afterHeader {
			rows = append(rows, r)
		} else if headerDivider.Match([]byte(r)) {
			afterHeader = true
		}
	}
	return rows
}

func getCols(row string) []string {
	var cols []string
	bracketCtrl := "06389"
	row = strings.ReplaceAll(row, `\|`, bracketCtrl)
	for _, col := range strings.Split(row, "|") {
		trimmedCol := strings.TrimSpace(col)
		if trimmedCol == "" {
			continue
		}
		trimmedCol = strings.ReplaceAll(trimmedCol, `\`, "")
		trimmedCol = strings.ReplaceAll(trimmedCol, `"`, `\"`)
		trimmedCol = strings.ReplaceAll(trimmedCol, bracketCtrl, `|`)
		cols = append(cols, trimmedCol)
	}
	return cols
}
