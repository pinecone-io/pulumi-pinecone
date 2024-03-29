// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PineconeDatabase.Pinecone
{
    [PineconeResourceType("pinecone:index:PineconeCollection")]
    public partial class PineconeCollection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The dimension of the vectors stored in each record held in the collection.
        /// </summary>
        [Output("dimension")]
        public Output<int> Dimension { get; private set; } = null!;

        /// <summary>
        /// The environment where the collection is hosted.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// The name of the collection to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The size of the collection in bytes.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The name of the index to be used as the source for the collection.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// The number of records stored in the collection.
        /// </summary>
        [Output("vectorCount")]
        public Output<int> VectorCount { get; private set; } = null!;


        /// <summary>
        /// Create a PineconeCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PineconeCollection(string name, PineconeCollectionArgs args, CustomResourceOptions? options = null)
            : base("pinecone:index:PineconeCollection", name, args ?? new PineconeCollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PineconeCollection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("pinecone:index:PineconeCollection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pinecone-io/pulumi-pinecone",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PineconeCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PineconeCollection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PineconeCollection(name, id, options);
        }
    }

    public sealed class PineconeCollectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the collection to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the index to be used as the source for the collection.
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        public PineconeCollectionArgs()
        {
        }
        public static new PineconeCollectionArgs Empty => new PineconeCollectionArgs();
    }
}
