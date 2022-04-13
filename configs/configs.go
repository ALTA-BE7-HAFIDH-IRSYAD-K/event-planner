package configs

import (
	"os"
	"sync"
)

type AppConfig struct {
	Port      string `yaml:"port"`
	SecretJWT string `yaml:"secretjwt"`
	Database  struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {

	var defaultConfig AppConfig
	defaultConfig.Port = "8000"
	defaultConfig.Database.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Database.Name = getEnv("DB_NAME", "event")
	defaultConfig.Database.Address = getEnv("ADDRESS", "localhost")
	defaultConfig.Database.Port = "3306"
	defaultConfig.Database.Username = getEnv("DB_USERNAME", "root")
	defaultConfig.Database.Password = getEnv("DB_PASSWORD", "admin123")

	return &defaultConfig
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
