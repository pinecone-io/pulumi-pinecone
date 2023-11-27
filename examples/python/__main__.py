import pulumi

pulumi.export("output", {
    "value": my_pinecone_index["result"],
})
