# VecForge Architecture

## Overview
```
┌─────────────┐     ┌─────────────┐
│   Client    │────▶│  Go Server  │
└─────────────┘     └──────┬──────┘
                          │
           ┌──────────────┼──────────────┐
           ▼              ▼              ▼
    ┌──────────┐  ┌──────────┐  ┌──────────┐
    │  Qdrant  │  │ Weaviate │  │  Milvus  │
    └──────────┘  └──────────┘  └──────────┘
           │              │              │
           └──────────────┼──────────────┘
                          ▼
                   ┌──────────┐
                   │ BARQ     │
                   │ Rank     │
                   │ Fusion   │
                   └──────────┘
```

## Components
- Go Server: HTTP handling, middleware
- Rust FFI: Vector search, rank fusion
- DB Providers: Qdrant, Weaviate, Milvus
- RAG Agent: Query rewrite, rerank, cache

## Performance
- p99 < 50ms
- 10k req/min
- 99.99% uptime
