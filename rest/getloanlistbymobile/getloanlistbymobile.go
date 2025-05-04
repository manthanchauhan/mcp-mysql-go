package getloanlistbymobile

import (
	"encoding/json"
	"fmt"
	"mcp-mysql-go/ig"
	"mcp-mysql-go/loan"
	"mcp-mysql-go/rest"
)

// Service provides methods for retrieving loan data
type Service struct {
	restClient *rest.Client
	apiBaseURL string
}

// NewService creates a new Loan service
func NewService(restClient *rest.Client, baseURL string) *Service {
	if baseURL == "" {
		baseURL = ig.API_BASE_URL + ig.API_USER_PROFILE
	}

	return &Service{
		restClient: restClient,
		apiBaseURL: baseURL,
	}
}

// GetLoanListByMobile retrieves loans by user's mobile number
func GetLoanListByMobile(mobile string, pageNo int, pageSize int) (*loan.LoanListResponse, error) {
	restClient := rest.NewClient(10)

	// Create a new Loan service
	s := NewService(restClient, "")

	if mobile == "" {
		return nil, fmt.Errorf("mobile is missing")
	}

	if pageNo <= 0 {
		pageNo = 1 // Default to first page
	}

	if pageSize <= 0 {
		pageSize = 10 // Default page size
	}

	// Build URL with query parameters
	url, err := rest.BuildURL(s.apiBaseURL, map[string]string{
		"mobile":   mobile,
		"pageNo":   fmt.Sprintf("%d", pageNo),
		"pageSize": fmt.Sprintf("%d", pageSize),
	})
	if err != nil {
		return nil, fmt.Errorf("error building URL: %w", err)
	}

	// Set headers
	headers := map[string]string{
		"Accept":         "application/json",
		"Content-Type":   "application/json",
		"x-access-token": ig.ADMIN_TOKEN,
		"device-env":     "development",
		"device-type":    "DASHBOARD",
		"device-version": "12.2.6",
	}

	// Make the request
	data, err := s.restClient.Get(url, headers)
	if err != nil {
		return nil, fmt.Errorf("error fetching loans by mobile: %w", err)
	}

	// Parse response
	var response loan.LoanListResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	// Check for API errors
	if response.Error != nil && *response.Error != "" {
		return nil, fmt.Errorf("API error: %s", *response.Error)
	}

	return &response, nil
}
