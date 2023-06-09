package config

import (
	"GoGinStarter/internal/log"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App     App     `yaml:"app"`
	DB      DB      `yaml:"db"`
	Cache   Cache   `yaml:"cache"`
	Session Session `yaml:"session"`
	Mail    Mail    `yaml:"mail"`
	Auth    Auth    `yaml:"auth"`
}

type App struct {
	Name        string `yaml:"name"`
	Env         string `yaml:"env"`
	Debug       bool   `yaml:"debug"`
	URL         string `yaml:"url"`
	Port        int    `yaml:"port"`
	Lang        string `yaml:"lang"`
	Maintenance bool   `yaml:"maintenance"`
}

type DB struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

type Cache struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database int    `default:"0" yaml:"database"`
	Password string `yaml:"password"`
}

type Session struct {
	Driver   string `yaml:"driver"`
	Lifetime int    `yaml:"lifetime"`
}

type Mail struct {
	Mailer      string `yaml:"mailer"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Encryption  string `yaml:"encryption"`
	FromAddress string `yaml:"from_address"`
	FromName    string `yaml:"from_name"`
}

type Auth struct {
	JWT JWT `yaml:"jwt"`
	OTP OTP `yaml:"otp"`
}

type JWT struct {
	Secret         string `yaml:"secret"`
	ExpirationTime int    `yaml:"expiration_time"`
}

type OTP struct {
	TokenLength    int `yaml:"token_length"`
	ExpirationTime int `yaml:"expiration_time"`
}

func ProvideConfig(log log.Log) *Config {
	// Open the YAML file
	file, err := os.Open(".yaml")
	if err != nil {
		log.Error("Failed to open config file: " + err.Error())
	}
	defer file.Close()

	// Decode the YAML data into a Config struct
	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Error("Failed to decode config file: " + err.Error())
	}

	return &config
}
