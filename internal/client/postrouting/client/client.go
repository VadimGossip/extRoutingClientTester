package client

import (
	"encoding/json"
	"fmt"

	def "github.com/VadimGossip/extRoutingClientTester/internal/client/postrouting"
	"github.com/VadimGossip/extRoutingClientTester/internal/client/postrouting/model"
	"github.com/VadimGossip/extRoutingClientTester/pkg/client/http"
)

type client struct {
	postroutingClient http.Client
}

func NewClient(postroutingClient http.Client) *client {
	return &client{
		postroutingClient: postroutingClient,
	}
}

var _ def.Client = (*client)(nil)

func (c *client) Send(req *model.Request) (*model.Response, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	respBytes, err := c.postroutingClient.SendPostRequest(reqBytes)
	if err != nil {
		return nil, err
	}

	res := &model.Response{}
	if err = json.Unmarshal(respBytes, res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res, nil
}
