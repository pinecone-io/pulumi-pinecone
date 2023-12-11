// File: provider/pkg/pinecone/provider.go
package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone/index"
)

const (
	Name    = "pinecone"
	Version = "0.1.0"
)

func Provider() p.Provider {
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[*index.PineconeIndex, index.PineconeIndexArgs, index.PineconeIndexState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"index": "index",
		},
		Config: infer.Config[index.PineconeProviderConfig](),
		//Init: func(ctx p.Context, c index.PineconeProviderConfig) (any, error) {
		//	return client.NewPineconeClient(c.APIToken, c.PineconeEnv), nil
		//},
	})
}
