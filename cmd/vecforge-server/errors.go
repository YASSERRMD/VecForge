package main

import (
	"log"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
	Err    error
}

func (e *AppError) Error() string {
	return e.Message
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Error: %v", err)
	
	code := http.StatusInternalServerError
	if e, ok := err.(*AppError); ok {
		code = e.Code
	}
	
	http.Error(w, err.Error(), code)
}
