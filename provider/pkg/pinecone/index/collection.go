package index

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone/client"
	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone/config"
	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone/utils"
	goprovider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type PineconeCollection struct{}

type PineconeCollectionArgs struct {
	CollectionName   string `pulumi:"name"`
	CollectionSource string `pulumi:"source"`
}

func (pc *PineconeCollectionArgs) Annotate(a infer.Annotator) {
	a.Describe(&pc.CollectionName, "The name of the collection to be created.")
	a.Describe(&pc.CollectionSource, "The name of the index to be used as the source for the collection.")
}

func (*PineconeCollection) Create(ctx context.Context, name string, args PineconeCollectionArgs, preview bool) (string, PineconeCollectionState, error) {
	pineconeConfig := infer.GetConfig[config.PineconeProviderConfig](ctx)

	if preview {
		goprovider.GetLogger(ctx).Debugf("Creating Pinecone collection %s", args.CollectionName)
		return "", PineconeCollectionState{
			PineconeCollectionArgs: PineconeCollectionArgs{
				CollectionName:   args.CollectionName,
				CollectionSource: args.CollectionSource,
			},
		}, nil
	}
	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			APIKey:    pineconeConfig.APIKey,
		},
	}

	pineconeClient, err := client.NewClientWithResponses("https://api.pinecone.io", client.WithHTTPClient(httpClient))
	if err != nil {
		return "", PineconeCollectionState{}, fmt.Errorf("failed to create Pinecone client: %w", err)
	}
	goprovider.GetLogger(ctx).Debugf("Creating Pinecone collection %s", args.CollectionName)
	resp, err := pineconeClient.CreateCollectionWithResponse(ctx, client.CreateCollectionJSONRequestBody{
		Name:   args.CollectionName,
		Source: args.CollectionSource,
	})
	if err != nil {
		goprovider.GetLogger(ctx).Debugf("Failed to create Pinecone collection %s with http status code %d", args.CollectionName, resp.StatusCode())
		return "", PineconeCollectionState{}, fmt.Errorf("failed to create Pinecone collection: %w", err)
	}
	goprovider.GetLogger(ctx).Debugf("Pinecone collection creaation responese: %s", string(resp.Body))

	ready := false
	for !ready {
		resp, err := pineconeClient.DescribeCollectionWithResponse(ctx, args.CollectionName)
		if err != nil {
			return "", PineconeCollectionState{}, fmt.Errorf("failed to describe Pinecone collection: %w", err)
		}
		if resp.JSON200.Status == client.CollectionModelStatusReady {
			ready = true
		}
	}
	return args.CollectionName, PineconeCollectionState{
		PineconeCollectionArgs: PineconeCollectionArgs{
			CollectionName:   args.CollectionName,
			CollectionSource: args.CollectionSource,
		},
		CollectionDimension:   resp.JSON201.Dimension,
		CollectionSize:        resp.JSON201.Size,
		CollectionVectorCount: resp.JSON201.VectorCount,
		CollectionEnvironment: resp.JSON201.Environment,
	}, nil
}

func (*PineconeCollection) Delete(ctx context.Context, id string, args PineconeCollectionArgs) error {
	pineconeConfig := infer.GetConfig[config.PineconeProviderConfig](ctx)
	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			APIKey:    pineconeConfig.APIKey,
		},
	}
	pineconeClient, err := client.NewClientWithResponses("https://api.pinecone.io", client.WithHTTPClient(httpClient))
	if err != nil {
		return fmt.Errorf("failed to create Pinecone client: %w", err)
	}
	response, err := pineconeClient.DeleteCollectionWithResponse(ctx, args.CollectionName)
	if err != nil {
		goprovider.GetLogger(ctx).Errorf("Failed to delete Pinecone collection: %s with http status code: %d", args.CollectionName, response.StatusCode())
		return fmt.Errorf("failed to delete Pinecone collection: %w", err)
	}
	goprovider.GetLogger(ctx).Debugf("Successfully deleted Pinecone collection: %s", args.CollectionName)
	return nil
}

func (*PineconeCollection) Read(ctx context.Context, id string, args PineconeCollectionArgs, state PineconeCollectionState) (canonicalID string, normalizedInputs PineconeCollectionArgs, normalizedState PineconeCollectionState, err error) {
	pineconeConfig := infer.GetConfig[config.PineconeProviderConfig](ctx)
	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			APIKey:    pineconeConfig.APIKey,
		},
	}
	pineconeClient, err := client.NewClientWithResponses("https://api.pinecone.io", client.WithHTTPClient(httpClient))
	if err != nil {
		return id, args, state, fmt.Errorf("failed to create Pinecone client: %w", err)
	}

	resp, err := pineconeClient.DescribeCollectionWithResponse(ctx, args.CollectionName)
	if err != nil {
		if resp.JSON404 != nil {
			goprovider.GetLogger(ctx).Debugf("Pinecone collection '%s' not found", args.CollectionName)
			return id, args, state, nil
		}
		return id, args, state, fmt.Errorf("error getting Pinecone collection details '%s': %w", args.CollectionName, err)
	}

	state.CollectionName = resp.JSON200.Name
	state.CollectionDimension = resp.JSON200.Dimension
	state.CollectionSize = resp.JSON200.Size
	state.CollectionVectorCount = resp.JSON200.VectorCount
	state.CollectionEnvironment = resp.JSON200.Environment

	return id, args, state, nil
}

type PineconeCollectionState struct {
	PineconeCollectionArgs
	CollectionSize        *int64  `pulumi:"size"`
	CollectionDimension   *int32  `pulumi:"dimension"`
	CollectionVectorCount *int32  `pulumi:"vectorCount"`
	CollectionEnvironment *string `pulumi:"environment"`
}

func (pcs *PineconeCollectionState) Annotate(a infer.Annotator) {
	a.Describe(&pcs.CollectionSize, "The size of the collection in bytes.")
	a.Describe(&pcs.CollectionDimension, "The dimension of the vectors stored in each record held in the collection.")
	a.Describe(&pcs.CollectionVectorCount, "The number of records stored in the collection.")
	a.Describe(&pcs.CollectionEnvironment, "The environment where the collection is hosted.")
}
