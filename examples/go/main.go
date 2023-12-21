package main

import (
	"github.com/pulumi/pulumi-pinecone/sdk/go/pinecone"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		cfg := config.New(ctx, "")

		myPineconeIndex, err := pinecone.NewPineconeIndex(ctx, "myPineconeIndex", &pinecone.PineconeIndexArgs{
			Name:      pulumi.String("example-index2"),
			Dimension: pulumi.Int(512),
			Metric:    pinecone.IndexMetricCosine,
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

		index, err := pinecone.LookupPineconeIndex(ctx, &pinecone.LookupPineconeIndexArgs{
			Name: "example-index2",
		})
		if err != nil {
			return err
		}
		ctx.Export("myPineconeIndexStatus", pulumi.Sprintf("%v", index.Status))

		podPineConeProvider, err := pinecone.NewProvider(ctx, "podPineConeProvider", &pinecone.ProviderArgs{
			APIKey: cfg.GetSecret("pod-api"),
		})

		myPineconeIndex2, err := pinecone.NewPineconeIndex(ctx, "myPineconeIndex2", &pinecone.PineconeIndexArgs{
			Name:      pulumi.String("example-index3"),
			Dimension: pulumi.Int(1536),
			Metric:    pinecone.IndexMetricCosine,
			Spec: &pinecone.PineconeSpecArgs{
				Pod: &pinecone.PineconePodSpecArgs{
					Environment: pulumi.String("dev"),
					MetaDataConfig: &pinecone.MetaDataConfigArgs{
						Indexed: pulumi.StringArray{
							pulumi.String("genre"),
							pulumi.String("title"),
							pulumi.String("imdb_rating"),
						},
					},
					//Shards:           pulumi.Int(1),
					Replicas:         pulumi.Int(1),
					Pods:             pulumi.Int(1),
					PodType:          pulumi.String("p1.x1"),
					SourceCollection: pulumi.String("movie-embeddings"),
				},
			},
		}, pulumi.Provider(podPineConeProvider))
		if err != nil {
			return err
		}
		ctx.Export("myPineconeIndex2", myPineconeIndex2.Name)
		return nil
	})
}
