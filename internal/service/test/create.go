package test

import (
	"github.com/VadimGossip/extRoutingClientTester/internal/logger"
	"go.uber.org/zap"
)

func (s *service) CreateTestTask(total, rps, pps int) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	taskId := s.lastId + 1

	s.tasks[taskId] = &task{
		requestsPerSec: rps,
		packPerSec:     pps,
		summary: &summary{
			total: 1000,
		},
	}
	logger.Info("New test created", zap.Int64("taskId", taskId), zap.Int("total", total), zap.Int("rps", rps), zap.Int("pps", pps))
	return taskId
}
