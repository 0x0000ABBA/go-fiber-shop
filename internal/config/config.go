package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"false"`
	
	Listen        struct {
		BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port   string `env:"PORT" env-default:"10000"`
	}

	AdminUser struct {
		Email    string `env:"ADMIN_EMAIL" env-default:"admin@admin.admin"`
		Password string `env:"ADMIN_PASSWORD" env-default:"admin"`
	}

	Postgres struct {
		Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
		Port     string `env:"POSTGRES_PORT" env-default:"5432"`
		User     string `env:"POSTGRES_USER" env-default:"postgres"`
		Password string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
		Database string `env:"POSTGRES_DATABASE" env-default:"postgres"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("Gathering config")
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "HELP_TEXT_PLACEHOLDER"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
