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
	total := 10000
	rps := 1000
	pps := 20
	maxWorkers := 10

	taskId := s.testService.CreateTestTask(total, rps, pps)
	if err := s.postCacheService.Refresh(int64(total)); err != nil {
		return err
	}

	reqChan := make(chan model.PostroutingRequest)
	go func() {
		for i := 0; i < total; i++ {
			req := s.postCacheService.GetRequest()
			reqChan <- req
		}
		close(reqChan)
	}()

	rateLimiter := rate.NewLimiter(rate.Limit(rps), pps)
	wg := &sync.WaitGroup{}
	for range maxWorkers {
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

				if err := s.send(taskId, &req); err != nil {
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
	s.testService.Print(taskId)

	return nil

}
