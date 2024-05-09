package config

import (
	"context"
	"fmt"
	goprovider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"os"
)

type PineconeProviderConfig struct {
	APIKey string `pulumi:"APIKey,optional" provider:"secret"`
}

func (c *PineconeProviderConfig) Annotate(a infer.Annotator) {
	a.Describe(&c.APIKey, "The API token for Pinecone.")
}

func (c *PineconeProviderConfig) Configure(ctx context.Context) error {
	goprovider.GetLogger(ctx).Debugf("Configuring Pinecone provider")
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
