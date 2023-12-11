// File: provider/pkg/pinecone/index/types.go
package index

import (
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type PineconeIndex struct{}

type PineconeIndexArgs struct {
	IndexName      string `pulumi:"name"`
	IndexDimension int    `pulumi:"dimension"`
	IndexMetric    string `pulumi:"metric"`
	IndexPods      int    `pulumi:"pods"`
	IndexReplicas  int    `pulumi:"replicas"`
	IndexPodType   string `pulumi:"podType"`
}

type PineconeProviderConfig struct {
	APIToken    string `pulumi:"apiToken" provider:"secret"`
	PineconeEnv string `pulumi:"pineconeEnv"`
	Name        string `pulumi:"name"`
}

type PineconeIndexState struct {
	IndexName string `pulumi:"indexName"`
}

func (c *PineconeProviderConfig) Annotate(a infer.Annotator) {
	a.Describe(&c.APIToken, "The API token for Pinecone.")
	a.Describe(&c.PineconeEnv, "The environment for the Pinecone API.")
}

func (c *PineconeProviderConfig) Configure(ctx p.Context) error {
	if c.APIToken == "" {
		return fmt.Errorf("API token is required")
	}
	if c.PineconeEnv == "" {
		return fmt.Errorf("Pinecone environment is required")
	}
	return nil
}
