// Action: (GET) index_describe
// About:
//   This operation returns information about a Pinecone index, including configuration information and deployment status.
// API Base URL: https://controller.gcp-starter.pinecone.io/databases/{indexName}
// Documentation URL: https://docs.pinecone.io/reference/describe_index
// Path Parameters:
// - indexName: (string) required: The name of the index to be described.
// Responses:
// - 200: Success. Returns the configuration information and deployment status of the index.
//   Response Body Fields:
//   - database: (object) Contains information about the index.
//     - name: (string) The name of the index.
//     - dimension: (integer) The dimensions of the vectors in the index.
//     - metric: (string) The distance metric used for similarity search.
//     - pods: (integer) The number of pods allocated to the index.
//     - replicas: (integer) The number of replicas of the index.
//     - pod_type: (string) The type of pod used for the index.
//     - metadata_config: (object) Configuration for Pinecone's internal metadata index.
//   - status: (object) Provides the deployment status of the index.
//     - ready: (boolean) Indicates if the index is ready.
//     - state: (string) The current state of the index (Initializing ScalingUp ScalingDown ScalingUpPodSize ScalingDownPodSize Terminating Ready InitializationFailed).
//     - host: (string) The host address of the index.
//     - port: (integer) The port number of the index.
