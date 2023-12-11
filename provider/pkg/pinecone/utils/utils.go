// File: provider/pkg/pinecone/utils/utils.go
package utils

import (
	"fmt"
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
