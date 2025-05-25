package toolimplementation

import (
	"fmt"
	"mcp-mysql-go/config"
	"os"

	mcp "github.com/manthanchauhan/mcp-go-util/mcp"

	"github.com/google/uuid"
)

var TOOL_GET_LOAN_CLOSURE_PROCESS_STEP_FOR_SUPPORT_AGENT = mcp.Tool{
	Name:        "get_loan_closure_process_step_for_support_agent",
	Description: "Get the loan closure process steps guide for support agents",
	InputSchema: mcp.InputSchema{
		Type:       mcp.OBJECT_TYPE,
		Required:   []string{},
		Properties: map[string]mcp.InputSchemaProperty{},
	},
}

func GetLoanClosureProcessStepForSupportAgent() *mcp.ToolResult {
	toolCallId := uuid.New().String()

	// Read the loan closure process file
	filePath := "/Users/manthanchauhan/mcp-mysql-go/loanclosureprocess.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		return &mcp.ToolResult{
			ToolName:   config.TOOL_GET_LOAN_CLOSURE_PROCESS_STEP_FOR_SUPPORT_AGENT,
			ToolCallId: toolCallId,
			Status:     mcp.STATUS_ERROR,
			Content: []mcp.ToolResultContent{
				{Type: mcp.CONTENT_TYPE_TEXT, Text: fmt.Sprintf("Error reading loan closure process file: %v", err)},
			},
		}
	}

	// Return the content as a successful result
	return &mcp.ToolResult{
		ToolName:   config.TOOL_GET_LOAN_CLOSURE_PROCESS_STEP_FOR_SUPPORT_AGENT,
		ToolCallId: toolCallId,
		Status:     mcp.STATUS_SUCCESS,
		Content: []mcp.ToolResultContent{
			{Type: mcp.CONTENT_TYPE_TEXT, Text: string(content)},
		},
	}
}
