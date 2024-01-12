import * as pulumi from "@pulumi/pulumi";
import * as pinecone from "@pinecone-database/pinecone";

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
