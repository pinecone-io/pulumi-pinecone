using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Pinecone = PineconeDatabase.Pinecone;

return await Deployment.RunAsync(() => 
{
    var myPineconeIndex = new Pinecone.Index("myPineconeIndex", new()
    {
        Name = "example-index",
        Dimension = 10,
        Spec = new Pinecone.Inputs.IndexSpecArgs
        {
            Serverless = new Pinecone.Inputs.IndexSpecServerlessArgs
            {
                Cloud = "aws",
                Region = "us-east-1",
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["name"] = myPineconeIndex.Name,
        ["host"] = myPineconeIndex.Host,
    };
});

