import * as pulumi from "@pulumi/pulumi";
import * as pinecone from "@pinecone-database/pulumi";

const myPineconeIndex = new pinecone.Index("myPineconeIndex", {
    name: "example-index",
    dimension: 10,
    spec: {
        serverless: {
            cloud: "aws",
            region: "us-east-1",
        },
    },
});
export const name = myPineconeIndex.name;
export const host = myPineconeIndex.host;
