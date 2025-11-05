package main

import (
	"github.com/pinecone-io/pulumi-pinecone/sdk/v2/go/pinecone"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		myPineconeIndex, err := pinecone.NewIndex(ctx, "myPineconeIndex", &pinecone.IndexArgs{
			Name:      pulumi.String("example-index2"),
			Dimension: pulumi.Int(10),
			Spec: pinecone.IndexSpecArgs{
				Serverless: pinecone.IndexSpecServerlessArgs{
					Cloud:  pulumi.String("aws"),
					Region: pulumi.String("us-east-1"),
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("name", myPineconeIndex.Name)
		ctx.Export("host", myPineconeIndex.Host)

		return nil
	})
}
