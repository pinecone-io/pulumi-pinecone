// File: provider/pkg/pinecone/client/client.go
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	utils "github.com/usrbinkat/pulumi-pinecone-native/provider/pkg/pinecone/utils"
)

type PineconeClient struct {
	APIToken    string
	PineconeEnv string
	BaseURL     string
}

func NewPineconeClient(apiToken, pineconeEnv string) *PineconeClient {
	return &PineconeClient{
		APIToken:    apiToken,
		PineconeEnv: pineconeEnv,
		BaseURL:     fmt.Sprintf("https://controller.%s.pinecone.io/databases", pineconeEnv),
	}
}

func (client *PineconeClient) CreateIndex(data utils.IndexData) (*http.Response, error) {
	url := client.BaseURL

	indexData := map[string]interface{}{
		"name":      data.GetIndexName(),
		"dimension": data.GetIndexDimension(),
		"metric":    data.GetIndexMetric(),
		"pods":      data.GetIndexPods(),
		"replicas":  data.GetIndexReplicas(),
		"pod_type":  data.GetIndexPodType(),
	}

	jsonData, err := json.Marshal(indexData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling index data: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Api-Key", client.APIToken)
	req.Header.Add("accept", "text/plain")
	req.Header.Add("content-type", "application/json")

	return http.DefaultClient.Do(req)
}

func (client *PineconeClient) DeleteIndex(indexName string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", client.BaseURL, indexName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating DELETE request: %w", err)
	}

	req.Header.Add("Api-Key", client.APIToken)
	req.Header.Add("accept", "text/plain")

	return http.DefaultClient.Do(req)
}

func (client *PineconeClient) ListIndexes() ([]string, error) {
	url := client.BaseURL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Api-Key", client.APIToken)
	req.Header.Add("accept", "application/json; charset=utf-8")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("non-success status code received: %s", res.Status)
	}

	var indexes []string
	if err := json.NewDecoder(res.Body).Decode(&indexes); err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return indexes, nil
}

func (client *PineconeClient) DescribeIndex(indexName string) (*utils.IndexDetails, error) {
	url := fmt.Sprintf("%s/%s", client.BaseURL, indexName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Api-Key", client.APIToken)
	req.Header.Add("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get index details: status code %d", resp.StatusCode)
	}

	var indexDetails utils.IndexDetails
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &indexDetails)
	if err != nil {
		return nil, err
	}

	return &indexDetails, nil
}
