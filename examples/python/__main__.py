import pulumi
import pulumi_pinecone as pinecone

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
