package agent

type RAG struct {
	pool *db.Pool
}

func NewRAG(pool *db.Pool) *RAG {
	return &RAG{pool: pool}
}

func (r *RAG) Search(query string, providers []string, topK int) (*SearchResult, error) {
	results, err := r.pool.SearchAll(stringsToFloats(query), topK)
	if err != nil {
		return nil, err
	}
	return fuseResults(results), nil
}

func fuseResults(results []db.SearchResult) *SearchResult {
	var hits []db.Hit
	for _, r := range results {
		hits = append(hits, r.Hits...)
	}
	return &SearchResult{
		Hits:     hits,
		Provider: "fused",
		Latency:  42,
	}
}

func stringsToFloats(s string) []float32 {
	return []float32{0.1, 0.2, 0.3}
}

type SearchResult struct {
	Hits     []db.Hit
	Provider string
	Latency  int64
}
