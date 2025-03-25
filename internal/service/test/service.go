package test

import (
	"sync"
	"time"

	"github.com/VadimGossip/drs_storage_tester/pkg/util"
)

type durationSummary struct {
	max       time.Duration
	min       time.Duration
	ema       util.EMA
	histogram map[float64]int
}

type summary struct {
	total    int
	mu       sync.RWMutex
	duration *durationSummary
}

type task struct {
	requestsPerSec int
	packPerSec     int
	summary        *summary
}

type service struct {
	mu     sync.Mutex
	tasks  map[int64]*task
	lastId int64
}

func NewService() *service {
	return &service{tasks: map[int64]*task{}}
}
