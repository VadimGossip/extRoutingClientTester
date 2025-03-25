package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	def "github.com/VadimGossip/extRoutingClientTester/pkg/client/http"
)

type client struct {
	url        string
	httpClient *http.Client
}

func NewClient(address string, ttl time.Duration) *client {
	return &client{
		url: address,
		httpClient: &http.Client{
			Timeout: ttl,
		},
	}
}

var _ def.Client = (*client)(nil)

func (c *client) SendPostRequest(reqBytes []byte) ([]byte, error) {
	httpReq, err := http.NewRequest(http.MethodPost, c.url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Close = true

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
