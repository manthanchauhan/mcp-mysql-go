package mcp

const (
	NUMBER_TYPE = "number"
	OBJECT_TYPE = "object"
	STRING_TYPE = "string"
)

type InputSchema struct {
	Type       string                         `json:"type"`
	Properties map[string]InputSchemaProperty `json:"properties"`
	Required   []string                       `json:"required"`
}
