package http

type Client interface {
	SendPostRequest(reqBytes []byte) ([]byte, error)
}
