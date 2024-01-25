package index

import (
	"fmt"
	"net/http"

	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone/client"
	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone/config"
	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone/utils"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

type LookupPineconeCollection struct{}

func (g *LookupPineconeCollection) Annotate(a infer.Annotator) {
	a.Describe(&g, "The result of a get operation on a Pinecone collection.")
}

func (*LookupPineconeCollection) Call(ctx p.Context, args LookupPineconeCollectionArgs) (LookupPineconeCollectionResult, error) {
	pineconeConfig := infer.GetConfig[config.PineconeProviderConfig](ctx)
	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			APIKey:    pineconeConfig.APIKey,
		},
	}
	pineconeClient, err := client.NewClientWithResponses("https://api.pinecone.io", client.WithHTTPClient(httpClient))
	if err != nil {
		return LookupPineconeCollectionResult{}, err
	}
	resp, err := pineconeClient.DescribeCollectionWithResponse(ctx, args.CollectionName)
	ctx.Logf(diag.Debug, "DescribeCollectionWithResponse: %v", resp.Status())
	if err != nil {
		ctx.Logf(diag.Error, "DescribeCollectionWithResponse: %v", resp.Status())
		return LookupPineconeCollectionResult{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return LookupPineconeCollectionResult{}, fmt.Errorf("DescribeCollectionWithResponse: %v", resp.Status())
	}
	return LookupPineconeCollectionResult{
		PineconeCollectionState: PineconeCollectionState{
			PineconeCollectionArgs: PineconeCollectionArgs{
				CollectionName: resp.JSON200.Name,
			},
			CollectionSize:        resp.JSON200.Size,
			CollectionDimension:   resp.JSON200.Dimension,
			CollectionVectorCount: resp.JSON200.VectorCount,
			CollectionEnvironment: resp.JSON200.Environment,
		},
	}, nil
}

type LookupPineconeCollectionArgs struct {
	CollectionName string `pulumi:"name"`
}

func (g *LookupPineconeCollectionArgs) Annotate(a infer.Annotator) {
	a.Describe(&g.CollectionName, "The name of the Pinecone collection.")
}

type LookupPineconeCollectionResult struct {
	PineconeCollectionState
}

func (g *LookupPineconeCollectionResult) Annotate(a infer.Annotator) {
	a.Describe(&g.PineconeCollectionArgs, "The result of a get operation on a Pinecone collection.")
}
