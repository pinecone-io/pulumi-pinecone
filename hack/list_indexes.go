// list_index.go
// Action: (GET) list_indexes
// about: This operation returns a list of your Pinecone indexes.
// API Base URL: https://controller.us-east1-gcp.pinecone.io/databases
// Documentation URL: https://docs.pinecone.io/reference/list_indexes
// Response:
// - 200: (array of strings) This operation returns a list of all the indexes that you have previously created, and which are associated with the given API key.

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
	// Validate the API token is set.
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

	// Create a new HTTP GET request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add necessary headers to the request.
	req.Header.Add("Api-Key", apiToken)
	req.Header.Add("accept", "application/json; charset=utf-8")

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
