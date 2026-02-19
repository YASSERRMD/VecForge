package main

import (
	"context"
	"log"
	"time"
)

type TraceID string

func newTraceID() TraceID {
	return TraceID("trace-" + time.Now().Format("20060102150405"))
}

func withTrace(ctx context.Context, id TraceID) context.Context {
	return context.WithValue(ctx, "trace_id", id)
}

func getTraceID(ctx context.Context) TraceID {
	if id, ok := ctx.Value("trace_id").(TraceID); ok {
		return id
	}
	return ""
}

func logWithTrace(id TraceID, msg string) {
	log.Printf("[%s] %s", id, msg)
}
