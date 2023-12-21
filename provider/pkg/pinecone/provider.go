// File: provider/pkg/pinecone/provider.go
package provider

//go:generate oapi-codegen -generate types,client -o ./client/pinecone.gen.go -package client ./swagger/pinecone-swagger.yaml

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone/config"
	"github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone/index"
)

func Provider() p.Provider {
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[*index.PineconeIndex, index.PineconeIndexArgs, index.PineconeIndexState](),
			infer.Resource[*index.PineconeCollection, index.PineconeCollectionArgs, index.PineconeCollectionState](),
		},
		Functions: []infer.InferredFunction{
			infer.Function[*index.LookupPineconeIndex, index.LookupPineconeIndexArgs, index.LookupPineconeIndexResult](),
			infer.Function[*index.LookupPineconeCollection, index.LookupPineconeCollectionArgs, index.LookupPineconeCollectionResult](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"pinecone": "index",
		},
		Metadata: schema.Metadata{
			DisplayName: "Pinecone",
			Description: "A Pulumi native provider for Pinecone",
			Keywords: []string{
				"pulumi",
				"command",
				"category/utility",
				"kind/native",
			},
			Homepage:          "https://pulumi.com",
			License:           "Apache-2.0",
			Repository:        "",
			PluginDownloadURL: "github://api.github.com/usrbinkat",
			Publisher:         "Pulumi",
			LogoURL:           "",
			LanguageMap: map[string]any{
				"go": map[string]any{
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/pulumi/pulumi-pinecone/sdk/go/pinecone",
				},
				"csharp": map[string]any{
					"packageReferences": map[string]string{
						"Pulumi": "3.*",
					},
					"rootNamespace": "Pulumi",
				},
				"nodejs": map[string]any{
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
					"packageName": "@pulumi/pinecone",
				},
				"python": map[string]any{
					"requires": map[string]string{
						"pulumi": ">=3.0.0,<4.0.0",
					},
					"packageName": "pulumi_pinecone",
				},
			},
		},
		Config: infer.Config[*config.PineconeProviderConfig](),
	})
}
