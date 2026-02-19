package agent

import (
	"sort"
	"github.com/YASSERRMD/VecForge/internal/db"
)

type Reranker struct{}

func NewReranker() *Reranker { return &Reranker{} }

func (r *Reranker) Rerank(hits []db.Hit, query string) []db.Hit {
	scores := make(map[string]float32)
	for _, h := range hits {
		scores[h.ID] = h.Score
	}
	
	sort.Slice(hits, func(i, j int) bool {
		return hits[i].Score > hits[j].Score
	})
	
	return hits
}

func (r *Reranker) RerankWithWeights(hits []db.Hit, weights map[string]float32) []db.Hit {
	for i := range hits {
		if w, ok := weights[hits[i].Provider]; ok {
			hits[i].Score *= w
		}
	}
	sort.Slice(hits, func(i, j int) bool {
		return hits[i].Score > hits[j].Score
	})
	return hits
}
