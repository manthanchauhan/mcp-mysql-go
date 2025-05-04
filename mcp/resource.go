package mcp

const (
	MIME_TYPE_TEXT = "text/plain"

	RESOURCE_TYPE_FILE = "file"
)

type Resource struct {
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MimeType    string `json:"mimeType"`
}
