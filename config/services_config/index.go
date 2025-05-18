package servicesconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ServiceConfig struct {
	Name       string `yaml:"name"`
	PathPrefix string `yaml:"path_prefix"`
	TargetURL  string `yaml:"target_url"`
}

type Config struct {
	Services []ServiceConfig `yaml:"services"`
}

func LoadConfigServices(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}
