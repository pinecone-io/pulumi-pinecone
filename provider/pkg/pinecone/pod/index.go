package pod

import (
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone/client"
	"github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone/config"
	"github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone/utils"
	"net/http"
)

type PineconePodIndex struct{}

type PodSpecPodType = string

type PodSpecReplicas = int32

type PodSpecShards = int32

type PineconePodSpec struct {
	Environment    string          `pulumi:"environment"`
	Replicas       PodSpecReplicas `pulumi:"replicas"`
	Shards         PodSpecShards   `pulumi:"shards"`
	PodType        PodSpecPodType  `pulumi:"podType"`
	Pods           int             `pulumi:"pods"`
	MetaDataConfig struct {
		Indexed *[]string `pulumi:"indexed"`
	} `pulumi:"metaDataConfig"`
	SourceCollection *string `pulumi:"sourceCollection"`
}

type PineconeSpec struct {
	Pod PineconePodSpec `pulumi:"pod,optional"`
}

type PineconePodIndexArgs struct {
	IndexName      string                `pulumi:"name"`
	IndexDimension client.IndexDimension `pulumi:"dimension"`
	IndexMetric    client.IndexMetric    `pulumi:"metric"`
	IndexSpec      PineconeSpec          `pulumi:"spec"`
}

func (pia *PineconePodIndexArgs) Annotate(a infer.Annotator) {
	a.Describe(&pia.IndexName, "The name of the Pinecone index.")
	a.Describe(&pia.IndexDimension, "The dimensions of the vectors in the index.")
	a.Describe(&pia.IndexMetric, "The metric used to compute the distance between vectors.")
	a.Describe(&pia.IndexSpec, "Describe how the index should be deployed.")
}

func (*PineconePodIndex) Create(ctx p.Context, name string, args PineconePodIndexArgs, preview bool) (string, PineconePodIndexState, error) {
	pineconeConfig := infer.GetConfig[config.PineconeProviderConfig](ctx)
	if err := utils.ValidateIndexName(args.IndexName); err != nil {
		return "", PineconePodIndexState{}, fmt.Errorf("invalid index name: %w", err)
	}

	if preview {
		ctx.Logf(diag.Debug, "Previewing Pinecone index creation: %s", args.IndexName)
		return args.IndexName, PineconePodIndexState{
			PineconePodIndexArgs: PineconePodIndexArgs{
				IndexName:      args.IndexName,
				IndexMetric:    args.IndexMetric,
				IndexDimension: args.IndexDimension,
				IndexSpec:      args.IndexSpec,
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
		return "", PineconePodIndexState{}, fmt.Errorf("failed to create Pinecone client: %w", err)
	}

	ctx.Logf(diag.Debug, "Creating Pinecone pod-based index %s", args.IndexName)
	resp, err := pineconeClient.CreateIndexWithResponse(ctx, client.CreateIndexJSONRequestBody{
		Name:      args.IndexName,
		Dimension: args.IndexDimension,
		Metric:    args.IndexMetric,
		Spec: client.CreateIndexRequest_Spec{
			Pod: &client.PodSpec{
				Environment: args.IndexSpec.Pod.Environment,
				MetadataConfig: &struct {
					Indexed *[]string `json:"indexed,omitempty"`
				}{
					Indexed: args.IndexSpec.Pod.MetaDataConfig.Indexed,
				},
				Pods:             args.IndexSpec.Pod.Pods,
				PodType:          args.IndexSpec.Pod.PodType,
				Replicas:         args.IndexSpec.Pod.Replicas,
				Shards:           args.IndexSpec.Pod.Shards,
				SourceCollection: args.IndexSpec.Pod.SourceCollection,
			},
		},
	})
	if err != nil {
		ctx.Logf(diag.Error, "Failed to create Pinecone pod-based index %s with http error code: %d", args.IndexName, resp.StatusCode())
		return "", PineconePodIndexState{}, fmt.Errorf("failed to create Pinecone pod-based index %w", err)
	}
	ctx.Logf(diag.Debug, "Pinecone pod-based index creation response: %s", string(resp.Body))
	ready := false
	for !ready {
		resp, err := pineconeClient.DescribeIndexWithResponse(ctx, args.IndexName)
		if err != nil {
			return "", PineconePodIndexState{}, fmt.Errorf("failed to describe Pinecone pod-based index: %w", err)
		}
		if resp.JSON200.Status.Ready {
			ready = true
		}
	}
	return args.IndexName, PineconePodIndexState{
		PineconePodIndexArgs: PineconePodIndexArgs{
			IndexName:      args.IndexName,
			IndexMetric:    args.IndexMetric,
			IndexDimension: args.IndexDimension,
			IndexSpec:      args.IndexSpec,
		},
		IndexHost: resp.JSON201.Host,
	}, nil
}

func (*PineconePodIndex) Delete(ctx p.Context, id string, state PineconePodIndexState) error {
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
	response, err := pineconeClient.DeleteIndexWithResponse(ctx, state.IndexName)
	if err != nil {
		ctx.Logf(diag.Error, "Failed to delete Pinecone pod-based index: %s with http status code: %d", state.IndexName, response.StatusCode())
		return fmt.Errorf("failed to delete Pinecone pod-based index: %w", err)
	}
	ctx.Logf(diag.Debug, "Pinecone pod-based index deletion response: %s", string(response.Body))
	return nil
}

func (*PineconePodIndex) Read(ctx p.Context, id string, args PineconePodIndexArgs, state PineconePodIndexState) (canonicalID string, normalizedInputs PineconePodIndexArgs, normalizedState PineconePodIndexState, err error) {
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
	indexDetails, err := pineconeClient.DescribeIndexWithResponse(ctx, state.IndexName)
	if err != nil {
		return id, args, state, fmt.Errorf("failed to describe Pinecone pod-based index: %w", err)
	}
	return state.IndexName, PineconePodIndexArgs{
			IndexName:      state.IndexName,
			IndexMetric:    state.IndexMetric,
			IndexDimension: state.IndexDimension,
			IndexSpec:      state.IndexSpec,
		}, PineconePodIndexState{
			PineconePodIndexArgs: PineconePodIndexArgs{
				IndexName:      state.IndexName,
				IndexMetric:    state.IndexMetric,
				IndexDimension: state.IndexDimension,
				IndexSpec:      state.IndexSpec,
			},
			IndexHost: indexDetails.JSON200.Host,
		}, nil
}

type PineconePodIndexState struct {
	PineconePodIndexArgs
	IndexHost string `pulumi:"host,omitempty"`
}

func (pia *PineconePodIndexState) Annotate(a infer.Annotator) {
	a.Describe(&pia.IndexName, "The name of the Pinecone index.")
	a.Describe(&pia.IndexDimension, "The dimensions of the vectors in the index.")
	a.Describe(&pia.IndexMetric, "The metric used to compute the distance between vectors.")
	a.Describe(&pia.IndexSpec, "Describe how the index should be deployed.")
}
