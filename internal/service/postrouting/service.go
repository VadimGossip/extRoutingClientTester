package postrouting

import (
	"github.com/VadimGossip/extRoutingClientTester/internal/client/postrouting"
	def "github.com/VadimGossip/extRoutingClientTester/internal/service"
)

type service struct {
	postClient       postrouting.Client
	postCacheService def.PostroutingCacheService
	testService      def.TestService
}

var _ def.PostroutingService = (*service)(nil)

func NewService(postClient postrouting.Client, postCacheService def.PostroutingCacheService, testService def.TestService) *service {
	return &service{
		postClient:       postClient,
		postCacheService: postCacheService,
		testService:      testService,
	}
}
