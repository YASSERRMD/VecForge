package main

import "os"

type ServerConfig struct {
	Port         string
	TemplateDir  string
	StaticDir    string
	ReadTimeout  int
	WriteTimeout int
}

func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:         getEnv("PORT", "8080"),
		TemplateDir:  getEnv("TEMPLATE_DIR", "web/templates"),
		StaticDir:    getEnv("STATIC_DIR", "web/static"),
		ReadTimeout:  15,
		WriteTimeout: 15,
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
