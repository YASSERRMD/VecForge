package main

import "os"

type EnvConfig struct {
	Port         string
	LogLevel     string
	RateLimit    int
	ReadTimeout  int
	WriteTimeout int
}

func LoadEnv() *EnvConfig {
	return &EnvConfig{
		Port:         getEnv("PORT", "8080"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		RateLimit:    parseInt("RATE_LIMIT", 100),
		ReadTimeout:  parseInt("READ_TIMEOUT", 15),
		WriteTimeout: parseInt("WRITE_TIMEOUT", 15),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func parseInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return 0
}
