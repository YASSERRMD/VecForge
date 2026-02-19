package db

import "errors"

var (
	ErrProviderNotFound = errors.New("provider not found")
	ErrNoProviders     = errors.New("no providers available")
	ErrSearchTimeout   = errors.New("search timeout")
	ErrCircuitOpen    = errors.New("circuit breaker open")
	ErrInvalidQuery   = errors.New("invalid query")
	ErrNotImplemented = errors.New("not implemented")
)

type SearchError struct {
	Provider string
	Err      error
}

func (e *SearchError) Error() string {
	return e.Provider + ": " + e.Err.Error()
}

func (e *SearchError) Unwrap() error {
	return e.Err
}

type ProviderError struct {
	Name string
	Code int
	Msg  string
}

func (e *ProviderError) Error() string {
	return e.Msg
}
