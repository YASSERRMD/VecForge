package agent

import "sort"

type Ranker struct{}

func NewRanker() *Ranker { return &Ranker{} }

func (r *Ranker) Fuse(results []db.SearchResult, k int) []db.Hit {
	scores := make(map[string]float32)
	
	for _, res := range results {
		for rank, hit := range res.Hits {
			weight := 1.0 / float32(k+rank+1)
			scores[hit.ID] += weight
		}
	}
	
	var fused []db.Hit
	for id, score := range scores {
		fused = append(fused, db.Hit{ID: id, Score: score})
	}
	
	sort.Slice(fused, func(i, j int) bool {
		return fused[i].Score > fused[j].Score
	})
	
	if len(fused) > k {
		fused = fused[:k]
	}
	return fused
}
