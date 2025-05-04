package mcp

const (
	STATUS_ERROR   = "error"
	STATUS_SUCCESS = "success"
)

const (
	CONTENT_TYPE_TEXT = "text"
)

type ToolResult struct {
	ToolName   string              `json:"tool_name"`
	ToolCallId string              `json:"tool_call_id"`
	Status     string              `json:"status"`
	Content    []ToolResultContent `json:"content,omitempty"`
}

type ToolResultContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
