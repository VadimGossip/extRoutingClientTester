package post_cache

import "github.com/VadimGossip/extRoutingClientTester/internal/model"

func (s *service) GetRequest() model.PostroutingRequest {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.offset < len(s.requests)-1 {
		s.offset++
	} else {
		s.offset = 0
	}
	return s.requests[s.offset]
}
