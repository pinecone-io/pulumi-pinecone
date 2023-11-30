using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Pinecone = Pulumi.Pinecone;

return await Deployment.RunAsync(() => 
{
    var myPineconeIndex = new Pinecone.PineconeIndexResource("myPineconeIndex", new()
    {
        Name = "example-index",
        Dimension = 512,
        Metric = "cosine",
        Pods = 1,
        Replicas = 1,
        PodType = "p1.x1",
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myPineconeIndex.Id },
        },
    };
});

