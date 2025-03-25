package model

type PostroutingRequest struct {
	IP      string
	ANumber string
	BNumber string
	Mark    int
	Tags    []string
}
