package agent

type Config struct {
	TopK          int
	MaxTokens     int
	Temperature   float32
	RewriteQuery  bool
	UseRerank    bool
	CacheEnabled  bool
}

func DefaultConfig() *Config {
	return &Config{
		TopK:         10,
		MaxTokens:    1000,
		Temperature:   0.7,
		RewriteQuery: true,
		UseRerank:   true,
		CacheEnabled: true,
	}
}

func (c *Config) Validate() error {
	if c.TopK < 1 || c.TopK > 100 {
		return ErrInvalidConfig
	}
	if c.Temperature < 0 || c.Temperature > 2 {
		return ErrInvalidConfig
	}
	return nil
}

var ErrInvalidConfig = &ConfigError{msg: "invalid config"}

type ConfigError struct{ msg string }

func (e *ConfigError) Error() string { return e.msg }
