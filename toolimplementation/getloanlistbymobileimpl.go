package toolimplementation

import (
	"encoding/json"
	"fmt"
	"mcp-mysql-go/config"
	"mcp-mysql-go/rest/getloanlistbymobile"

	mcp "github.com/manthanchauhan/mcp-go-util/mcp"

	"github.com/google/uuid"
)

var TOOL_GET_LOAN_LIST_BY_MOBILE = mcp.Tool{
	Name:        "get_loan_list_by_mobile",
	Description: "Get user's active loans list by mobile number",
	InputSchema: mcp.InputSchema{
		Type:     mcp.OBJECT_TYPE,
		Required: []string{"mobile"},
		Properties: map[string]mcp.InputSchemaProperty{
			"mobile": {
				Type:        mcp.STRING_TYPE,
				Description: "Mobile number of the user",
			},
			"page_no": {
				Type:        mcp.NUMBER_TYPE,
				Description: "Page number for pagination",
			},
			"page_size": {
				Type:        mcp.NUMBER_TYPE,
				Description: "Number of results per page",
			},
		},
	},
}

func GetLoanListByMobile(mobile string, pageNo, pageSize int) *mcp.ToolResult {
	toolCallId := uuid.New().String()

	// Call the REST API to get loans by mobile
	loans, err := getloanlistbymobile.GetLoanListByMobile(mobile, pageNo, pageSize)
	if err != nil {
		return &mcp.ToolResult{
			ToolName:   config.TOOL_GET_LOAN_LIST_BY_MOBILE,
			ToolCallId: toolCallId,
			Status:     mcp.STATUS_ERROR,
			Content: []mcp.ToolResultContent{
				{Type: mcp.CONTENT_TYPE_TEXT, Text: fmt.Sprintf("Error getting loans for mobile %s: %v", mobile, err)},
			},
		}
	}

	// Convert loans to JSON for a better display
	loansJSON, err := json.MarshalIndent(loans.Result, "", "  ")
	if err != nil {
		return &mcp.ToolResult{
			ToolName:   config.TOOL_GET_LOAN_LIST_BY_MOBILE,
			ToolCallId: toolCallId,
			Status:     mcp.STATUS_ERROR,
			Content: []mcp.ToolResultContent{
				{Type: mcp.CONTENT_TYPE_TEXT, Text: fmt.Sprintf("Error marshaling loan data: %v", err)},
			},
		}
	}

	return &mcp.ToolResult{
		ToolName:   config.TOOL_GET_LOAN_LIST_BY_MOBILE,
		ToolCallId: toolCallId,
		Status:     mcp.STATUS_SUCCESS,
		Content: []mcp.ToolResultContent{
			{Type: mcp.CONTENT_TYPE_TEXT, Text: string(loansJSON)},
		},
	}
}
