// File: provider/cmd/pulumi-resource-pinecone/main.go
package main

import (
	"fmt"
	"os"
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
	pinecone "github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone"
	"github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/version"
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
