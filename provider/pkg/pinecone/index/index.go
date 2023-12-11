// File: provider/pkg/pinecone/index/index.go
package index

import (
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone/client"
	utils "github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone/utils"
)

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

	return args.IndexName, PineconeIndexState{IndexName: args.IndexName}, nil
}

// PineconeIndexArgs methods to satisfy the IndexData interface
func (args PineconeIndexArgs) GetIndexName() string {
	return args.IndexName
}

func (args PineconeIndexArgs) GetIndexDimension() int {
	return args.IndexDimension
}

func (args PineconeIndexArgs) GetIndexMetric() string {
	return args.IndexMetric
}

func (args PineconeIndexArgs) GetIndexPods() int {
	return args.IndexPods
}

func (args PineconeIndexArgs) GetIndexReplicas() int {
	return args.IndexReplicas
}

func (args PineconeIndexArgs) GetIndexPodType() string {
	return args.IndexPodType
}
