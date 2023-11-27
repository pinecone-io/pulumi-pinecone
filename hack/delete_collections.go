// Action: (DELETE) delete_collection
// About: This operation deletes an existing collection.
// API Base URL: https://controller.us-east1-gcp.pinecone.io/collections/{collectionName}
// Documentation URL: https://docs.pinecone.io/reference/delete_collection
// Parameters:
// - URL Path:
//   - collectionName: (string) required: The name of the collection to be deleted.
// Response:
// - 202: (string) The collection has been successfully deleted.
// - 404: The collection does not found.
// - 500: Internal error. Can be caused by invalid parameters.

package collection

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://controller.us-east1-gcp.pinecone.io/collections/collectionName"

	req, _ := http.NewRequest("DELETE", url, nil)

	req.Header.Add("accept", "text/plain")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}
