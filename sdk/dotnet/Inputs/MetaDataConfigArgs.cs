// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PineconeDatabase.Pinecone.Inputs
{

    public sealed class MetaDataConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("indexed")]
        private InputList<string>? _indexed;

        /// <summary>
        ///  Indexed By default, all metadata is indexed; to change this behavior, use this property to specify an array of metadata fields which should be indexed.
        /// </summary>
        public InputList<string> Indexed
        {
            get => _indexed ?? (_indexed = new InputList<string>());
            set => _indexed = value;
        }

        public MetaDataConfigArgs()
        {
        }
        public static new MetaDataConfigArgs Empty => new MetaDataConfigArgs();
    }
}
