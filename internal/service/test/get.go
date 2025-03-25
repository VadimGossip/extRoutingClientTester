package test

import (
	"github.com/VadimGossip/extRoutingClientTester/internal/model"
)

func (s *service) GetTestTasks() ([]model.TestTask, error) {
	testTasks, err := s.repoTest.GetTestTasks()
	if err != nil {
		return nil, err
	}

	for i := range testTasks {
		s.createTestTask(testTasks[i].ID, testTasks[i].Total, testTasks[i].Rps, testTasks[i].Pps)
	}
	return testTasks, nil
}
