package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	version    string = "dev"
	commitHash string = "-"
)

var cfg *Config

func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

type Grpc struct {
	Port              int    `yaml:"port"`
	MaxConnectionIdle int64  `yaml:"maxConnectionIdle"`
	Timeout           int64  `yaml:"timeout"`
	MaxConnectionAge  int64  `yaml:"maxConnectionAge"`
	Host              string `yaml:"host"`
}

type Rest struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	Version     string
	CommitHash  string
}

// type Metrics struct {
// 	Port int    `yaml:"port"`
// 	Host string `yaml:"host"`
// 	Path string `yaml:"path"`
// }

// type Status struct{
// 	Port          int    `yaml:"port"`
// 	Host          string `yaml:"host"`
// 	VersionPath   string `yaml:"versionPath"`
// 	LivenessPath  string `yaml:"livenessPath"`
// 	ReadinessPath string `yaml:"readinessPath"`
// }

// type Jaeger struct{
// 	Service string `yaml:"service"`
// 	Host    string `yaml:"host"`
// 	Port    string `yaml:"port"`
// }


type Database struct{
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Migrations string `yaml:"migrations"`
	Name       string `yaml:"name"`
	SslMode    string `yaml:"sslmode"`
	Driver     string `yaml:"driver"`
}

type Config struct {
	Project Project `yaml:"project"`
	Grpc    Grpc    `yaml:"grpc"`
	Rest    Rest    `yaml:"rest"`
	// Metrics Metrics `yaml:"metrics"`
	// Status Status `yaml:"status"`
	// Jaeger Jaeger `yaml:"jaeger"`
	Database Database `yaml:"database"`
}

func ReadConfigYML(filePath string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
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