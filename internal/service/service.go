package service

import (
	"context"
	"time"

	"github.com/VadimGossip/extRoutingClientTester/internal/model"
)

type EventService interface {
	RunEventGeneration(ctx context.Context, total, rps, pps int) chan int
}

type PostroutingCacheService interface {
	GetRequest() model.PostroutingRequest
	Refresh(limit int64) error
}

type PostroutingService interface {
	RunTests(ctx context.Context) error
}

type TestService interface {
	GetTestTasks() ([]model.TestTask, error)
	AddDurationToSummary(taskId int64, dur time.Duration)
	Print(taskId int64)
}
