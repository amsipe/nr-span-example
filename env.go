package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	HTTPPort        string `default:"5500" required:"true"`
	NewRelicKey     string `envconfig:"NEW_RELIC_LICENSE_KEY" required:"true"`
	NewRelicAppName string `envconfig:"NEW_RELIC_APP_NAME"`

	DBHost     string `required:"true" default:"exampledb"`
	DBPort     string `required:"true" default:"3306"`
	DBUsername string `required:"true" default:"root"`
	DBPassword string `required:"true" default:"password"`
}

func getEnvConfig() (config *envConfig, err error) {
	_ = godotenv.Load()

	config = &envConfig{}
	log.Println(config)
	err = envconfig.Process("", config)
	return config, err
}
