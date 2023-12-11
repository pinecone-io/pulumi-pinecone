// File: provider/pkg/pinecone/utils/types.go
package utils

type IndexDetails struct {
	Name           string `json:"name"`
	Dimension      int    `json:"dimension"`
	Metric         string `json:"metric"`
	Pods           int    `json:"pods"`
	Replicas       int    `json:"replicas"`
	PodType        string `json:"pod_type"`
	MetadataConfig struct {
		Indexed []string `json:"indexed"`
	} `json:"metadata_config"`
	Status struct {
		Ready bool   `json:"ready"`
		State string `json:"state"`
		Host  string `json:"host"`
		Port  int    `json:"port"`
	} `json:"status"`
}

type IndexData interface {
	GetIndexName() string
	GetIndexDimension() int
	GetIndexMetric() string
	GetIndexPods() int
	GetIndexReplicas() int
	GetIndexPodType() string
}

// PineconeIndexArgs methods to satisfy the IndexData interface
func (args PineconeIndexArgs) GetIndexName() string {
	return args.IndexName
}

func (args PineconeIndexArgs) GetIndexDimension() int {
	return args.IndexDimension
}

func (args PineconeIndexArgs) GetIndexMetric() string {
	return args.IndexMetric
}

func (args PineconeIndexArgs) GetIndexPods() int {
	return args.IndexPods
}

func (args PineconeIndexArgs) GetIndexReplicas() int {
	return args.IndexReplicas
}

func (args PineconeIndexArgs) GetIndexPodType() string {
	return args.IndexPodType
}

type PineconeIndexArgs struct {
	IndexName      string `pulumi:"name"`
	IndexDimension int    `pulumi:"dimension"`
	IndexMetric    string `pulumi:"metric"`
	IndexPods      int    `pulumi:"pods"`
	IndexReplicas  int    `pulumi:"replicas"`
	IndexPodType   string `pulumi:"podType"`
}
