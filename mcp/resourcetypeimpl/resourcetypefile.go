package resourcetypeimpl

import (
	"os"
)

func ReadFileResource(path string) (*string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	contentStr := string(content)

	return &contentStr, nil
}
