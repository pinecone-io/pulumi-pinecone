// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace PineconeDatabase.Pinecone
{
    [EnumType]
    public readonly struct IndexMetric : IEquatable<IndexMetric>
    {
        private readonly string _value;

        private IndexMetric(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IndexMetric Dotproduct { get; } = new IndexMetric("dotproduct");
        public static IndexMetric Cosine { get; } = new IndexMetric("cosine");
        public static IndexMetric Euclidean { get; } = new IndexMetric("euclidean");

        public static bool operator ==(IndexMetric left, IndexMetric right) => left.Equals(right);
        public static bool operator !=(IndexMetric left, IndexMetric right) => !left.Equals(right);

        public static explicit operator string(IndexMetric value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IndexMetric other && Equals(other);
        public bool Equals(IndexMetric other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServerlessSpecCloud : IEquatable<ServerlessSpecCloud>
    {
        private readonly string _value;

        private ServerlessSpecCloud(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServerlessSpecCloud Aws { get; } = new ServerlessSpecCloud("aws");
        public static ServerlessSpecCloud Azure { get; } = new ServerlessSpecCloud("azure");
        public static ServerlessSpecCloud Gcp { get; } = new ServerlessSpecCloud("gcp");

        public static bool operator ==(ServerlessSpecCloud left, ServerlessSpecCloud right) => left.Equals(right);
        public static bool operator !=(ServerlessSpecCloud left, ServerlessSpecCloud right) => !left.Equals(right);

        public static explicit operator string(ServerlessSpecCloud value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServerlessSpecCloud other && Equals(other);
        public bool Equals(ServerlessSpecCloud other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
