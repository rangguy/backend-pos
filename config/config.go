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
	Port                  int      `json:"port"`
	AppName               string   `json:"appName"`
	AppEnv                string   `json:"appEnv"`
	SignatureKey          string   `json:"signatureKey"`
	Database              Database `json:"database"`
	RateLimiterMaxRequest int      `json:"rateLimiterMaxRequest"`
	RateLimiterTimeSecond int      `json:"rateLimiterTimeSecond"`
	JwtSecretKey          string   `json:"jwtSecretKey"`
	JwtExpirationTime     int      `json:"jwtExpirationTime"`
}

type Database struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnection     int    `json:"maxOpenConnection"`
	MaxLifetimeConnection int    `json:"maxLifetimeConnection"`
	MaxIdleConnection     int    `json:"maxIdleConnection"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

func Init() {
	err := util.BindFromJSON(&Config, "config.json", ".")
	if err != nil {
		logrus.Infof("failed to bind config: %v", err)
		err = util.BindFromConsul(&Config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_KEY"))
		if err != nil {
			panic(err)
		}
	}

	overrideFromEnv()

	logrus.Infof("Config loaded: %s (env: %s, port: %d)", Config.AppName, Config.AppEnv, Config.Port)
	logrus.Infof("Database: %s@%s:%d/%s", Config.Database.Username, Config.Database.Host, Config.Database.Port, Config.Database.Name)
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

	if Config.Database.Password == "" {
		logrus.Warn("DB_PASSWORD is not set!")
	}
	if Config.JwtSecretKey == "" {
		logrus.Warn("JWT_SECRET_KEY is not set!")
	}
}

func overrideFromEnv() {
	if port := os.Getenv("APP_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			Config.Port = p
		}
	}

	if appName := os.Getenv("APP_NAME"); appName != "" {
		Config.AppName = appName
	}

	if appEnv := os.Getenv("APP_ENV"); appEnv != "" {
		Config.AppEnv = appEnv
	}

	if signKey := os.Getenv("SIGNATURE_KEY"); signKey != "" {
		Config.SignatureKey = signKey
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		Config.Database.Host = dbHost
	}

	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		if p, err := strconv.Atoi(dbPort); err == nil {
			Config.Database.Port = p
		}
	}

	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		Config.Database.Name = dbName
	}

	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		Config.Database.Username = dbUser
	}

	if dbPass := os.Getenv("DB_PASSWORD"); dbPass != "" {
		Config.Database.Password = dbPass
	}

	if jwtKey := os.Getenv("JWT_SECRET_KEY"); jwtKey != "" {
		Config.JwtSecretKey = jwtKey
	}

	if jwtExp := os.Getenv("JWT_EXPIRATION_TIME"); jwtExp != "" {
		if exp, err := strconv.Atoi(jwtExp); err == nil {
			Config.JwtExpirationTime = exp
		}
	}

	if maxReq := os.Getenv("RATE_LIMITER_MAX_REQUEST"); maxReq != "" {
		if req, err := strconv.Atoi(maxReq); err == nil {
			Config.RateLimiterMaxRequest = req
		}
	}

	if timeSec := os.Getenv("RATE_LIMITER_TIME_SECOND"); timeSec != "" {
		if sec, err := strconv.Atoi(timeSec); err == nil {
			Config.RateLimiterTimeSecond = sec
		}
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
