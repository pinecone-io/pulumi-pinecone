import pulumi
import pulumi_pinecone as pinecone

my_pinecone_index = pinecone.PineconeIndexResource("myPineconeIndex",
    name="example-index",
    dimension=512,
    metric="cosine",
    pods=1,
    replicas=1,
    pod_type="p1.x1")
pulumi.export("output", {
    "value": my_pinecone_index.id,
})
