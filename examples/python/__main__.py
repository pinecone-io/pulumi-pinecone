import pulumi
import pinecone_pulumi as pinecone

pinecone_index = pinecone.PineconeIndex(
    "my-index",
    name="my-index",
    spec=pinecone.PineconeSpecArgs(
        pod=pinecone.PineconePodSpecArgs(
            environment="gcp-starter",
            pod_type="starter",
            replicas=1,
        ),
    ),
    metric="cosine",
    dimension=1536,
) 
pulumi.export("output", {
    "value": pinecone_index.host,
})
