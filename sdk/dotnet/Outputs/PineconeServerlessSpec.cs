// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PineconeDatabase.Pinecone.Outputs
{

    [OutputType]
    public sealed class PineconeServerlessSpec
    {
        /// <summary>
        /// The public cloud where you would like your index hosted.
        /// </summary>
        public readonly PineconeDatabase.Pinecone.ServerlessSpecCloud Cloud;
        /// <summary>
        /// The region where you would like your index to be created. Different cloud providers have different regions available.
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private PineconeServerlessSpec(
            PineconeDatabase.Pinecone.ServerlessSpecCloud cloud,

            string region)
        {
            Cloud = cloud;
            Region = region;
        }
    }
}
