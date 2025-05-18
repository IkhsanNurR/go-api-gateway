package pathconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PublicPaths []string `yaml:"public_paths"`
}

func LoadConfigPath(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}
