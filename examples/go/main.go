package main

import (
	"github.com/pulumi/pulumi-pinecone/sdk/go/pinecone"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myPineconeIndex, err := pinecone.NewPineconeIndexResource(ctx, "myPineconeIndex", &pinecone.PineconeIndexResourceArgs{
			Name:      pulumi.String("example-index"),
			Dimension: pulumi.Int(512),
			Metric:    pulumi.String("cosine"),
			Pods:      pulumi.Int(1),
			Replicas:  pulumi.Int(1),
			PodType:   pulumi.String("p1.x1"),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", map[string]interface{}{
			"value": myPineconeIndex.ID(),
		})
		return nil
	})
}
