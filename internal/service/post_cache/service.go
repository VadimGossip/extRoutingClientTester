package post_cache

import (
	"sync"

	"github.com/VadimGossip/extRoutingClientTester/internal/model"
	"github.com/VadimGossip/extRoutingClientTester/internal/repository"
	def "github.com/VadimGossip/extRoutingClientTester/internal/service"
)

var _ def.PostroutingCacheService = (*service)(nil)

type service struct {
	postroutingRepo repository.PostroutingRequestRepository
	requests        []model.PostroutingRequest
	mu              sync.Mutex
	offset          int
}

func NewService(postroutingRepo repository.PostroutingRequestRepository) *service {
	s := &service{postroutingRepo: postroutingRepo}
	return s
}
