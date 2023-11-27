// File: provider/cmd/pulumi-resource-pinecone/main.go
package main

import (
	p "github.com/pulumi/pulumi-go-provider"
	pinecone "github.com/usrbinkat/pulumi-pinecone-native/provider"
)

// Serve the provider against Pulumi's Provider protocol.
func main() {
	p.RunProvider(pinecone.Name, pinecone.Version, pinecone.Provider())
}
