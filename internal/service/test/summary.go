package test

import (
	"time"

	"github.com/VadimGossip/extRoutingClientTester/internal/service/test/model"
	"github.com/VadimGossip/extRoutingClientTester/pkg/ema"
	"github.com/VadimGossip/extRoutingClientTester/pkg/util"
)

func (s *service) AddDurationToSummary(taskId int64, dur time.Duration) {
	if testTask, ok := s.tasks[taskId]; ok {
		testTask.Summary.Mu.Lock()
		defer testTask.Summary.Mu.Unlock()
		var durInit bool
		if testTask.Summary.Duration == nil {
			testTask.Summary.Duration = &model.DurationSummary{
				Ema:       ema.NewEMA(0.01),
				Histogram: make(map[float64]int),
			}
			durInit = true
		}

		if testTask.Summary.Duration.Min > dur || durInit {
			testTask.Summary.Duration.Min = dur
		}
		if testTask.Summary.Duration.Max < dur || durInit {
			testTask.Summary.Duration.Max = dur
		}
		testTask.Summary.Duration.Ema.Add(float64(dur.Nanoseconds()))
		testTask.Summary.Duration.Histogram[(util.RoundFloat(float64(dur.Milliseconds()/100), 0)*100)+100]++
	}
}
