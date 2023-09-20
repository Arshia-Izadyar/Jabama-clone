package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Logger   LoggerConfig
	Otp      OtpConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port    int
	RunMode string
}

type JWTConfig struct {
	Secret                     string
	RefreshSecret              string
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
}

type PostgresConfig struct {
	Host            string
	User            string
	Password        string
	DbName          string
	SslMode         string
	Port            int
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               int
	Password           string
	Db                 int
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	PoolTimeout        int
	IdleCheckFrequency int
}
type OtpConfig struct {
	Digits     int
	ExpireTime time.Duration
	Limiter    time.Duration
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}

func configPath(env string) string {
	if env == "production" {
		return "../config/config-production.yml"
	}
	if env == "docker" {
		return "../config/config-docker.yml"
	}
	return "../config/config-development.yml"
}

func loadConfig(fileName, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(fileName)
	v.SetConfigType(fileType)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func parseConfig(v *viper.Viper) (cfg *Config, err error) {
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func GetConfig() *Config {
	env := os.Getenv("APP_ENV")
	cfgPath := configPath(env)
	v, err := loadConfig(cfgPath, "yml")
	if err != nil {
		return nil
	}
	cfg, err := parseConfig(v)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
