package resourceimplementation

import mcp "github.com/manthanchauhan/mcp-go-util/mcp"

var RESOURCE_LOAN_CLOSURE_STEPS = mcp.Resource{
	Uri:         "file:///Users/manthanchauhan/mcp-mysql-go/loanclosureprocess.txt",
	Name:        "Loan Closure Process",
	Description: "Loan Closure steps for customer support agent",
	MimeType:    mcp.MIME_TYPE_TEXT,
}
