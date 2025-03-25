package test

import (
	"time"

	"github.com/VadimGossip/extRoutingClientTester/pkg/ema"
	"github.com/VadimGossip/extRoutingClientTester/pkg/util"
)

func (s *service) AddDurationToSummary(taskId int64, dur time.Duration) {
	if testTask, ok := s.tasks[taskId]; ok {
		testTask.summary.mu.Lock()
		defer testTask.summary.mu.Unlock()
		var durInit bool
		if testTask.summary.duration == nil {
			testTask.summary.duration = &durationSummary{
				ema:       ema.NewEMA(0.01),
				histogram: make(map[float64]int),
			}
			durInit = true
		}

		if testTask.summary.duration.min > dur || durInit {
			testTask.summary.duration.min = dur
		}
		if testTask.summary.duration.max < dur || durInit {
			testTask.summary.duration.max = dur
		}
		testTask.summary.duration.ema.Add(float64(dur.Nanoseconds()))
		testTask.summary.duration.histogram[(util.RoundFloat(float64(dur.Milliseconds()/100), 0)*100)+100]++
	}
}
