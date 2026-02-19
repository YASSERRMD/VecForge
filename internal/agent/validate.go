package agent

import "strings"

type Validator struct{}

func NewValidator() *Validator { return &Validator{} }

func (v *Validator) ValidateQuery(query string) error {
	if strings.TrimSpace(query) == "" {
		return ErrEmptyQuery
	}
	if len(query) > 1000 {
		return ErrQueryTooLong
	}
	return nil
}

func (v *Validator) ValidateResult(hits []Hit) error {
	if len(hits) == 0 {
		return ErrNoResults
	}
	return nil
}

var (
	ErrEmptyQuery   = &ValidationError{msg: "query is empty"}
	ErrQueryTooLong = &ValidationError{msg: "query exceeds 1000 chars"}
	ErrNoResults    = &ValidationError{msg: "no results found"}
)

type ValidationError struct{ msg string }

func (e *ValidationError) Error() string { return e.msg }
