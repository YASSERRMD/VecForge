package db

import "fmt"

type Config struct {
	URL      string  `validate:"required,url"`
	Timeout  int     `validate:"min=1,max=60"`
	MaxConns int     `validate:"min=1,max=100"`
}

func (c *Config) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("URL is required")
	}
	if c.Timeout < 1 || c.Timeout > 60 {
		return fmt.Errorf("timeout must be between 1 and 60")
	}
	return nil
}

func (c *Config) WithDefaults() *Config {
	if c.Timeout == 0 {
		c.Timeout = 5
	}
	if c.MaxConns == 0 {
		c.MaxConns = 10
	}
	return c
}
