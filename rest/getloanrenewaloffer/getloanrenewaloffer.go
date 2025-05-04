package getloanrenewaloffer

import (
	"encoding/json"
	"fmt"
	"mcp-mysql-go/ig"
	"mcp-mysql-go/loan"
	"mcp-mysql-go/rest"
)

// Service provides methods for retrieving loan renewal offers
type Service struct {
	restClient *rest.Client
	apiBaseURL string
}

// NewService creates a new Loan Renewal Offer service
func NewService(restClient *rest.Client, baseURL string) *Service {
	if baseURL == "" {
		baseURL = ig.API_BASE_URL + ig.API_SPIN_WHEEL
	}

	return &Service{
		restClient: restClient,
		apiBaseURL: baseURL,
	}
}

// GetLoanRenewalOffer retrieves the best closure retention offer for a loan
func GetLoanRenewalOffer(loanId int) (*loan.RenewalOfferResponse, error) {
	restClient := rest.NewClient(10)

	// Create a new Loan service
	s := NewService(restClient, "")

	if loanId <= 0 {
		return nil, fmt.Errorf("invalid loan ID: %d", loanId)
	}

	// Build URL with the loan ID
	url := fmt.Sprintf("%s/%d/best-closure-retention-offer", s.apiBaseURL, loanId)

	// Set headers
	headers := map[string]string{
		"Accept":         "application/json",
		"Content-Type":   "application/json",
		"x-access-token": ig.USER_TOKEN,
		"device-env":     "development",
		"device-type":    "DASHBOARD",
		"device-version": "12.2.6",
	}

	// Make the request
	data, err := s.restClient.Get(url, headers)
	if err != nil {
		return nil, fmt.Errorf("error fetching loan renewal offer: %w", err)
	}

	// Parse response
	var response loan.RenewalOfferResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	// Check for API errors
	if response.Error != nil && *response.Error != "" {
		return nil, fmt.Errorf("API error: %s", *response.Error)
	}

	return &response, nil
}
