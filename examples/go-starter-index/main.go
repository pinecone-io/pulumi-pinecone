package main

import (
	"github.com/pinecone-io/pulumi-pinecone/sdk/go/pinecone"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		myPineconeIndex, err := pinecone.NewPineconeIndex(ctx, "myPineconeIndex", &pinecone.PineconeIndexArgs{
			Name:      pulumi.String("example-index3"),
			Dimension: pulumi.Int(1536),
			Metric:    pinecone.IndexMetricCosine,
			Spec: &pinecone.PineconeSpecArgs{
				Pod: &pinecone.PineconePodSpecArgs{
					Environment: pulumi.String("gcp-starter"),
					Shards:      pulumi.Int(1),
					Replicas:    pulumi.Int(1),
					//Pods:        pulumi.Int(1),
					PodType: pulumi.String("starter"),
					//SourceCollection: pulumi.String("movie-embeddings"),
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("myPineconeIndex2", myPineconeIndex.Host)

		return nil
	})
}
