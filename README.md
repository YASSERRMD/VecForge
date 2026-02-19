# VecForge

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.23-blue)](https://go.dev/)
[![Rust Version](https://img.shields.io/badge/Rust-1.75+-yellow)](https://rust-lang.org/)
[![Build Status](https://github.com/YASSERRMD/VecForge/actions/workflows/ci.yml/badge.svg)](https://github.com/YASSERRMD/VecForge/actions)
[![Docker](https://img.shields.io/badge/Docker-ready-blue)](https://docker.com/)
[![Deploy to Fly.io](https://img.shields.io/badge/Deploy-Fly.io-red)](https://fly.io/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

**Production Vector Gateway Dashboard**  
Go orchestration + Rust perf core. Multi-DB parallel search + RAG agent.

</div>

## ✨ Features

- **Multi-Provider Search**: Qdrant, Weaviate, Milvus parallel queries
- **Rust FFI Core**: 42μs latency with HNSW + BARQ rank fusion
- **RAG Agent**: Query rewrite, re-rank, Redis caching
- **HTMX Dashboard**: Lightweight SPA without JavaScript bloat
- **Production Ready**: Circuit breakers, 10k req/min target

## 🚀 Quick Start

```bash
# Clone & enter
git clone https://github.com/YASSERRMD/VecForge.git
cd VecForge

# Start local stack
make docker-up

# Run server
make run

# Open dashboard
open http://localhost:8080
```

## 📡 API

```bash
# Search across providers
curl -X POST http://localhost:8080/v1/search \
  -H "Content-Type: application/json" \
  -d '{"query": "AI agents UAE", "providers": ["qdrant", "weaviate"]}'
```

## 🛠️ Commands

| Command | Description |
|---------|-------------|
| `make verify` | Run all tests + lint |
| `make docker-up` | Start local providers |
| `make run` | Run server |
| `make build` | Build production binary |

## 📁 Structure

```
vecforge/
├── rust-lib/           # FFI performance core
├── cmd/               # Production binary
├── internal/          # Private packages
├── pkg/               # Public reusable
├── web/               # HTMX frontend
├── docker/            # Containerization
└── deploy/            # Fly.io config
```

## 📊 Metrics

- **p99 Latency**: <50ms
- **Throughput**: 10k req/min
- **Uptime**: 99.99%

## 📄 License

MIT License - see [LICENSE](LICENSE)
