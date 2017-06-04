package config

import "github.com/caarlos0/env"

// Config - configure environment variables, settings
type config struct {
	DbURL           string `env:"MONGODB_URI" envDefault:"mongodb://localhost/"`
	DbName          string `env:"MONGODB_NAME" envDefault:"go-webapp"`
	Port            string `env:"PORT" envDefault:"3000"`
	FbSecret        string `env:"FB_SECRET,required"`
	FbWebHookToken  string `env:"FB_WEBHOOK_TOKEN,required"`
	IgSecret        string `env:"IG_SECRET,required"`
	AuthTokenSecret string `env:"AUTH_TOKEN_SECRET,required"`
	RedisURL        string `env:"REDIS_URL" envDefault:"localhost:6379"`
	RedisPwd        string `env:"REDIS_PWD"`
}

// MakeConfig - add values to config from environment variables
func (c *config) MakeConfig() {

	err := env.Parse(c)

	if err != nil {
		panic("Error while parsing config. Please check all env. variables")
	}
}

// Settings - export settings, singleton
var Settings config
