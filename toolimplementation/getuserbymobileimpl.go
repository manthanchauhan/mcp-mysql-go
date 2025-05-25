package toolimplementation

import (
	"encoding/json"
	"fmt"
	"mcp-mysql-go/config"
	"mcp-mysql-go/rest/getuserbymobile"

	"github.com/google/uuid"
	mcp "github.com/manthanchauhan/mcp-go-util/mcp"
)

var TOOL_GET_USER_BY_MOBILE = mcp.Tool{
	Name:        "get_user_by_mobile",
	Description: "Get user information by mobile number",
	InputSchema: mcp.InputSchema{
		Type:     mcp.OBJECT_TYPE,
		Required: []string{"mobile"},
		Properties: map[string]mcp.InputSchemaProperty{
			"mobile": {
				Type:        mcp.STRING_TYPE,
				Description: "Mobile number of the user",
			},
		},
	},
}

func GetUserByMobile(mobile string) *mcp.ToolResult {
	toolCallId := uuid.New().String()

	// Call the REST API to get user by mobile
	user, err := getuserbymobile.GetUserByMobile(mobile)
	if err != nil {
		return &mcp.ToolResult{
			ToolName:   config.TOOL_GET_USER_BY_MOBILE,
			ToolCallId: toolCallId,
			Status:     mcp.STATUS_ERROR,
			Content: []mcp.ToolResultContent{
				{Type: mcp.CONTENT_TYPE_TEXT, Text: fmt.Sprintf("Error getting user with mobile %s", mobile)},
			},
		}
	}

	// Convert user to JSON for a better display
	userJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return &mcp.ToolResult{
			ToolName:   config.TOOL_GET_USER_BY_MOBILE,
			ToolCallId: toolCallId,
			Status:     mcp.STATUS_ERROR,
			Content: []mcp.ToolResultContent{
				{Type: mcp.CONTENT_TYPE_TEXT, Text: "Error marshaling user data"},
			},
		}
	}

	return &mcp.ToolResult{
		ToolName:   config.TOOL_GET_USER_BY_MOBILE,
		ToolCallId: toolCallId,
		Status:     mcp.STATUS_SUCCESS,
		Content: []mcp.ToolResultContent{
			{Type: mcp.CONTENT_TYPE_TEXT, Text: string(userJSON)},
		},
	}
}
