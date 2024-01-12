// File: provider/cmd/pulumi-resource-pinecone/main.go
package main

import (
	"fmt"
	"os"
	"strings"

	pinecone "github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone"
	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/version"
	p "github.com/pulumi/pulumi-go-provider"
)

// Serve the provider against Pulumi's Provider protocol.
func main() {
	trimmedVersion := strings.TrimPrefix(version.Version, "v")
	err := p.RunProvider("pinecone", trimmedVersion, pinecone.Provider())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
