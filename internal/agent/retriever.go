package agent

import "github.com/YASSERRMD/VecForge/internal/db"

type Retriever struct {
	pool     *db.ConnectionPool
	embedder Embedder
	topK     int
}

func NewRetriever(pool *db.ConnectionPool, embedder Embedder, topK int) *Retriever {
	return &Retriever{pool: pool, embedder: embedder, topK: topK}
}

func (r *Retriever) Retrieve(query string) ([]db.Hit, error) {
	vec, err := r.embedder.Embed(query)
	if err != nil {
		return nil, err
	}
	
	providers := r.pool.GetAll()
	var allHits []db.Hit
	
	for _, p := range providers {
		hits, err := p.Search(vec, r.topK)
		if err != nil {
			continue
		}
		allHits = append(allHits, hits...)
	}
	
	return allHits, nil
}

func (r *Retriever) SetTopK(k int) { r.topK = k }
func (r *Retriever) TopK() int { return r.topK }
