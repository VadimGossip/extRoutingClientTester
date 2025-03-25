package app

import (
	"log"

	"github.com/VadimGossip/extRoutingClientTester/internal/client/postrouting"
	postClient "github.com/VadimGossip/extRoutingClientTester/internal/client/postrouting/client"
	"github.com/VadimGossip/extRoutingClientTester/internal/config"
	clientCfg "github.com/VadimGossip/extRoutingClientTester/internal/config/client"
	"github.com/VadimGossip/extRoutingClientTester/internal/repository"
	postReqRepo "github.com/VadimGossip/extRoutingClientTester/internal/repository/post_request"
	testRepo "github.com/VadimGossip/extRoutingClientTester/internal/repository/test"
	"github.com/VadimGossip/extRoutingClientTester/internal/service"
	eventService "github.com/VadimGossip/extRoutingClientTester/internal/service/event"
	postCacheService "github.com/VadimGossip/extRoutingClientTester/internal/service/post_cache"
	postService "github.com/VadimGossip/extRoutingClientTester/internal/service/postrouting"
	testService "github.com/VadimGossip/extRoutingClientTester/internal/service/test"
	"github.com/VadimGossip/extRoutingClientTester/pkg/client/http"
	httpClient "github.com/VadimGossip/extRoutingClientTester/pkg/client/http/client"
)

type serviceProvider struct {
	postroutingClientConfig config.PostroutingClientConfig

	postroutingHTTPClient http.Client
	postroutingClient     postrouting.Client

	postReqRepo repository.PostroutingRequestRepository
	testRepo    repository.TestRepository

	testService        service.TestService
	eventService       service.EventService
	postCacheService   service.PostroutingCacheService
	postroutingService service.PostroutingService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PostroutingClientConfig() config.PostroutingClientConfig {
	if s.postroutingClientConfig == nil {
		cfg, err := clientCfg.NewPostroutingClientConfig()
		if err != nil {
			log.Fatalf("failed to get postroutingClientConfig: %s", err)
		}

		s.postroutingClientConfig = cfg
	}

	return s.postroutingClientConfig
}

func (s *serviceProvider) PostroutingHTTPClient() http.Client {
	if s.postroutingHTTPClient == nil {
		s.postroutingHTTPClient = httpClient.NewClient(s.PostroutingClientConfig().Url(), s.PostroutingClientConfig().TTL())
	}

	return s.postroutingHTTPClient
}

func (s *serviceProvider) PostroutingClient() postrouting.Client {
	if s.postroutingClient == nil {
		s.postroutingClient = postClient.NewClient(s.PostroutingHTTPClient())
	}

	return s.postroutingClient
}

func (s *serviceProvider) PostroutingRequestRepository() repository.PostroutingRequestRepository {
	if s.postReqRepo == nil {
		s.postReqRepo = postReqRepo.NewRepository()
	}
	return s.postReqRepo
}

func (s *serviceProvider) TestRepository() repository.TestRepository {
	if s.testRepo == nil {
		s.testRepo = testRepo.NewRepository()
	}
	return s.testRepo
}

func (s *serviceProvider) TestService() service.TestService {
	if s.testService == nil {
		s.testService = testService.NewService(s.TestRepository())
	}
	return s.testService
}

func (s *serviceProvider) EventService() service.EventService {
	if s.eventService == nil {
		s.eventService = eventService.NewService()
	}
	return s.eventService
}

func (s *serviceProvider) PostroutingCacheService() service.PostroutingCacheService {
	if s.postCacheService == nil {
		s.postCacheService = postCacheService.NewService(s.PostroutingRequestRepository())
	}
	return s.postCacheService
}

func (s *serviceProvider) PostroutingService() service.PostroutingService {
	if s.postroutingService == nil {
		s.postroutingService = postService.NewService(s.PostroutingClient(), s.PostroutingCacheService(), s.TestService())
	}
	return s.postroutingService
}
