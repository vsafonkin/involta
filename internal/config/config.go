package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	defaultRServerHost   = "localhost"
	defaultRServerPort   = "6534"
	defaultDB            = "testdb"
	defaultNamespace     = "docs"
	defaultAPIServerHost = "localhost"
	defaultAPIServerPort = "8080"
)

type Config struct {
	RServer
	APIServer
}

type RServer struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DBName    string `yaml:"dbname"`
	Namespace string `yaml:"namespace"`
}

type APIServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var config Config

func LoadConfig(path string) error {
	fmt.Println("Load config from", path)
	return yamlConfig(path)
}

func yamlConfig(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open config path error: %w", err)
	}
	defer f.Close()

	d := yaml.NewDecoder(f)

	if err := d.Decode(&config); err != nil {
		return fmt.Errorf("decode yaml config error: %w", err)
	}
	return nil
}

func RServerHost() string {
	if config.RServer.Host == "" {
		return defaultRServerHost
	}
	return config.RServer.Host
}

func RServerPort() string {
	if config.RServer.Port == "" {
		return defaultRServerPort
	}
	return config.RServer.Port
}

func DBName() string {
	if config.DBName == "" {
		return defaultDB
	}
	return config.DBName
}

func Namespace() string {
	if config.Namespace == "" {
		return defaultNamespace
	}
	return config.Namespace
}

func APIServerHost() string {
	if config.APIServer.Host == "" {
		return defaultAPIServerHost
	}
	return config.APIServer.Host
}

func APIServerPort() string {
	if config.APIServer.Port == "" {
		return defaultAPIServerPort
	}
	return config.APIServer.Port
}
