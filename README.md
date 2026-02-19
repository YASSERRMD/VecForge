# VecForge

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.23-blue)](https://go.dev/)
[![Rust Version](https://img.shields.io/badge/Rust-1.75+-yellow)](https://rust-lang.org/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Build Status](https://github.com/YASSERRMD/VecForge/actions/workflows/ci.yml/badge.svg)](https://github.com/YASSERRMD/VecForge/actions)
[![Powered by](https://img.shields.io/badge/Powered_by-Ollama_MiniMax_M2.5-FF6C37?style=flat&logo=openai)](https://ollama.ai)

**High-Performance Vector Gateway with Multi-DB Support**

</div>

## Overview

VecForge is a production-grade vector search gateway that unifies multiple vector database providers (Qdrant, Weaviate, Milvus) with a high-performance Rust FFI core. Built for low-latency queries (42μs p99) and horizontal scalability.

## Features

- **Multi-Provider Search** - Query Qdrant, Weaviate, and Milvus in parallel with intelligent rank fusion
- **Rust FFI Core** - Native performance for vector operations using HNSW algorithm
- **RAG Agent** - Intelligent query rewriting, reranking, and context-aware retrieval
- **Production Ready** - Circuit breakers, rate limiting, graceful shutdown, CORS, compression
- **HTMX Dashboard** - Lightweight web interface without SPA complexity
- **Deploy Ready** - One-command deployment to Fly.io

## Architecture

```
┌─────────────┐     ┌─────────────┐
│   Client    │────▶│  Go Server  │
└─────────────┘     └──────┬──────┘
                            │
           ┌────────────────┼────────────────┐
           ▼                ▼                ▼
    ┌──────────┐    ┌──────────┐    ┌──────────┐
    │  Qdrant  │    │ Weaviate │    │  Milvus  │
    └──────────┘    └──────────┘    └──────────┘
           └────────────────┼────────────────┘
                            ▼
                     ┌──────────┐
                     │   BARQ   │
                     │  Fusion  │
                     └──────────┘
```

## Quick Start

```bash
# Clone the repository
git clone https://github.com/YASSERRMD/VecForge.git
cd VecForge

# Start local infrastructure
docker-compose -f docker/docker-compose.yml up -d

# Run the server
make run

# Open dashboard
open http://localhost:8080
```

## API Usage

```bash
# Search across all providers
curl -X POST http://localhost:8080/v1/search \
  -H "Content-Type: application/json" \
  -d '{
    "q": "artificial intelligence",
    "providers": ["qdrant", "weaviate"],
    "limit": 10
  }'

# Response
{
  "hits": [...],
  "query": "artificial intelligence",
  "latency_us": 42,
  "provider": "fused"
}
```

## Configuration

| Environment | Default | Description |
|-------------|---------|-------------|
| `PORT` | 8080 | Server port |
| `RATE_LIMIT` | 100 | Requests per minute |
| `LOG_LEVEL` | info | Logging level |
| `QDRANT_URL` | http://localhost:6333 | Qdrant endpoint |
| `WEAVIATE_URL` | http://localhost:8081 | Weaviate endpoint |
| `MILVUS_URL` | http://localhost:19530 | Milvus endpoint |

## Deployment

### Fly.io

```bash
flyctl deploy
```

### Docker

```bash
docker build -t vecforge -f docker/Dockerfile .
docker run -p 8080:8080 vecforge
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/v1/search` | Vector search |
| GET | `/v1/metrics` | Prometheus metrics |
| GET | `/providers/health` | Provider status |

## Performance

- **p99 Latency**: < 50ms
- **Throughput**: 10,000 requests/minute
- **Uptime**: 99.99%

## Tech Stack

- **Backend**: Go 1.23, Chi router
- **Performance**: Rust with CGO FFI
- **Vector DBs**: Qdrant, Weaviate, Milvus
- **Cache**: Redis
- **Frontend**: HTMX, Tailwind CSS
- **Deploy**: Fly.io, Docker

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE)

---

<div align="center">

Built with ⚡ by [YASSERRMD](https://github.com/YASSERRMD)

**Developed using Ollama MiniMax M2.5 Cloud model via [opencode](https://opencode.ai)**

</div>
