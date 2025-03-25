package model

import (
	"sync"
	"time"

	"github.com/VadimGossip/drs_storage_tester/pkg/util"
)

type DurationSummary struct {
	Max       time.Duration
	Min       time.Duration
	Ema       util.EMA
	Histogram map[float64]int
}

type Summary struct {
	Total    int
	Mu       sync.RWMutex
	Duration *DurationSummary
}

type TestTask struct {
	RequestsPerSec int
	PackPerSec     int
	Summary        *Summary
}
