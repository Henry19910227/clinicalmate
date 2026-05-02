package app

import (
	model "clinicalmate/internal/model/config"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	cfg *model.Config
}

func New(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	var cfg model.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}
	return &config{cfg: &cfg}
}

func (c *config) Config() *model.Config {
	return c.cfg
}
