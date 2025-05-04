package rest

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// GetJSON performs a GET request and unmarshals the JSON response into the provided result
func (c *Client) GetJSON(url string, headers map[string]string, result interface{}) error {
	data, err := c.Get(url, headers)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}

	return nil
}

// PostJSON performs a POST request and unmarshals the JSON response into the provided result
func (c *Client) PostJSON(url string, payload interface{}, headers map[string]string, result interface{}) error {
	data, err := c.Post(url, payload, headers)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}

	return nil
}

// BuildURL builds a URL with the given base URL and query parameters
func BuildURL(baseURL string, queryParams map[string]string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("error parsing base URL: %w", err)
	}

	q := u.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}
