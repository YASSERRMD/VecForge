package db

import (
	"errors"
	"time"
)

type RetryConfig struct {
	MaxRetries int
	Delay      time.Duration
	MaxDelay   time.Duration
}

func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries: 3,
		Delay:      100 * time.Millisecond,
		MaxDelay:   1 * time.Second,
	}
}

func WithRetry(fn func() error, cfg *RetryConfig) error {
	var err error
	delay := cfg.Delay
	
	for i := 0; i < cfg.MaxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		
		if i < cfg.MaxRetries-1 {
			time.Sleep(delay)
			delay *= 2
			if delay > cfg.MaxDelay {
				delay = cfg.MaxDelay
			}
		}
	}
	
	return errors.New("max retries exceeded")
}

type RetryableError struct {
	Err error
}

func (e *RetryableError) Error() string {
	return e.Err.Error()
}

func IsRetryable(err error) bool {
	return true
}
