// Action: (GET) list_collections
// About: This operation returns a list of all the collections in your current project.
// API Base URL: https://controller.us-east1-gcp.pinecone.io/collections
// Documentation URL: https://docs.pinecone.io/reference/list_collections
// Response Body: array of strings

package collection

import (
	"fmt"
	"io"
	"net/http"
)

func list_collections() {

	url := "https://controller.us-east1-gcp.pinecone.io/collections"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json; charset=utf-8")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}
