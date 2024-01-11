package index

import (
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone/client"
	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone/config"
	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone/utils"
	"net/http"
)

type LookupPineconeIndex struct{}

func (g *LookupPineconeIndex) Annotate(a infer.Annotator) {
	a.Describe(&g, "The result of a get operation on a Pinecone index.")
}

func (*LookupPineconeIndex) Call(ctx p.Context, args LookupPineconeIndexArgs) (LookupPineconeIndexResult, error) {
	pineconeConfig := infer.GetConfig[config.PineconeProviderConfig](ctx)
	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			APIKey:    pineconeConfig.APIKey,
		},
	}
	pineconeClient, err := client.NewClientWithResponses("https://api.pinecone.io", client.WithHTTPClient(httpClient))
	if err != nil {
		return LookupPineconeIndexResult{}, err
	}
	resp, err := pineconeClient.DescribeIndexWithResponse(ctx, args.IndexName)
	ctx.Logf(diag.Debug, "DescribeIndexWithResponse: %v", resp.Status())
	if err != nil {
		ctx.Logf(diag.Error, "DescribeIndexWithResponse: %v", resp.Status())
		return LookupPineconeIndexResult{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return LookupPineconeIndexResult{}, fmt.Errorf("DescribeIndexWithResponse: %v", resp.Status())
	}
	return LookupPineconeIndexResult{
		PineconeIndexArgs: PineconeIndexArgs{
			IndexName:      resp.JSON200.Name,
			IndexMetric:    IndexMetric(resp.JSON200.Metric),
			IndexDimension: resp.JSON200.Dimension,
			IndexSpec: PineconeSpec{
				Serverless: PineconeServerlessSpec{
					Cloud:  ServerlessSpecCloud(resp.JSON200.Spec.Serverless.Cloud),
					Region: resp.JSON200.Spec.Serverless.Region,
				},
			},
		},
		IndextStatus: resp.JSON200.Status.Ready,
		IndexHost:    resp.JSON200.Host,
	}, nil
}

type LookupPineconeIndexArgs struct {
	IndexName string `pulumi:"name"`
}

func (g *LookupPineconeIndexArgs) Annotate(a infer.Annotator) {
	a.Describe(&g.IndexName, "The name of the Pinecone index.")
}

type LookupPineconeIndexResult struct {
	PineconeIndexArgs
	IndexHost    string `pulumi:"host,omitempty"`
	IndextStatus bool   `pulumi:"status,omitempty"`
}

func (g *LookupPineconeIndexResult) Annotate(a infer.Annotator) {
	a.Describe(&g, "The result of a get operation on a Pinecone index.")
}
