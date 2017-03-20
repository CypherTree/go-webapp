package config

import "os"

// Config - configure environment variables, settings
type config struct {
	DbURL           string
	DbName          string
	Port            string
	FbSecret        string
	IgSecret        string
	AuthTokenSecret string
	RedisURL        string
	RedisPwd        string
}

// MakeConfig - add values to config from environment variables
func (c *config) MakeConfig() {

	// default values
	c.DbURL = "mongodb://localhost/gowebapp-try"
	c.DbName = "gowebapp-try"
	c.Port = "3000"
	c.RedisURL = "localhost:6379"

	// get enviroment vars and override default values
	if DbURL := os.Getenv("MONGODB_URI"); DbURL != "" {
		c.DbURL = DbURL
	}

	if Port := os.Getenv("PORT"); Port != "" {
		c.Port = Port
	}

	if FbSecret := os.Getenv("FB_SECRET"); FbSecret != "" {
		c.FbSecret = FbSecret
	}

	if IgSecret := os.Getenv("IG_SECRET"); IgSecret != "" {
		c.IgSecret = IgSecret
	}

	if AuthTokenSecret := os.Getenv("AUTH_TOKEN_SECRET"); AuthTokenSecret != "" {
		c.AuthTokenSecret = AuthTokenSecret
	}

	if RedisURL := os.Getenv("REDIS_URL"); RedisURL != "" {
		c.RedisURL = RedisURL
	}

	if RedisPwd := os.Getenv("REDIS_PWD"); RedisPwd != "" {
		c.RedisPwd = RedisPwd
	}
}

// Settings - export settings, singleton
var Settings config
