package main

import (
	"encoding/json"
	"log"
	"mcp-mysql-go/config"
	"mcp-mysql-go/jsonrpc"
	"mcp-mysql-go/mcp"
	"mcp-mysql-go/mcp/resourcetypeimpl"
	"mcp-mysql-go/resourceimplementation"
	"mcp-mysql-go/rest/getloanrenewaloffer"
	"mcp-mysql-go/toolimplementation"
	"os"
	"strings"
)

const SERVER_NAME = "mcp-mysql"
const SERVER_VERSION = "0.0.1"

var DECODER = json.NewDecoder(os.Stdin)
var ENCODER = json.NewEncoder(os.Stdout)

func runMcpServer() {
	log.SetOutput(os.Stderr)
	log.Printf("Hi starting...")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Printf("Starting %s MCP server...", config.ServerName)

	for {
		var request jsonrpc.JSONRPCRequest

		if err := DECODER.Decode(&request); err != nil {
			SendErrorResponse(nil, jsonrpc.PARSE_ERROR, "Parse error", nil)
			continue
		}

		log.Printf("Received request: %+v", request)

		if request.JSONRPC != jsonrpc.VERSION {
			SendErrorResponse(request.Id, jsonrpc.INVALID_JSON, "Invalid JSON", nil)
			continue
		}

		switch request.Method {
		case mcp.INITIALIZE:
			SendInitializeResponse(request.Id)
		case mcp.INITIALIZED:
			log.Printf("Server initialized successfully")
		case mcp.LIST_TOOLS:
			SendListToolsResponse(request.Id)
		case mcp.CALL_TOOLS:
			SendToolCallResponse(request)
		case mcp.LIST_PROMPTS:
			SendListPromptsResponse(request.Id)
		case mcp.LIST_RESOURCES:
			SendListResourcesResponse(request.Id)
		case mcp.READ_RESOURCE:
			sendReadResourceResponse(&request)
		default:
			SendErrorResponse(request.Id, jsonrpc.METHOD_NOT_FOUND, "Method not found", nil)
		}
	}
}

func SendErrorResponse(id any, code int, message string, data any) {
	var response *jsonrpc.JSONRPCResponse = jsonrpc.CreateErrorResponse(id, code, message, data)

	log.Printf("Sending error response: %+v", *response)

	if err := ENCODER.Encode(response); err != nil {
		log.Printf("Error sending response: %v", err)
	}
}

