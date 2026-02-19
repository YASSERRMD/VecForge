# VecForge API

## Endpoints

### GET /health
Returns server health status.

### POST /v1/search
Search vector databases.

Request:
```json
{"q": "query", "providers": ["qdrant"], "limit": 10}
```

Response:
```json
{"hits": [], "query": "", "latency_us": 42, "provider": "fused"}
```

### GET /v1/metrics
Returns server metrics.
