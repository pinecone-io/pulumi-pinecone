package main

import (
	"github.com/pinecone-io/pulumi-pinecone/sdk/go/pinecone"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		myPineconeIndex, err := pinecone.NewPineconeIndex(ctx, "myPineconeIndex", &pinecone.PineconeIndexArgs{
			Name:   pulumi.String("example-index2"),
			Metric: pinecone.IndexMetricCosine,
			Spec: &pinecone.PineconeSpecArgs{
				Serverless: &pinecone.PineconeServerlessSpecArgs{
					Cloud:  pinecone.ServerlessSpecCloudAws,
					Region: pulumi.String("us-west-2"),
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("myPineconeIndex", myPineconeIndex.Name)
		ctx.Export("myPineconeIndexHost", myPineconeIndex.Host)

		myPineconeIndex2, err := pinecone.NewPineconeIndex(ctx, "myPineconeIndex2", &pinecone.PineconeIndexArgs{
			Name:      pulumi.String("example-index3"),
			Dimension: pulumi.Int(1536),
			Metric:    pinecone.IndexMetricCosine,
			Spec: &pinecone.PineconeSpecArgs{
				Pod: &pinecone.PineconePodSpecArgs{
					Environment: pulumi.String("gcp-starter"),
					MetaDataConfig: &pinecone.MetaDataConfigArgs{
						Indexed: pulumi.StringArray{
							pulumi.String("genre"),
							pulumi.String("title"),
							pulumi.String("imdb_rating"),
						},
					},
					Shards:   pulumi.Int(1),
					Replicas: pulumi.Int(1),
					Pods:     pulumi.Int(1),
					PodType:  pulumi.String("p1.x1"),
					//SourceCollection: pulumi.String("movie-embeddings"),
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("myPineconeIndex2", myPineconeIndex2.Host)

		myPineconeIndex3, err := pinecone.NewPineconeIndex(ctx, "myPineconeIndex3", &pinecone.PineconeIndexArgs{
			Name:      pulumi.String("example-index4"),
			Dimension: pulumi.Int(1536),
			Metric:    pinecone.IndexMetricCosine,
			Spec: &pinecone.PineconeSpecArgs{
				Pod: &pinecone.PineconePodSpecArgs{
					Environment: pulumi.String("us-east-1-aws"),
					MetaDataConfig: &pinecone.MetaDataConfigArgs{
						Indexed: pulumi.StringArray{
							pulumi.String("genre"),
							pulumi.String("title"),
							pulumi.String("imdb_rating"),
						},
					},
					Shards:   pulumi.Int(1),
					Replicas: pulumi.Int(1),
					Pods:     pulumi.Int(1),
					PodType:  pulumi.String("p1.x1"),
					//SourceCollection: pulumi.String("movie-embeddings"),
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("myPineconeIndex3", myPineconeIndex3.Host)

		return nil
	})
}
