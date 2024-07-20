package configs

import (
	"os"
	"strconv"
)

var appConfig *Config

type Config struct {
	postgresUser     string
	postgresPassword string
	postgresHost     string
	postgresPort     string
	postgresDB       string
}

func LoadConfig() {
	config := &Config{
		postgresUser:     getEnv("POSTGRES_USER", "wanrun"),
		postgresPassword: getEnv("POSTGRES_PASSWORD", ""),
		postgresHost:     getEnv("POSTGRES_HOST", "postgres"),
		postgresPort:     getEnv("POSTGRES_PORT", "5432"),
		postgresDB:       getEnv("POSTGRES_DB", "wanrun"),
	}
	appConfig = config
}

// 環境変数の取得
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// 環境変数取得のintバージョン
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// configの構造体いのポインタ取得
func AppConfig() *Config {
	return appConfig
}

/*
Getter
*/
func (c *Config) PostgresUser() string {
	return c.postgresUser
}

func (c *Config) PostgresPassword() string {
	return c.postgresPassword
}

func (c *Config) PostgresHost() string {
	return c.postgresHost
}

func (c *Config) PostgresPort() string {
	return c.postgresPort
}

func (c *Config) PostgresDB() string {
	return c.postgresDB
}
