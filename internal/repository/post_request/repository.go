package post_request

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/VadimGossip/extRoutingClientTester/internal/logger"
	"github.com/VadimGossip/extRoutingClientTester/internal/model"
	def "github.com/VadimGossip/extRoutingClientTester/internal/repository"
	"github.com/VadimGossip/extRoutingClientTester/internal/repository/post_request/converter"
	repoModel "github.com/VadimGossip/extRoutingClientTester/internal/repository/post_request/model"
	"go.uber.org/zap"
)

type repository struct {
}

var _ def.PostroutingRequestRepository = (*repository)(nil)

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) readFileBytes(path string, maxLines int64) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = f.Close(); err != nil {
			logger.Error("Postrouting request repository error",
				zap.String("method", "readFileBytes"),
				zap.String("problem", "file close"),
				zap.Error(err),
			)
			return
		}
	}()
	result := make([]byte, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Bytes()...)
		maxLines--
		if maxLines == 0 {
			return result, nil
		}
	}
	return result, nil
}

func (r *repository) GetRequests(limit int64) ([]model.PostroutingRequest, error) {
	repoRequests := make([]repoModel.PostroutingRequest, 0)
	bytes, err := r.readFileBytes("./data/postrouting_requests.json", limit)
	if err != nil {
		return nil, fmt.Errorf("postrouting request repository. GetRequests. Can't read file, err = %s", err)
	}

	if err = json.Unmarshal(bytes, &repoRequests); err != nil {
		return nil, fmt.Errorf("postrouting request repository. GetRequests. Can't unmarshal data, err = %s", err)
	}

	return converter.ToPostroutingRequestSFromRepo(repoRequests), nil
}
