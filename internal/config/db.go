package config

var (
	DB_HOST     = Get("DB_HOST", "127.0.0.1")
	DB_PORT     = Get("DB_PORT", "3306")
	DB_NAME     = Get("DB_NAME", "database")
	DB_USERNAME = Get("DB_USERNAME", "root")
	DB_PASSWORD = Get("DB_PASSWORD", "password")
)
