package main

import (
	"strconv"
	"sync/atomic"
)

var (
	requestCount   uint64
	errorCount    uint64
	requestLatency uint64
)

func incRequests()   { atomic.AddUint64(&requestCount, 1) }
func incErrors()    { atomic.AddUint64(&errorCount, 1) }
func addLatency(ns int64) { atomic.AddUint64(&requestLatency, uint64(ns)) }

func prometheusMetrics() string {
	return `# HELP vecforge_requests_total Total requests
# TYPE vecforge_requests_total counter
vecforge_requests_total ` + strconv.FormatUint(atomic.LoadUint64(&requestCount), 10) + `
# HELP vecforge_errors_total Total errors
# TYPE vecforge_errors_total counter
vecforge_errors_total ` + strconv.FormatUint(atomic.LoadUint64(&errorCount), 10) + `
# HELP vecforge_request_latency_ns Total latency in nanoseconds
# TYPE vecforge_request_latency_ns counter
vecforge_request_latency_ns ` + strconv.FormatUint(atomic.LoadUint64(&requestLatency), 10)
}
