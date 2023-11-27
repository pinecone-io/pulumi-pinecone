// create_index.go
// Action: (POST) create_index
// About:
//   This operation creates a Pinecone index. You can use it to specify the measure of similarity, the dimension
//   of vectors to be stored in the index, the numbers of replicas to use, and more.
// API Base URL: https://controller.us-east1-gcp.pinecone.io/databases
// Documentation URL: https://docs.pinecone.io/reference/create_index
// Parameters:
// - name: (string) required: The name of the index to be created. The maximum length is 45 characters.
// - dimension: (integer) required: The dimensions of the vectors to be inserted in the index.
// - metric: (string) The distance metric to be used for similarity search. You can use 'euclidean', 'cosine', or 'dotproduct'.
// - pods: (integer) The number of pods for the index to use,including replicas.
// - replicas: (integer) The number of replicas. Replicas duplicate your index. They provide higher availability and throughput.
// - pod_type: (string) The type of pod to use. One of s1, p1, or p2 appended with . and one of x1, x2, x4, or x8.
// - metadata_config: (object | null) Configuration for the behavior of Pinecone's internal metadata index. By default, all metadata is indexed; when metadata_config is present, only specified metadata fields are indexed. To specify metadata fields to index, provide a JSON object of the following form: {"indexed": ["example_metadata_field"]}
// - source_collection: (string) The name of the collection to create an index from
// Responses:
// - 201: (string) The index has been successfully created
// - 400: Bad request. Encountered when request exceeds quota or an invalid index name.
// - 409: Index of given name already exists.
// - 500: Internal error. Can be caused by invalid parameters.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {

	// Set the index name.
	indexName := "magic-index"
	// Validate the index name.
	if err := validateIndexName(indexName); err != nil {
		fmt.Println("Invalid index name:", err)
		return
	}

	// Retrieve the API token from the environment variable.
	apiToken := os.Getenv("PC_API_TOKEN")
	if apiToken == "" {
		fmt.Println("API token not found in environment variables")
		return
	}

	// Retrieve the Pinecone environment from the environment variable.
	pineconeEnv := os.Getenv("PC_API_ENVIRONMENT")
	if pineconeEnv == "" {
		fmt.Println("Pinecone environment not found in environment variables")
		return
	}

	// Construct the API endpoint URL using the Pinecone environment.
	url := fmt.Sprintf("https://controller.%s.pinecone.io/databases", pineconeEnv)

	// Define the payload for creating the index.
	// This can be modified to include other parameters as needed.
	indexData := map[string]interface{}{
		"name":      indexName,
		"dimension": 512,
		"metric":    "cosine",
		"pods":      1,
		"replicas":  1,
		"pod_type":  "p1.x1",
	}
	jsonData, err := json.Marshal(indexData)
	if err != nil {
		fmt.Println("Error marshaling index data:", err)
		return
	}

	// Create a new HTTP POST request with the payload.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add necessary headers to the request.
	req.Header.Add("Api-Key", apiToken)
	req.Header.Add("accept", "text/plain")
	req.Header.Add("content-type", "application/json")

	// Send the request to the server.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	// Check if the HTTP response status code is in the 2xx range.
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		fmt.Println("Non-success status code received:", res.Status)
		body, _ := io.ReadAll(res.Body)
		fmt.Println("Response body:", string(body))
		return
	}

	// Read and print the response body.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}

// validateIndexName checks if the index name contains only alphanumeric characters and dashes.
func validateIndexName(name string) error {
	matched, err := regexp.MatchString("^[a-z0-9-]+$", name)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("index name can only contain lower case alphanumeric characters and dashes")
	}
	return nil
}
