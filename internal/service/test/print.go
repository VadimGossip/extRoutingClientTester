package test

import (
	"fmt"
	"time"

	"github.com/VadimGossip/extRoutingClientTester/internal/logger"
	"go.uber.org/zap"
)

func (s *service) Print(taskId int64) {
	if testTask, ok := s.tasks[taskId]; ok {
		if testTask.Summary.Duration == nil {
			fmt.Printf("No duration statistics for task_id %d\n", taskId)
			return
		}
		fmt.Printf("Total Histogram %+v\n", testTask.Summary.Duration.Histogram)
		fmt.Printf("Request EMA Answer TotalDuration %+v\n", time.Duration(testTask.Summary.Duration.Ema.Value()))
		fmt.Printf("Request Min Answer Duration %+v\n", testTask.Summary.Duration.Min)
		fmt.Printf("Request Max Answer Duration %+v\n", testTask.Summary.Duration.Max)
		return
	}
	logger.Info("Unknown task id. Nothing to print", zap.Int64("task_id", taskId))
}
