package config

import (
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"os"
)

type PineconeProviderConfig struct {
	APIKey string `pulumi:"APIKey,optional" provider:"secret"`
}

func (c *PineconeProviderConfig) Annotate(a infer.Annotator) {
	a.Describe(&c.APIKey, "The API token for Pinecone.")
}

func (c *PineconeProviderConfig) Configure(ctx p.Context) error {
	ctx.Logf(diag.Debug, "Configuring Pinecone provider")
	if c.APIKey == "" {
		APIKey, exists := os.LookupEnv("PINECONE_API_KEY")
		if exists {
			c.APIKey = APIKey
			return nil
		}
		return fmt.Errorf("API key is required")
	}
	return nil
}
