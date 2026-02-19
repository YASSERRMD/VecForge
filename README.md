# VecForge

Production Vector Gateway - Go + Rust FFI

## Features
- Multi-DB: Qdrant, Weaviate, Milvus
- Rust FFI for 42μs latency
- RAG Agent with query rewrite
- Circuit breaker, retry, cache
- HTMX dashboard

## Architecture
```
User → POST /v1/search → Go Server → Rust FFI → BARQ Fusion → Redis Cache
```

## Quick Start
```bash
docker-compose -f docker/docker-compose.yml up -d
make run
curl -X POST http://localhost:8080/v1/search -d '{"q":"AI"}'
```
