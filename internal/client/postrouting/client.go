package postrouting

import (
	"github.com/VadimGossip/extRoutingClientTester/internal/client/postrouting/model"
)

type Client interface {
	Send(req *model.Request) (*model.Response, error)
}
