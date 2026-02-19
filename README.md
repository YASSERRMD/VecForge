# VecForge

Production Vector Gateway - Go + Rust FFI

## Features
- Multi-DB: Qdrant, Weaviate, Milvus
- Rust FFI for 42μs latency
- RAG Agent with query rewrite
- HTMX Dashboard
- Production ready: Rate limiting, CORS, Compression

## Quick Start
```bash
docker-compose -f docker/docker-compose.yml up -d
make run
# Open http://localhost:8080
```

## Deploy
```bash
flyctl deploy
```

## API
- GET /health
- POST /v1/search
- GET /v1/metrics

## License: MIT
