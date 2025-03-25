package post_request

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/VadimGossip/extRoutingClientTester/internal/logger"
	"github.com/VadimGossip/extRoutingClientTester/internal/model"
	def "github.com/VadimGossip/extRoutingClientTester/internal/repository"
	"go.uber.org/zap"
)

type repository struct {
}

var _ def.TestRepository = (*repository)(nil)

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) readFileBytes(path string) ([]byte, error) {
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
	}
	return result, nil
}

func (r *repository) GetTestTasks() ([]model.TestTask, error) {
	testTasks := make([]model.TestTask, 0)
	bytes, err := r.readFileBytes("./data/test_tasks.json")
	if err != nil {
		return nil, fmt.Errorf("test repository. GetTestTasks. Can't read file, err = %s", err)
	}

	if err = json.Unmarshal(bytes, &testTasks); err != nil {
		return nil, fmt.Errorf("test repository repository. GetTestTasks. Can't unmarshal data, err = %s", err)
	}

	for i := range testTasks {
		testTasks[i].ID = int64(i + 1)
	}

	return testTasks, nil
}
