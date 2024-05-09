// File: provider/pkg/pinecone/utils/utils.go
package utils

import (
	"encoding/json"
	"fmt"
	"github.com/pinecone-io/pulumi-pinecone/provider/pkg/pinecone/client"
	"net/http"
	"regexp"
)

func ValidateIndexName(name string) error {
	matched, err := regexp.MatchString("^[a-z0-9-]+$", name)
	if err != nil {
		return fmt.Errorf("regex error: %w", err)
	}
	if !matched {
		return fmt.Errorf("index name can only contain lower case alphanumeric characters and dashes")
	}
	return nil
}

type CustomTransport struct {
	Transport http.RoundTripper
	APIKey    string
}

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Status int `json:"status"`
}

func PrintErrorMessageFromResponse(response []byte) (*string, error) {
	jsonError := ErrorResponse{}
	err := json.Unmarshal(response, &jsonError)
	if err != nil {
		return nil, err
	}
	return &jsonError.Error.Message, nil
}

func (c *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Api-Key", c.APIKey)
	req.Header.Add("User-Agent", "source_tag=pulumi-pinecone")
	return c.Transport.RoundTrip(req)
}

var IndexDimensionDefault client.IndexDimension = 1536