func SendResponse(id any, result any) {
	resp := jsonrpc.JSONRPCResponse{
		JSONRPC: jsonrpc.VERSION,
		Result:  result,
		Id:      id,
		Error:   nil,
	}

	log.Printf("Sending response: %+v", resp)

	if err := ENCODER.Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func SendInitializeResponse(id any) {
	result := mcp.InitializeResult{
		ProtocolVersion: mcp.VERSION,
		ServerInfo: mcp.ServerInfo{
			Name:    SERVER_NAME,
			Version: SERVER_VERSION,
		},
		Capabilities: mcp.Capabilities{
			Tools:     &mcp.Capability{ListChanged: false},
			Resources: nil,
			Prompts:   nil,
		},
	}

	SendResponse(id, result)
}

func SendListToolsResponse(id any) {
	result := mcp.ListToolsResult{
		Tools: []mcp.Tool{
			toolimplementation.TOOL_GET_ORDER_LIST_BY_USER_ID,
			toolimplementation.TOOL_GET_USER_BY_EMAIL,
			toolimplementation.TOOL_GET_USER_BY_MOBILE,
			toolimplementation.TOOL_GET_LOAN_LIST_BY_MOBILE,
			toolimplementation.TOOL_GET_LOAN_CLOSURE_PROCESS_STEP_FOR_SUPPORT_AGENT,
			toolimplementation.TOOL_GET_LOAN_RENEWAL_OFFER,
		},
	}

	SendResponse(id, result)
}

func SendListPromptsResponse(id any) {
	result := mcp.PromptListResponse{
		Prompts: []any{},
	}

	SendResponse(id, result)
}

func SendListResourcesResponse(id any) {
	result := mcp.ResourceListResponse{
		Resources: []mcp.Resource{
			resourceimplementation.RESOURCE_LOAN_CLOSURE_STEPS,
		},
	}

	SendResponse(id, result)
}

func SendToolCallResponse(request jsonrpc.JSONRPCRequest) {
	paramsMap, ok := request.Params.(map[string]any)

	if !ok {
		SendErrorResponse(request.Id, jsonrpc.INVALID_PARAMS, "Invalid params", nil)
		return
	}

	toolName, ok := paramsMap["name"].(string)

	if !ok {
		SendErrorResponse(request.Id, jsonrpc.INVALID_PARAMS, "Invalid params", nil)
		return
	}

	switch toolName {
	case config.TOOL_GET_USER_BY_EMAIL:
		handleToolGetUserByEmail(request.Id, paramsMap)
	case config.TOOL_GET_ORDER_LIST_BY_USER_ID:
		handleToolGetOrderListByUserId(request.Id, paramsMap)
	case config.TOOL_GET_USER_BY_MOBILE:
		handleToolGetUserByMobile(request.Id, paramsMap)
	case config.TOOL_GET_LOAN_LIST_BY_MOBILE:
		handleToolGetLoanListByMobile(request.Id, paramsMap)
	case config.TOOL_GET_LOAN_CLOSURE_PROCESS_STEP_FOR_SUPPORT_AGENT:
		handleToolGetLoanClosureProcessStepForSupportAgent(request.Id)
	case config.TOOL_GET_LOAN_RENEWAL_OFFER:
		handleToolGetLoanRenewalOffer(request.Id, paramsMap)
	default:
		SendErrorResponse(request.Id, jsonrpc.INVALID_PARAMS, "Invalid tool name", nil)
	}

	return
}

func handleToolGetUserByEmail(requestId any, paramsMap map[string]any) {
	arguments, ok := paramsMap["arguments"].(map[string]interface{})
	if !ok {
		SendErrorResponse(requestId, jsonrpc.INVALID_PARAMS, "Invalid params", nil)
		return
	}

	email, ok := arguments["email"].(string)
	if !ok {
		SendErrorResponse(requestId, jsonrpc.INVALID_PARAMS, "Invalid params", nil)
		return
	}
	user := toolimplementation.GetUserByEmail(email)
	SendResponse(requestId, user)
}

func handleToolGetOrderListByUserId(requestId any, paramsMap map[string]any) {
	arguments, ok := paramsMap["arguments"].(map[string]any)
	if !ok {
		SendErrorResponse(requestId, jsonrpc.INVALID_PARAMS, "Invalid params", nil)
		return
	}

	userId, ok := arguments["user_id"].(float64)
	if !ok {
		SendErrorResponse(requestId, jsonrpc.INVALID_PARAMS, "Invalid params", nil)
		return
	}

	orderList := toolimplementation.GetOrderListByUserId(userId)
	SendResponse(requestId, orderList)
}

func handleToolGetUserByMobile(requestId any, paramsMap map[string]any) {
	arguments, ok := paramsMap["arguments"].(map[string]any)
	if !ok {
		SendErrorResponse(requestId, jsonrpc.INVALID_PARAMS, "Invalid params", nil)
		return
	}

	mobile, ok := arguments["mobile"].(string)
	if !ok {
		SendErrorResponse(requestId, jsonrpc.INVALID_PARAMS, "Invalid mobile parameter", nil)
		return
	}

	// Call the tool implementation
	result := toolimplementation.GetUserByMobile(mobile)
	SendResponse(requestId, result)
}

func handleToolGetLoanListByMobile(requestId any, paramsMap map[string]any) {
	arguments, ok := paramsMap["arguments"].(map[string]any)
	if !ok {
		SendErrorResponse(requestId, jsonrpc.INVALID_PARAMS, "Invalid params", nil)
		return
	}

	mobile, ok := arguments["mobile"].(string)
	if !ok {
		SendErrorResponse(requestId, jsonrpc.INVALID_PARAMS, "Invalid mobile parameter", nil)
		return
	}

	// Get optional parameters with defaults
	pageNo := 1
	pageSize := 10

	if pageNoValue, exists := arguments["page_no"]; exists {
		if pageNoFloat, ok := pageNoValue.(float64); ok {
			pageNo = int(pageNoFloat)
		}
	}

	if pageSizeValue, exists := arguments["page_size"]; exists {
		if pageSizeFloat, ok := pageSizeValue.(float64); ok {
			pageSize = int(pageSizeFloat)
		}
	}

	// Call the tool implementation
	result := toolimplementation.GetLoanListByMobile(mobile, pageNo, pageSize)
	SendResponse(requestId, result)
}

func sendReadResourceResponse(request *jsonrpc.JSONRPCRequest) {
	paramsMap, ok := request.Params.(map[string]any)

	if !ok {
		SendErrorResponse(request.Id, jsonrpc.INVALID_PARAMS, "Invalid params", nil)
		return
	}

	resourceUri, ok := paramsMap["uri"].(string)

	if !ok {
		SendErrorResponse(request.Id, jsonrpc.INVALID_PARAMS, "Invalid params", nil)
		return
	}

	uriComponents := strings.Split(resourceUri, "://")

	if len(uriComponents) != 2 {
		SendErrorResponse(request.Id, jsonrpc.INVALID_PARAMS, "Invalid Uri", nil)
		return
	}

	resourceType, resourcePath := uriComponents[0], uriComponents[1]

	switch resourceType {
	case mcp.RESOURCE_TYPE_FILE:
		text, err := resourcetypeimpl.ReadFileResource(resourcePath)

		result := mcp.ResourceResult{
			Uri: resourceUri,
		}

		if err != nil {
			SendErrorResponse(request.Id, jsonrpc.RESOURCE_NOT_FOUND, "Resource not found", result)
			return
		}

		response := mcp.ResourceResultResponse{
			Contents: []mcp.ResourceResult{
				{Uri: resourceUri, MimeType: mcp.MIME_TYPE_TEXT, Text: *text},
			},
		}

		SendResponse(request.Id, response)
	default:
		SendErrorResponse(request.Id, jsonrpc.RESOURCE_NOT_FOUND, "Resource not found", nil)
	}
}

func handleToolGetLoanClosureProcessStepForSupportAgent(requestId any) {
	// This tool doesn't require any parameters, so we directly call the implementation
	result := toolimplementation.GetLoanClosureProcessStepForSupportAgent()
	SendResponse(requestId, result)
}

func handleToolGetLoanRenewalOffer(requestId any, paramsMap map[string]any) {
	arguments, ok := paramsMap["arguments"].(map[string]any)
	if !ok {
		SendErrorResponse(requestId, jsonrpc.INVALID_PARAMS, "Invalid params", nil)
		return
	}

	loanIdFloat, ok := arguments["loan_id"].(float64)
	if !ok {
		SendErrorResponse(requestId, jsonrpc.INVALID_PARAMS, "Invalid loan_id parameter", nil)
		return
	}

	loanId := int(loanIdFloat)

	// Call the tool implementation
	result := toolimplementation.GetLoanRenewalOffer(loanId)
	SendResponse(requestId, result)
}

func main() {
	// runMcpServer()
	getloanrenewaloffer.GetLoanRenewalOffer(107695)
}
