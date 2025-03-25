package postrouting

import (
	"context"
	"sync"

	"github.com/VadimGossip/extRoutingClientTester/internal/logger"
	"github.com/VadimGossip/extRoutingClientTester/internal/model"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

func (s *service) RunTests(ctx context.Context) error {
	testTasks, err := s.testService.GetTestTasks()
	if err != nil {
		return err
	}

	for _, task := range testTasks {
		logger.Info("Test task started", zap.Int64("task_id", task.ID), zap.Int("total", task.Total), zap.Int("rps", task.Rps), zap.Int("pps", task.Pps), zap.Int("max_workers", task.MaxWorkers))
		if err = s.postCacheService.Refresh(int64(task.Total)); err != nil {
			return err
		}

		reqChan := make(chan model.PostroutingRequest)
		go func() {
			for i := 0; i < task.Total; i++ {
				req := s.postCacheService.GetRequest()
				reqChan <- req
			}
			close(reqChan)
		}()

		rateLimiter := rate.NewLimiter(rate.Limit(task.Rps), task.Pps)
		wg := &sync.WaitGroup{}
		for range task.MaxWorkers {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for req := range reqChan {
					if err := rateLimiter.Wait(ctx); err != nil {
						logger.Error("Postrouting service error",
							zap.String("method", "Run"),
							zap.String("problem", "rateLimiter wait"),
							zap.Error(err),
						)
					}

					if err := s.send(task.ID, &req); err != nil {
						logger.Error("Postrouting service error",
							zap.String("method", "Run"),
							zap.String("problem", "send"),
							zap.Error(err),
						)
						return
					}
				}
			}()
		}

		wg.Wait()
		logger.Info("Test task stopped", zap.Int64("task_id", task.ID), zap.Int("total", task.Total), zap.Int("rps", task.Rps), zap.Int("pps", task.Pps), zap.Int("max_workers", task.MaxWorkers))
		s.testService.Print(task.ID)
	}

	return nil

}
