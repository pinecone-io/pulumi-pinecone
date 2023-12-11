import * as pulumi from "@pulumi/pulumi";
import * as pinecone from "@pulumi/pinecone";

const myPineconeIndex = new pinecone.PineconeIndex("myPineconeIndex", {
    name: "example-index",
    dimension: 512,
    metric: "cosine",
    pods: 1,
    replicas: 1,
    podType: "p1.x1",
});
export const output = {
    value: myPineconeIndex.id,
};
