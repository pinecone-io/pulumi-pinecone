import pulumi
import pinecone_pulumi as pinecone

my_pinecone_index = pinecone.Index("myPineconeIndex",
    name="example-index",
    dimension=10,
    spec={
        "serverless": {
            "cloud": "aws",
            "region": "us-east-1",
        },
    })
pulumi.export("name", my_pinecone_index.name)
pulumi.export("host", my_pinecone_index.host)
