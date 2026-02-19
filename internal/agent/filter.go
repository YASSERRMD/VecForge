package agent

type Filter interface {
	Filter(hits []Hit) []Hit
}

type ScoreFilter struct {
	minScore float32
}

func NewScoreFilter(min float32) *ScoreFilter {
	return &ScoreFilter{minScore: min}
}

func (f *ScoreFilter) Filter(hits []Hit) []Hit {
	var filtered []Hit
	for _, h := range hits {
		if h.Score >= f.minScore {
			filtered = append(filtered, h)
		}
	}
	return filtered
}

type ProviderFilter struct {
	allowed []string
}

func NewProviderFilter(providers []string) *ProviderFilter {
	return &ProviderFilter{allowed: providers}
}

func (f *ProviderFilter) Filter(hits []Hit) []Hit {
	set := make(map[string]bool)
	for _, p := range f.allowed {
		set[p] = true
	}
	
	var filtered []Hit
	for _, h := range hits {
		if set[h.Source] {
			filtered = append(filtered, h)
		}
	}
	return filtered
}

type DedupeFilter struct{}

func NewDedupeFilter() *DedupeFilter { return &DedupeFilter{} }

func (f *DedupeFilter) Filter(hits []Hit) []Hit {
	seen := make(map[string]bool)
	var filtered []Hit
	for _, h := range hits {
		if !seen[h.Content] {
			seen[h.Content] = true
			filtered = append(filtered, h)
		}
	}
	return filtered
}
