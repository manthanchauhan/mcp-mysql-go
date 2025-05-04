package mcp

// Capabilities represents MCP server capabilities
type Capabilities struct {
	Tools     *Capability `json:"tools,omitempty"`
	Resources *Capability `json:"resources,omitempty"`
	Prompts   *Capability `json:"prompts,omitempty"`
}
