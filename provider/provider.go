// File: provider/provider.go
package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

const (
	Name    = "pinecone"
	Version = "0.1.0"

	diagInfo       = "info"
	controllerBase = "https://controller.%s.pinecone.io/databases"
)

// PineconeProviderConfig represents the configuration for the Pinecone provider
type PineconeProviderConfig struct {
	APIToken    string `pulumi:"apiToken" provider:"secret"`
	PineconeEnv string `pulumi:"pineconeEnv"`
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

// PineconeIndex represents a Pinecone index resource
type PineconeIndex struct{}

type PineconeIndexArgs struct {
	IndexName      string `pulumi:"name"`
	IndexDimension int    `pulumi:"dimension"`
	IndexMetric    string `pulumi:"metric"`
	IndexPods      int    `pulumi:"pods"`
	IndexReplicas  int    `pulumi:"replicas"`
	IndexPodType   string `pulumi:"podType"`
}

func (pia *PineconeIndexArgs) Annotate(a infer.Annotator) {
	a.Describe(&pia.IndexName, "The name of the Pinecone index.")
	a.Describe(&pia.IndexDimension, "The dimensions of the vectors in the index.")
	a.Describe(&pia.IndexMetric, "The metric used to compute the distance between vectors.")
	a.Describe(&pia.IndexPods, "The number of pods to use for the index.")
	a.Describe(&pia.IndexReplicas, "The number of replicas to use for the index.")
	a.Describe(&pia.IndexPodType, "The type of pods to use for the index.")
}

type PineconeIndexState struct {
	IndexName string `pulumi:"indexName"`
}

// Create a new Pinecone index
func (*PineconeIndex) Create(ctx p.Context, name string, args PineconeIndexArgs, preview bool) (string, PineconeIndexState, error) {
	config := infer.GetConfig[PineconeProviderConfig](ctx)

	if err := validateIndexName(args.IndexName); err != nil {
		return "", PineconeIndexState{}, fmt.Errorf("invalid index name: %w", err)
	}

	if preview {
		ctx.Logf(diagInfo, "Previewing Pinecone index creation: %s", args.IndexName)
		return name, PineconeIndexState{IndexName: args.IndexName}, nil
	}

	ctx.Logf(diagInfo, "Creating Pinecone index: %s", args.IndexName)
	response, err := createPineconeIndex(config, args)
	if err != nil {
		return "", PineconeIndexState{}, fmt.Errorf("failed to create Pinecone index: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", PineconeIndexState{}, fmt.Errorf("HTTP error code: %s", response.Status)
	}

	return args.IndexName, PineconeIndexState{IndexName: args.IndexName}, nil
}

// validateIndexName checks if the index name is valid
func validateIndexName(name string) error {
	matched, err := regexp.MatchString("^[a-z0-9-]+$", name)
	if err != nil {
		return fmt.Errorf("regex error: %w", err)
	}
	if !matched {
		return fmt.Errorf("index name can only contain lower case alphanumeric characters and dashes")
	}
	return nil
}

// createPineconeIndex creates a new Pinecone index
func createPineconeIndex(config PineconeProviderConfig, args PineconeIndexArgs) (*http.Response, error) {
	url := fmt.Sprintf(controllerBase, config.PineconeEnv)

	indexData := map[string]interface{}{
		"name":      args.IndexName,
		"dimension": args.IndexDimension,
		"metric":    args.IndexMetric,
		"pods":      args.IndexPods,
		"replicas":  args.IndexReplicas,
		"pod_type":  args.IndexPodType,
	}

	jsonData, err := json.Marshal(indexData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling index data: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Api-Key", config.APIToken)
	req.Header.Add("accept", "text/plain")
	req.Header.Add("content-type", "application/json")

	return http.DefaultClient.Do(req)
}

func Provider() p.Provider {
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{infer.Resource[*PineconeIndex, PineconeIndexArgs, PineconeIndexState]()},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
		Config: infer.Config[PineconeProviderConfig](),
	})
}
