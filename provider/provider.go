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
)

func Provider() p.Provider {
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{infer.Resource[*PineconeIndex, PineconeIndexArgs, PineconeIndexState]()},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"pinecone": "index",
		},
		Config: infer.Config[Config](),
	})
}

type Config struct {
	APIToken    string `pulumi:"apiToken" provider:"secret"`
	PineconeEnv string `pulumi:"pineconeEnv"`
}

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.APIToken, "The API token for Pinecone.")
	a.Describe(&c.PineconeEnv, "The environment for the Pinecone API.")
}

func (c *Config) Configure(ctx p.Context) error {
	if c.APIToken == "" {
		return fmt.Errorf("API token is required")
	}
	if c.PineconeEnv == "" {
		return fmt.Errorf("Pinecone environment is required")
	}
	return nil
}

type PineconeIndex struct{}

type PineconeIndexArgs struct {
	IndexName      string `pulumi:"indexName"`
	IndexDimension int    `pulumi:"indexDimension"`
	IndexMetric    string `pulumi:"indexMetric"`
	IndexPods      int    `pulumi:"indexPods"`
	IndexReplicas  int    `pulumi:"indexReplicas"`
	IndexPodType   string `pulumi:"indexPodType"`
}

func (pia *PineconeIndexArgs) Annotate(a infer.Annotator) {
	a.Describe(&pia.IndexName, "The name of the Pinecone index.")
	a.Describe(&pia.IndexDimension, "The dimensions of the vectors in the index.")
}

type PineconeIndexState struct {
	IndexName string `pulumi:"indexName"`
}

func (*PineconeIndex) Create(ctx p.Context, name string, args PineconeIndexArgs, preview bool) (string, PineconeIndexState, error) {
	config := infer.GetConfig[Config](ctx)

	// Validate index name
	if err := validateIndexName(args.IndexName); err != nil {
		return "", PineconeIndexState{}, err
	}

	// Construct API URL
	url := fmt.Sprintf("https://controller.%s.pinecone.io/databases", config.PineconeEnv)

	// Define index parameters.
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
		return "", PineconeIndexState{}, fmt.Errorf("error marshaling index data: %s", err)
	}

	// Create POST request with index params.
	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", PineconeIndexState{}, fmt.Errorf("error creating request: %s", err)
	}

	// Add headers to request.
	req.Header.Add("Api-Key", config.APIToken)
	req.Header.Add("accept", "text/plain")
	req.Header.Add("content-type", "application/json")

	// Send request.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", PineconeIndexState{}, fmt.Errorf("error sending request: %s", err)
	}
	defer res.Body.Close()

	// Check response status code.
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", PineconeIndexState{}, fmt.Errorf("HTTP Error code: %s", res.Status)
	}

	// Return state.
	return args.IndexName, PineconeIndexState{IndexName: args.IndexName}, nil
}

func validateIndexName(name string) error {
	matched, err := regexp.MatchString("^[a-z0-9-]+$", name)
	if err != nil {
		return fmt.Errorf("error validating index name: %s", err)
	}
	if !matched {
		return fmt.Errorf("index name can only contain lower case alphanumeric characters and dashes")
	}
	return nil
}
