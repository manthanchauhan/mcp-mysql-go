package jsonrpc

const VERSION = "2.0"

const (
	PARSE_ERROR           = -32700 // Parse error 		Invalid JSON was received by the server. An error occurred on the server while parsing the JSON text.
	INVALID_JSON          = -32600 // Invalid Request	The JSON sent is not a valid Request object.
	METHOD_NOT_FOUND      = -32601 // Method not found	The method does not exist / is not available.
	INVALID_PARAMS        = -32602 // Invalid params	Invalid method parameter(s).
	INTERNAL_SERVER_ERROR = -32603 // Internal error	Internal JSON-RPC error.
	RESOURCE_NOT_FOUND    = -32002 // Resource not found
)
