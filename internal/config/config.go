package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
}

var configuration *Config

func LoadConfig() error {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "local" // 默認環境
	}

	configFile := fmt.Sprintf("configs/appsettings.%s.json", env)
	absPath, err := filepath.Abs(configFile)
	if err != nil {
		return fmt.Errorf("無法獲取配置文件路徑: %v", err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		return fmt.Errorf("無法打開配置文件: %v", err)
	}
	defer file.Close()

	configuration = new(Config)
	if err := json.NewDecoder(file).Decode(configuration); err != nil {
		return fmt.Errorf("無法解析配置文件: %v", err)
	}

	return nil
}

func GetConfig() *Config {
	return configuration
}
