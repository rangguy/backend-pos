package config

import (
	"backend/common/util"
	"github.com/sirupsen/logrus"
	_ "github.com/spf13/viper/remote"
	"os"
	"strconv"
)

var Config AppConfig

type AppConfig struct {
	Port                  int
	AppName               string
	AppEnv                string
	SignatureKey          string
	Database              Database
	RateLimiterMaxRequest int
	RateLimiterTimeSecond int
	JwtSecretKey          string
	JwtExpirationTime     int
}

type Database struct {
	Host                  string
	Port                  int
	Name                  string
	Username              string
	Password              string
	MaxOpenConnection     int
	MaxLifetimeConnection int
	MaxIdleConnection     int
	MaxIdleTime           int
}

func Init() {
	// 1️⃣ Default dari ENV (SOURCE OF TRUTH)
	loadFromEnv()

	// 2️⃣ Optional: override dari config.json (LOCAL DEV)
	if err := util.BindFromJSON(&Config, "config.json", "."); err == nil {
		logrus.Info("Config loaded from config.json")
	}

	// 3️⃣ Optional: override dari Consul
	if os.Getenv("CONSUL_HTTP_URL") != "" {
		if err := util.BindFromConsul(
			&Config,
			os.Getenv("CONSUL_HTTP_URL"),
			os.Getenv("CONSUL_HTTP_KEY"),
		); err != nil {
			logrus.Warnf("Failed to load config from Consul: %v", err)
		}
	}

	// 4️⃣ ENV selalu override terakhir
	overrideFromEnv()

	validate()

	logrus.Infof(
		"Config loaded: %s (%s) port=%d",
		Config.AppName,
		Config.AppEnv,
		Config.Port,
	)
	logrus.Infof(
		"Database: %s@%s:%d/%s",
		Config.Database.Username,
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.Name,
	)
}

func loadFromEnv() {
	Config = AppConfig{
		Port:         getEnvInt("APP_PORT", 8085),
		AppName:      getEnv("APP_NAME", "backend-pos"),
		AppEnv:       getEnv("APP_ENV", "development"),
		SignatureKey: getEnv("SIGNATURE_KEY", ""),
		Database: Database{
			Host:                  getEnv("DB_HOST", ""),
			Port:                  getEnvInt("DB_PORT", 5432),
			Name:                  getEnv("DB_NAME", ""),
			Username:              getEnv("DB_USER", ""),
			Password:              getEnv("DB_PASSWORD", ""),
			MaxOpenConnection:     getEnvInt("DB_MAX_OPEN_CONN", 10),
			MaxLifetimeConnection: getEnvInt("DB_MAX_LIFETIME_CONN", 10),
			MaxIdleConnection:     getEnvInt("DB_MAX_IDLE_CONN", 10),
			MaxIdleTime:           getEnvInt("DB_MAX_IDLE_TIME", 10),
		},
		RateLimiterMaxRequest: getEnvInt("RATE_LIMITER_MAX_REQUEST", 1000),
		RateLimiterTimeSecond: getEnvInt("RATE_LIMITER_TIME_SECOND", 60),
		JwtSecretKey:          getEnv("JWT_SECRET_KEY", ""),
		JwtExpirationTime:     getEnvInt("JWT_EXPIRATION_TIME", 1440),
	}
}

func overrideFromEnv() {
	if v := os.Getenv("APP_PORT"); v != "" {
		Config.Port, _ = strconv.Atoi(v)
	}
	if v := os.Getenv("APP_NAME"); v != "" {
		Config.AppName = v
	}
	if v := os.Getenv("APP_ENV"); v != "" {
		Config.AppEnv = v
	}
	if v := os.Getenv("SIGNATURE_KEY"); v != "" {
		Config.SignatureKey = v
	}
	if v := os.Getenv("DB_HOST"); v != "" {
		Config.Database.Host = v
	}
	if v := os.Getenv("DB_PORT"); v != "" {
		Config.Database.Port, _ = strconv.Atoi(v)
	}
	if v := os.Getenv("DB_NAME"); v != "" {
		Config.Database.Name = v
	}
	if v := os.Getenv("DB_USER"); v != "" {
		Config.Database.Username = v
	}
	if v := os.Getenv("DB_PASSWORD"); v != "" {
		Config.Database.Password = v
	}
	if v := os.Getenv("JWT_SECRET_KEY"); v != "" {
		Config.JwtSecretKey = v
	}
	if v := os.Getenv("JWT_EXPIRATION_TIME"); v != "" {
		Config.JwtExpirationTime, _ = strconv.Atoi(v)
	}
	if v := os.Getenv("RATE_LIMITER_MAX_REQUEST"); v != "" {
		Config.RateLimiterMaxRequest, _ = strconv.Atoi(v)
	}
	if v := os.Getenv("RATE_LIMITER_TIME_SECOND"); v != "" {
		Config.RateLimiterTimeSecond, _ = strconv.Atoi(v)
	}
}

func validate() {
	if Config.Database.Host == "" {
		logrus.Fatal("DB_HOST is required")
	}
	if Config.Database.Name == "" {
		logrus.Fatal("DB_NAME is required")
	}
	if Config.Database.Username == "" {
		logrus.Fatal("DB_USER is required")
	}
	if Config.JwtSecretKey == "" {
		logrus.Fatal("JWT_SECRET_KEY is required")
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
