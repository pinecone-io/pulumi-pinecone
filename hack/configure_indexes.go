// configure_index.go
// Action: (PATCH) configure_index
// About: This operation specifies the pod type and number of replicas for an index. Not supported by projects on the gcp-starter environment.
// API Base URL: https://controller.us-east1-gcp.pinecone.io/databases/{indexName}
// Documentation URL: https://docs.pinecone.io/reference/configure_index
// Parameters:
// - URL Path:
//   - indexName (string) required: The name of the index.
// - Body:
//   - replicas (integer) The desired number of replicas for the index.
//   - pod_type (string) The new pod type for the index. One of s1, p1, or p2 appended with . and one of x1, x2, x4, or x8.
// Responses:
// - 202 (string) The index has been successfully updated.
// - 400 Bad request, not enough quota.
// - 404 Index not found.
// - 500 Internal error. Can be caused by invalid parameters.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
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

	// Set the index name.
	indexName := "your-index-name" // Replace with the desired index name

	// Construct the API endpoint URL using the Pinecone environment and index name.
	url := fmt.Sprintf("https://controller.%s.pinecone.io/databases/%s", pineconeEnv, indexName)

	// Define the payload for configuring the index.
	// This can be modified to include other parameters as needed.
	configData := map[string]interface{}{
		"replicas": 2,       // Replace with the desired number of replicas
		"pod_type": "p1.x1", // Replace with the desired pod type
	}
	jsonData, err := json.Marshal(configData)
	if err != nil {
		fmt.Println("Error marshaling configuration data:", err)
		return
	}

	// Create a new HTTP PATCH request with the payload.
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
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
