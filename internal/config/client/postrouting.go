package client

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"

	def "github.com/VadimGossip/extRoutingClientTester/internal/config"
	"github.com/VadimGossip/extRoutingClientTester/internal/logger"
)

const (
	postroutingUrlEnvName = "POSTROUTING_URL"
	postroutingTTLEnvName = "POSTROUTING_TTL"
)

type postroutingClientConfig struct {
	url string
	ttl time.Duration
}

var _ def.PostroutingClientConfig = (*postroutingClientConfig)(nil)

func NewPostroutingClientConfig() (*postroutingClientConfig, error) {
	cfg := &postroutingClientConfig{}
	if err := cfg.setFromEnv(); err != nil {
		return nil, fmt.Errorf("postroutingClientConfig set from env err: %s", err)
	}

	logger.Infof("postroutingClientConfig: [%+v]", *cfg)
	return cfg, nil
}

func (cfg *postroutingClientConfig) setFromEnv() error {
	var err error
	cfg.url = os.Getenv(postroutingUrlEnvName)
	if len(cfg.url) == 0 {
		return fmt.Errorf("postroutingClientConfig host not found")
	}

	ttlStr := os.Getenv(postroutingTTLEnvName)
	if len(ttlStr) == 0 {
		return fmt.Errorf("postroutingClientConfig ttl not found")
	}

	cfg.ttl, err = time.ParseDuration(ttlStr)
	if err != nil {
		return errors.Wrap(err, "failed to parse postroutingClientConfig ttl")
	}

	return nil
}

func (cfg *postroutingClientConfig) Url() string {
	return cfg.url
}

func (cfg *postroutingClientConfig) TTL() time.Duration {
	return cfg.ttl
}
