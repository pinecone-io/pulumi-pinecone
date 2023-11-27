// Action: (DELETE) delete_index
// About: This operation deletes an existing index.
// API Base URL: https://controller.us-east1-gcp.pinecone.io/databases/{indexName}
// Documentation URL: https://docs.pinecone.io/reference/delete_index
// Params: (required)
// - indexName: path param (string) required: The name of the index.
// Response:
// - 202: (string) The index has been successfully deleted.
// - 404: Index not found.
// - 500: Internal error. Can be caused by invalid parameters.

package main

import (
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
	indexName := "magic-index" // Replace with the desired index name

	// Construct the API endpoint URL using the Pinecone environment and index name.
	url := fmt.Sprintf("https://controller.%s.pinecone.io/databases/%s", pineconeEnv, indexName)

	// Create a new HTTP DELETE request.
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add necessary headers to the request.
	req.Header.Add("Api-Key", apiToken)
	req.Header.Add("accept", "text/plain")

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
