package agent

import (
	"strings"
	"unicode"
)

type QueryRewriter struct{}

func NewQueryRewriter() *QueryRewriter { return &QueryRewriter{} }

func (r *QueryRewriter) Rewrite(query string) string {
	query = strings.TrimSpace(query)
	query = r.lowercase(query)
	query = r.removePunctuation(query)
	query = r.expandAbbreviations(query)
	query = r.addSynonyms(query)
	return query
}

func (r *QueryRewriter) lowercase(s string) string {
	return strings.ToLower(s)
}

func (r *QueryRewriter) removePunctuation(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, s)
}

func (r *QueryRewriter) expandAbbreviations(s string) string {
	abbrevs := map[string]string{
		"ai": "artificial intelligence",
		"ml": "machine learning",
		"nlp": "natural language processing",
	}
	for abbr, full := range abbrevs {
		s = strings.ReplaceAll(s, abbr, full)
	}
	return s
}

func (r *QueryRewriter) addSynonyms(s string) string {
	synonyms := map[string][]string{
		"ai":     {"artificial intelligence", "machine intelligence"},
		"search": {"find", "retrieve", "lookup"},
	}
	for word, syns := range synonyms {
		if strings.Contains(s, word) {
			s = s + " " + strings.Join(syns, " ")
		}
	}
	return s
}
