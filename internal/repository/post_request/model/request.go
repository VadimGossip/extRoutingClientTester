package model

type PostroutingRequest struct {
	IP      string   `json:"ip"`
	Prefix  string   `json:"prefix"`
	ANumber string   `json:"anumber"`
	BNumber string   `json:"bnumber"`
	Mark    int      `json:"mark"`
	Tags    []string `json:"tags,omitempty"`
}
