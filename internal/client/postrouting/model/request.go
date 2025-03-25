package model

type Request struct {
	IP      string   `json:"dst_ip"`
	Anumber string   `json:"anumber"`
	Bnumber string   `json:"bnumber"`
	Mark    int      `json:"mark"`
	Tags    []string `json:"tags,omitempty"`
}
