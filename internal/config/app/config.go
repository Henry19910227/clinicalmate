package app

import (
	model "clinicalmate/internal/model/config"
)

type config struct {
	cfg *model.AppConfig
}

func New(cfg *model.AppConfig) Config {
	return &config{cfg: cfg}
}

func (c *config) Name() string {
	return c.cfg.Name
}

func (c *config) Ip() string {
	return c.cfg.Ip
}

func (c *config) Port() int {
	return c.cfg.Port
}
