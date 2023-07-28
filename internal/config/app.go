package config

const (
	ENV_LOCAL = "local"
)

var (
	APP_ENV  = Get("APP_ENV", "local")
	APP_PORT = Get("APP_PORT", "8080")
)
