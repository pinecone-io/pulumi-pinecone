// Action: (POST) create_collection
// About: This operation creates a Pinecone collection. Not supported by projects on the gcp-starter environment.
// API Base URL: https://controller.us-east1-gcp.pinecone.io/collections
// Documentation URL: https://docs.pinecone.io/reference/create_collection
// Parameters:
// - name: string (required) The name of the collection to be created.
// - source: string (required) The name of the source index to be used as the source for the collection.
// Response:
// - 201: (string) The collection has been successfully created.
// - 400: Bad request. Request exceeds quota or collection name is invalid.
// - 409: A collection with the name provided already exists.
// - 500: Internal error. Can be caused by invalid parameters.

package collection

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://controller.us-east1-gcp.pinecone.io/collections"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("accept", "text/plain")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}
