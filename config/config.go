package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	AppConfig struct {
		Port    string `yaml:"port"`
		LogFile string `yaml:"log-file"`
	} `yaml:"app"`

	DBconn string `yaml:"db_conn"`
}

func LoadServerConfig(path string) (ServerConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return ServerConfig{}, err
	}
	defer f.Close()

	var cfg ServerConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return ServerConfig{}, err
	}

	return cfg, nil
}
