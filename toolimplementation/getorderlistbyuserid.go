package toolimplementation

import (
	"fmt"
	"mcp-mysql-go/config"
	"mcp-mysql-go/mcp"
	"mcp-mysql-go/order"

	"github.com/google/uuid"
)

var TOOL_GET_ORDER_LIST_BY_USER_ID mcp.Tool = mcp.Tool{
	Name:        "get_order_list_by_user_id",
	Description: "Get Order List by User Id",
	InputSchema: mcp.InputSchema{
		Type:     mcp.OBJECT_TYPE,
		Required: []string{"user_id"},
		Properties: map[string]mcp.InputSchemaProperty{
			"user_id": {
				Type:        mcp.NUMBER_TYPE,
				Description: "Id of the user",
			},
		},
	},
}

var ORDERS []order.Order = []order.Order{
	{Id: 1, UserId: 1, ItemName: "Dishwasher Liquid", Amount: 14, CreatedAt: 1746259289000},
	{Id: 2, UserId: 1, ItemName: "Onion", Amount: 4, CreatedAt: 1746172887000},
	{Id: 3, UserId: 1, ItemName: "Ginger", Amount: 3, CreatedAt: 1746086487000},
	{Id: 4, UserId: 2, ItemName: "Garlic", Amount: 7, CreatedAt: 1746259289000},
}

func GetOrderListByUserId(userId float64) *mcp.ToolResult {
	orderList := []order.Order{}

	for _, order := range ORDERS {
		if order.UserId == userId {
			orderList = append(orderList, order)
		}
	}

	toolCallId := uuid.New().String()

	if len(orderList) == 0 {
		return &mcp.ToolResult{
			ToolName:   config.TOOL_GET_ORDER_LIST_BY_USER_ID,
			ToolCallId: toolCallId,
			Status:     mcp.STATUS_ERROR,
			Content: []mcp.ToolResultContent{
				{
					Type: mcp.CONTENT_TYPE_TEXT,
					Text: fmt.Sprintf("No orders found for user id %v", userId),
				},
			},
		}
	}

	return &mcp.ToolResult{
		ToolName:   config.TOOL_GET_ORDER_LIST_BY_USER_ID,
		ToolCallId: toolCallId,
		Status:     mcp.STATUS_SUCCESS,
		Content: []mcp.ToolResultContent{
			{
				Type: mcp.CONTENT_TYPE_TEXT,
				Text: fmt.Sprintf("%+v", orderList),
			},
		},
	}
}
