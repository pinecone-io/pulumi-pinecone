// File: provider/pkg/pinecone/utils/utils.go
package utils

import (
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

func (c *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Api-Key", c.APIKey)
	return c.Transport.RoundTrip(req)
}

var IndexDimensionDefault client.IndexDimension = 1536
