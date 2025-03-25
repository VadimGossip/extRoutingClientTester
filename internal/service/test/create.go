package test

import (
	"github.com/VadimGossip/extRoutingClientTester/internal/logger"
	"github.com/VadimGossip/extRoutingClientTester/internal/service/test/model"
	"go.uber.org/zap"
)

func (s *service) createTestTask(taskId int64, total, rps, pps int) int64 {
	s.tasks[taskId] = &model.TestTask{
		RequestsPerSec: rps,
		PackPerSec:     pps,
		Summary: &model.Summary{
			Total: total,
		},
	}
	logger.Info("New test created", zap.Int64("taskId", taskId), zap.Int("total", total), zap.Int("rps", rps), zap.Int("pps", pps))
	return taskId
}
