package config

import "net/url"

const (
	ENV_LOCAL = "local"
)

var u, _ = url.Parse(Get("APP_URL", "http://localhost:8080"))
var relative, _ = url.Parse("/")

var (
	APP_ENV   = Get("APP_ENV", "local")
	APP_URL   = u.ResolveReference(relative).String()
	APP_PORT  = Get("APP_PORT", "8080")
	TIME_ZONE = Get("TIME_ZONE", "UTC")
)
