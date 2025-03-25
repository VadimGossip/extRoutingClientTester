package config

import (
	"time"
)

type PostroutingClientConfig interface {
	Url() string
	TTL() time.Duration
}
