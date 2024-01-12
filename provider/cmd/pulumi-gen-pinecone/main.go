package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	pinecone "github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone"
	providerVersion "github.com/pinecone-io/pulumi-pinecone/provider/pkg/version"
)

// copied from encoding/json for use with JSONMarshal above
func MarshalIndent(v any) ([]byte, error) {

	// json.Marshal normally escapes HTML. This one doesn't
	// https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	b := buffer.Bytes()

	// serialize and pretty print
	var buf bytes.Buffer
	prefix, indent := "", "    "
	err = json.Indent(&buf, b, prefix, indent)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	flag.Usage = func() {
		const usageFormat = "Usage: %s <schema-file>"
		fmt.Fprintf(flag.CommandLine.Output(), usageFormat, os.Args[0])
		flag.PrintDefaults()
	}

	var version string
	flag.StringVar(&version, "version", providerVersion.Version, "the provider version to record in the generated code")

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}
	s, err := pinecone.Schema(version)
	if err != nil {
		panic(err)
	}

	// sort keys
	var arg map[string]any
	err = json.Unmarshal([]byte(s), &arg)
	if err != nil {
		panic(err)
	}

	// remove version key
	delete(arg, "version")

	// use custom marshal indent to skip html escaping
	out, err := MarshalIndent(arg)
	if err != nil {
		panic(err)
	}

	schemaPath := args[0]
	err = os.WriteFile(schemaPath, out, 0600)
	if err != nil {
		panic(err)
	}
}
