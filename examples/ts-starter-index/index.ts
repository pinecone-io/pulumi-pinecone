import * as pinecone from "@pinecone-database/pulumi";

const myPineconeIndex = new pinecone.PineconeIndex("myPineconeIndex", {
    name: "example-index",
    metric: pinecone.IndexMetric.Cosine,
    spec: {
        pod: {
            podType: "starter",
            replicas: 1,
            pods: 1,
            environment: "gcp-starter",
        }
    },
    dimension: 1536,
});
export const output = {
    value: myPineconeIndex.host,
};
