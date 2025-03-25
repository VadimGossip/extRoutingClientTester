package postrouting

import (
	"time"

	postModel "github.com/VadimGossip/extRoutingClientTester/internal/client/postrouting/model"
	"github.com/VadimGossip/extRoutingClientTester/internal/model"
)

func (s *service) send(taskId int64, request *model.PostroutingRequest) error {
	ts := time.Now()
	if request == nil {
		return nil
	}

	req := postModel.Request{
		IP:      request.IP,
		Anumber: request.ANumber,
		Bnumber: request.BNumber,
		Mark:    request.Mark,
		Tags:    request.Tags,
	}

	_, err := s.postClient.Send(&req)
	if err != nil {
		return err
	}

	s.testService.AddDurationToSummary(taskId, time.Since(ts))

	return nil

}
