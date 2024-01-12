---
title: Pinecone
meta_desc: Provides an overview of the Pinecone Provider for Pulumi.
layout: package
---

This Pulumi Pinecone Provider enables you to manage your [Pinecone](https://www.pinecone.io/) collections and indexes using any language of Pulumi Infrastructure as Code.

## Example

{{< chooser language "javascript,typescript,python,go" >}}


{{% choosable language javascript %}}

```javascript
"use strict";
const pulumi = require("@pulumi/pulumi");
const pinecone = require("@pinecone-database/pulumi");

const entity = new pinecone.PineconeIndex("myPineconeIndex", {
    name: "example-index",
    metric: pinecone.IndexMetric.Cosine,
    spec: {
        serverless: {
            cloud: pinecone.ServerlessSpecCloud.Aws,
            region: "us-west-2",
        },
    },
});

exports.host = myPineconeIndex.host;
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as pinecone from "@pinecone-database/pulumi";

const myPineconeIndex = new pinecone.PineconeIndex("myPineconeIndex", {
    name: "example-index",
    metric: pinecone.IndexMetric.Cosine,
    spec: {
        serverless: {
            cloud: pinecone.ServerlessSpecCloud.Aws,
            region: "us-west-2",
        },
    },
});
export const output = {
    value: myPineconeIndex.host,
};
```

{{% /choosable %}}

{{% choosable language python %}}

```python
"""A Python Pulumi program"""
import pulumi
import pinecone_pulumi as pinecone

my_pinecone_index = pinecone.PineconeIndex("myPineconeIndex",
   name="example-index",
   metric=pinecone.IndexMetric.COSINE,
   spec=pinecone.PineconeSpecArgs(
       serverless=pinecone.PineconeServerlessSpecArgs(
           cloud=pinecone.ServerlessSpecCloud.AWS,
           region="us-west-2",
       ),
   ))
pulumi.export("output", {
    "value": my_pinecone_index.host,
})
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pinecone-io/pulumi-pinecone/sdk/go/pinecone"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		myPineconeIndex, err := pinecone.NewPineconeIndex(ctx, "myPineconeIndex", &pinecone.PineconeIndexArgs{
			Name:      pulumi.String("example-index2"),
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
```

{{% /choosable %}}

{{< /chooser >}}
