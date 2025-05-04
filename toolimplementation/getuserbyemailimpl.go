package toolimplementation

import (
	"fmt"
	"mcp-mysql-go/config"
	"mcp-mysql-go/mcp"
	"mcp-mysql-go/user"
	"strings"

	"github.com/google/uuid"
)

var TOOL_GET_USER_BY_EMAIL = mcp.Tool{
	Name:        "get_order_list_by_user_id",
	Description: "Get Order List by User Id",
	InputSchema: mcp.InputSchema{
		Type:     mcp.OBJECT_TYPE,
		Required: []string{"user_id"},
		Properties: map[string]mcp.InputSchemaProperty{
			"user_id": mcp.InputSchemaProperty{
				Type:        mcp.NUMBER_TYPE,
				Description: "Id of the user",
			},
		},
	},
}

var USERS []user.User = []user.User{
	{Id: 1, Name: "Manthan Chauhan", Email: "manthanchauhan913@gmail.com"},
	{Id: 2, Name: "Harshita Agarwal", Email: "harshitaagarwal998@gmail.com"},
}

func GetUserByEmail(email string) *mcp.ToolResult {
	toolCallId := uuid.New().String()

	for _, user := range USERS {
		if strings.EqualFold(user.Email, email) {
			return &mcp.ToolResult{
				ToolName:   config.TOOL_GET_USER_BY_EMAIL,
				ToolCallId: toolCallId,
				Status:     mcp.STATUS_SUCCESS,
				Content: []mcp.ToolResultContent{
					{Type: mcp.CONTENT_TYPE_TEXT, Text: fmt.Sprintf("%+v", user)},
				},
			}
		}
	}

	return &mcp.ToolResult{
		ToolName:   config.TOOL_GET_USER_BY_EMAIL,
		ToolCallId: toolCallId,
		Status:     mcp.STATUS_ERROR,
		Content: []mcp.ToolResultContent{
			{Type: mcp.CONTENT_TYPE_TEXT, Text: fmt.Sprintf("User with email %v not found", email)},
		},
	}
}
