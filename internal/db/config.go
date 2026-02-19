package db

import "time"

type Config struct {
	URL      string
	Timeout  time.Duration
	MaxConns int
}

func DefaultConfig() *Config {
	return &Config{
		URL:      "http://localhost:6333",
		Timeout:  5 * time.Second,
		MaxConns: 10,
	}
}
