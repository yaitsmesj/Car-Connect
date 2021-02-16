package config

import (
	"encoding/json"
	"os"

	"github.com/yaitsmesj/api-service/logger"
)

type Config struct {
	URL string `json:"RabitMQURL"`
}

var config *Config

// GetConfig returns instance of config
func GetConfig() *Config {
	if config == nil {
		loadConfig()
	}
	return config
}

// LoadConfig loads configurarion file
func loadConfig() {
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	logger.LogMessage(err, "Could not Load Configs", "Successfully Loaded Configs")

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
}
