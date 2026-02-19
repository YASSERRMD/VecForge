package main

import (
	"log"
	"os"
)

func initLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

func logRequest(method, path, ip string, status int, latency int64) {
	log.Printf("%s %s %s %d %dms", method, path, ip, status, latency)
}

func logError(err error) {
	log.Printf("ERROR: %v", err)
}
