// File: provider/pkg/pinecone/client/client.go
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PineconeClient struct {
	APIToken    string
	PineconeEnv string
	BaseURL     string
}

type IndexData interface {
	GetIndexName() string
	GetIndexDimension() int
	GetIndexMetric() string
	GetIndexPods() int
	GetIndexReplicas() int
	GetIndexPodType() string
}

func NewPineconeClient(apiToken, pineconeEnv string) *PineconeClient {
	return &PineconeClient{
		APIToken:    apiToken,
		PineconeEnv: pineconeEnv,
		BaseURL:     fmt.Sprintf("https://controller.%s.pinecone.io/databases", pineconeEnv),
	}
}

func (client *PineconeClient) CreateIndex(data IndexData) (*http.Response, error) {
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
