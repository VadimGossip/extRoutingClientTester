package model

type TestTask struct {
	ID         int64
	Total      int `json:"total"`
	Rps        int `json:"rps"`
	Pps        int `json:"pps"`
	MaxWorkers int `json:"max_workers"`
}
