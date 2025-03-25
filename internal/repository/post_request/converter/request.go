package converter

import (
	"github.com/VadimGossip/extRoutingClientTester/internal/model"
	repoModel "github.com/VadimGossip/extRoutingClientTester/internal/repository/post_request/model"
)

func ToPostroutingRequestFromRepo(request *repoModel.PostroutingRequest) *model.PostroutingRequest {
	if request == nil {
		return nil
	}

	bNumber := request.BNumber
	if request.Prefix != "" {
		bNumber = request.Prefix + request.BNumber
	}

	return &model.PostroutingRequest{
		ANumber: request.ANumber,
		BNumber: bNumber,
		IP:      request.IP,
		Mark:    request.Mark,
		Tags:    request.Tags,
	}
}

func ToPostroutingRequestSFromRepo(requests []repoModel.PostroutingRequest) []model.PostroutingRequest {
	result := make([]model.PostroutingRequest, 0, len(requests))
	for i := range requests {
		req := ToPostroutingRequestFromRepo(&requests[i])
		if req == nil {
			continue
		}
		result = append(result, *req)
	}
	return result
}
