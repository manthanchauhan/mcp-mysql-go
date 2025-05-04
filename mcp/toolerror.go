package mcp

const (
	NOT_FOUND = "NOT_FOUND"
)

type ToolError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
