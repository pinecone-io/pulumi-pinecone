package shim

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/version"
	pinecone "github.com/pinecone-io/terraform-provider-pinecone/pinecone/provider"
)

func NewProvider() provider.Provider {
	p := pinecone.New(version.Version)
	return p()
}
