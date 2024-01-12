package main

import (
	"github.com/pinecone-io/pulumi-pinecone/sdk/go/pinecone"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

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

		return nil
	})
}
