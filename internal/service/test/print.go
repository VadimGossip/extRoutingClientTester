package test

import (
	"fmt"
	"time"

	"github.com/VadimGossip/extRoutingClientTester/internal/logger"
	"go.uber.org/zap"
)

func (s *service) Print(taskId int64) {
	if testTask, ok := s.tasks[taskId]; ok {
		fmt.Printf("Total Histogram %+v\n", testTask.summary.duration.histogram)
		fmt.Printf("Request EMA Answer TotalDuration %+v\n", time.Duration(testTask.summary.duration.ema.Value()))
		fmt.Printf("Request Min Answer Duration %+v\n", testTask.summary.duration.min)
		fmt.Printf("Request Max Answer Duration %+v\n", testTask.summary.duration.max)
		return
	}
	logger.Info("Unknown task id. Nothing to print", zap.Int64("task_id", taskId))
}
