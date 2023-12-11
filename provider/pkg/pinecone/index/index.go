// File: provider/pkg/pinecone/index/index.go
package index

import (
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone/client"

	utils "github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone/utils"
)

type PineconeIndex struct{}

type PineconeProviderConfig struct {
	APIToken    string `pulumi:"apiToken" provider:"secret"`
	PineconeEnv string `pulumi:"pineconeEnv"`
	Name        string `pulumi:"name"`
}

type PineconeIndexState struct {
	IndexName string `pulumi:"indexName"`
	IndexDimension int actualIndex.Dimension,
	IndexMetric int    actualIndex.Metric,
	IndexPods int      actualIndex.Pods,
	IndexReplicas int  actualIndex.Replicas,
	IndexPodType string   actualIndex.PodType,
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

const diagInfo = "info"

func (pia *PineconeIndexArgs) Annotate(a infer.Annotator) {
	a.Describe(&pia.IndexName, "The name of the Pinecone index.")
	a.Describe(&pia.IndexDimension, "The dimensions of the vectors in the index.")
	a.Describe(&pia.IndexMetric, "The metric used to compute the distance between vectors.")
	a.Describe(&pia.IndexPods, "The number of pods to use for the index.")
	a.Describe(&pia.IndexReplicas, "The number of replicas to use for the index.")
	a.Describe(&pia.IndexPodType, "The type of pods to use for the index.")
}

func (*PineconeIndex) Create(ctx p.Context, name string, args PineconeIndexArgs, preview bool) (string, PineconeIndexState, error) {
	config := infer.GetConfig[PineconeProviderConfig](ctx)
	if err := utils.ValidateIndexName(args.IndexName); err != nil {
		return "", PineconeIndexState{}, fmt.Errorf("invalid index name: %w", err)
	}

	if preview {
		ctx.Logf(diagInfo, "Previewing Pinecone index creation: %s", args.IndexName)
		return name, PineconeIndexState{IndexName: args.IndexName}, nil
	}

	// Create a new instance of PineconeClient
	pineconeClient := client.NewPineconeClient(config.APIToken, config.PineconeEnv)

	ctx.Logf(diagInfo, "Creating Pinecone index: %s", args.IndexName)
	response, err := pineconeClient.CreateIndex(args)
	if err != nil {
		return "", PineconeIndexState{}, fmt.Errorf("failed to create Pinecone index: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", PineconeIndexState{}, fmt.Errorf("HTTP error code: %s", response.Status)
	}

	// Update state based on actual created index
	if !preview {
		actualIndex, err := pineconeClient.DescribeIndex(args.IndexName)
		if err != nil {
			return "", PineconeIndexState{}, err
		}

		return args.IndexName, PineconeIndexState{
			IndexName:      actualIndex.Name,
			IndexDimension: actualIndex.Dimension,
			IndexMetric:    actualIndex.Metric,
			IndexPods:      actualIndex.Pods,
			IndexReplicas:  actualIndex.Replicas,
			IndexPodType:   actualIndex.PodType,
		}, nil
	}
	return args.IndexName, PineconeIndexState{IndexName: args.IndexName}, nil
}

func (pi *PineconeIndex) Delete(ctx p.Context, id string, state PineconeIndexState) error {
	// Retrieve the configuration from the context
	config := infer.GetConfig[PineconeProviderConfig](ctx)

	// Create a new Pinecone client instance
	pineconeClient := client.NewPineconeClient(config.APIToken, config.PineconeEnv)

	// Call the DeleteIndex method with the index name
	_, err := pineconeClient.DeleteIndex(state.IndexName)
	if err != nil {
		return fmt.Errorf("error deleting Pinecone index '%s': %w", state.IndexName, err)
	}

	// Log the successful deletion - Corrected logging statement
	ctx.Logf(diagInfo, "Successfully deleted Pinecone index: %s", state.IndexName)
	return nil
}

func (pi *PineconeIndex) Read(ctx p.Context, id string, state *PineconeIndexState) error {
	config := infer.GetConfig[PineconeProviderConfig](ctx)
	pineconeClient := client.NewPineconeClient(config.APIToken, config.PineconeEnv)

	// Call DescribeIndex to get the current state of the index
	indexDetails, err := pineconeClient.DescribeIndex(state.IndexName)
	if err != nil {
		if err.Error() == "index not found" {
			ctx.Logf(diagInfo, "Pinecone index '%s' not found, marking as deleted", state.IndexName)
			return nil
		}
		return fmt.Errorf("error getting Pinecone index details: %w", err)
	}

	// Update state with actual index details
	state.IndexName = indexDetails.Name
	state.IndexDimension = indexDetails.Dimension
	state.IndexMetric = indexDetails.Metric
	state.IndexPods = indexDetails.Pods
	state.IndexReplicas = indexDetails.Replicas
	state.IndexPodType = indexDetails.PodType

	return nil
}
