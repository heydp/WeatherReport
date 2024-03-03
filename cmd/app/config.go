package main

import "os"

func GetAppEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	return env
}

type AppConfig struct {
	Env  string `json:"env"`
	Host string `json:"host"`
	Port string `json:"port"`
}
