package config

import "os"

type env struct {
	Type       string
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
}

var global *env

func GetEnvVar() *env {
	return global
}

func IsLocal() bool {
	return global.Type == "local"
}

func IsDev() bool {
	return global.Type == "dev"
}

func init() {
	global = &env{
		Type:       os.Getenv("ENV_TYPE"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
