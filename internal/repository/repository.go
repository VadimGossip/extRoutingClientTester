package repository

import "github.com/VadimGossip/extRoutingClientTester/internal/model"

type PostroutingRequestRepository interface {
	GetRequests(limit int64) ([]model.PostroutingRequest, error)
}
