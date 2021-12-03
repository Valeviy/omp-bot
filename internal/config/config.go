package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Build information -ldflags .
const (
	version    string = "dev"
	commitHash string = "-"
)

var cfg *Config

// GetConfigInstance returns service config
func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	ServiceName string `yaml:"serviceName"`
	Version     string
	CommitHash  string
}

// Metrics - contains all parameters metrics information.
type Metrics struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

// Jaeger - contains all parameters metrics information.
type Jaeger struct {
	Service string `yaml:"service"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
}

// Status config for service.
type Status struct {
	Port          int    `yaml:"port"`
	Host          string `yaml:"host"`
	VersionPath   string `yaml:"versionPath"`
	LivenessPath  string `yaml:"livenessPath"`
	ReadinessPath string `yaml:"readinessPath"`
}

// Telemetry config for logs.
type Telemetry struct {
	GraylogPath string `yaml:"graylogPath"`
}

// Bot config for telegram bot.
type Bot struct {
	Debug           bool   `yaml:"debug"`
	Timeout         int    `yaml:"timeout"`
	ConnectAttempts uint   `yaml:"connectAttempts"`
	PerPage         uint64 `yaml:"perPage"`
}

// EquipmentRequestAPI config for equipment request api.
type EquipmentRequestAPI struct {
	Address string `yaml:"address"`
}

// EquipmentRequestFacadeAPI config for equipment request facade api.
type EquipmentRequestFacadeAPI struct {
	Address string `yaml:"address"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project                   Project                   `yaml:"project"`
	Metrics                   Metrics                   `yaml:"metrics"`
	Jaeger                    Jaeger                    `yaml:"jaeger"`
	Status                    Status                    `yaml:"status"`
	Telemetry                 Telemetry                 `yaml:"telemetry"`
	Bot                       Bot                       `yaml:"bot"`
	EquipmentRequestAPI       EquipmentRequestAPI       `yaml:"equipmentRequestApi"`
	EquipmentRequestFacadeAPI EquipmentRequestFacadeAPI `yaml:"equipmentRequestFacadeApi"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(filePath string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}

	//nolint
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash

	return nil
}
