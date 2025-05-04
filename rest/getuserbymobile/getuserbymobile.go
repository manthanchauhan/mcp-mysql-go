package getuserbymobile

import (
	"encoding/json"
	"fmt"
	"mcp-mysql-go/ig"
	"mcp-mysql-go/iguser"
	"mcp-mysql-go/rest"
)

// Service provides methods for retrieving IgUser data
type Service struct {
	restClient *rest.Client
	apiBaseURL string
}

// NewService creates a new IgUser service
func NewService(restClient *rest.Client, baseURL string) *Service {
	if baseURL == "" {
		baseURL = ig.API_BASE_URL + ig.API_ADMIN_USER_LIST
	}

	return &Service{
		restClient: restClient,
		apiBaseURL: baseURL,
	}
}

// GetUserByMobile retrieves a user by their mobile number
func GetUserByMobile(mobile string) (*iguser.IgUser, error) {
	restClient := rest.NewClient(10)

	// Create a new IgUser service
	s := NewService(restClient, "")

	if mobile == "" {
		return nil, fmt.Errorf("mobile is missing")
	}

	url, err := rest.BuildURL(s.apiBaseURL, map[string]string{
		"mobile": mobile,
	})
	if err != nil {
		return nil, fmt.Errorf("error building URL: %w", err)
	}

	headers := map[string]string{
		"Accept":         "application/json",
		"Content-Type":   "application/json",
		"x-access-token": ig.ADMIN_TOKEN,
	}

	data, err := s.restClient.Get(url, headers)
	if err != nil {
		return nil, fmt.Errorf("error fetching user by mobile: %w", err)
	}

	// Define response structure based on the actual API response format
	var apiResponse struct {
		Result []struct {
			ID              int     `json:"id"`
			Mobile          string  `json:"mobile"`
			Name            string  `json:"name"`
			Gender          *string `json:"gender"`
			Email           *string `json:"email"`
			Status          string  `json:"status"`
			IsEmailVerified bool    `json:"isEmailVerified"`
			ProfilePicture  *string `json:"profilePicture"`
			UserType        string  `json:"userType"`
			CreatedAt       int64   `json:"createdAt"`
		} `json:"result"`
		Status int     `json:"status"`
		Time   int64   `json:"time"`
		Error  *string `json:"error"`
	}

	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	// Check for API errors
	if apiResponse.Error != nil && *apiResponse.Error != "" {
		return nil, fmt.Errorf("API error: %s", *apiResponse.Error)
	}

	// Check if we got any results
	if len(apiResponse.Result) == 0 {
		return nil, fmt.Errorf("no user found with mobile: %s", mobile)
	}

	// Convert to our IgUser model
	userInfo := apiResponse.Result[0]
	igUser := &iguser.IgUser{
		ID:       userInfo.ID,
		Name:     userInfo.Name,
		Mobile:   userInfo.Mobile,
		Email:    "", // Initialize with empty string
		Username: "", // We don't have this in the response
		PhotoURL: "", // Initialize with empty string
	}

	// Set email if available
	if userInfo.Email != nil {
		igUser.Email = *userInfo.Email
	}

	// Set photo URL if available
	if userInfo.ProfilePicture != nil {
		igUser.PhotoURL = *userInfo.ProfilePicture
	}

	return igUser, nil
}
