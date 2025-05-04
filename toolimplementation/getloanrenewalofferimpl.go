package toolimplementation

import (
	"encoding/json"
	"fmt"
	"mcp-mysql-go/config"
	"mcp-mysql-go/mcp"
	"mcp-mysql-go/rest/getloanrenewaloffer"

	"github.com/google/uuid"
)

var TOOL_GET_LOAN_RENEWAL_OFFER = mcp.Tool{
	Name:        "get_loan_renewal_offer",
	Description: "Get the best closure retention offer for a loan",
	InputSchema: mcp.InputSchema{
		Type:     mcp.OBJECT_TYPE,
		Required: []string{"loan_id"},
		Properties: map[string]mcp.InputSchemaProperty{
			"loan_id": {
				Type:        mcp.NUMBER_TYPE,
				Description: "ID of the loan",
			},
		},
	},
}

func GetLoanRenewalOffer(loanId int) *mcp.ToolResult {
	toolCallId := uuid.New().String()

	// Call the REST API to get loan renewal offer
	renewalOffer, err := getloanrenewaloffer.GetLoanRenewalOffer(loanId)
	if err != nil {
		return &mcp.ToolResult{
			ToolName:   config.TOOL_GET_LOAN_RENEWAL_OFFER,
			ToolCallId: toolCallId,
			Status:     mcp.STATUS_ERROR,
			Content: []mcp.ToolResultContent{
				{Type: mcp.CONTENT_TYPE_TEXT, Text: fmt.Sprintf("Error getting renewal offer for loan ID %d: %v", loanId, err)},
			},
		}
	}

	// Convert renewal offer to JSON for a better display
	offerJSON, err := json.MarshalIndent(renewalOffer.Result, "", "  ")
	if err != nil {
		return &mcp.ToolResult{
			ToolName:   config.TOOL_GET_LOAN_RENEWAL_OFFER,
			ToolCallId: toolCallId,
			Status:     mcp.STATUS_ERROR,
			Content: []mcp.ToolResultContent{
				{Type: mcp.CONTENT_TYPE_TEXT, Text: fmt.Sprintf("Error marshaling renewal offer data: %v", err)},
			},
		}
	}

	return &mcp.ToolResult{
		ToolName:   config.TOOL_GET_LOAN_RENEWAL_OFFER,
		ToolCallId: toolCallId,
		Status:     mcp.STATUS_SUCCESS,
		Content: []mcp.ToolResultContent{
			{Type: mcp.CONTENT_TYPE_TEXT, Text: string(offerJSON)},
		},
	}
}
