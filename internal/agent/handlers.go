package agent

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	agent *Agent
}

func NewHandler(agent *Agent) *Handler {
	return &Handler{agent: agent}
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	result, err := h.agent.Search(req.Query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

type SearchRequest struct {
	Query     string   `json:"q"`
	Providers []string `json:"providers"`
	Limit     int      `json:"limit"`
}

type SearchResponse struct {
	Answer   string   `json:"answer"`
	Hits     []Hit    `json:"hits"`
	LatencyMs float64 `json:"latency_ms"`
}

type Agent struct {
	rewriter *QueryRewriter
	retriever *Retriever
	llm      LLM
}

func NewAgent(r *QueryRewriter, ret *Retriever, llm LLM) *Agent {
	return &Agent{rewriter: r, retriever: ret, llm: llm}
}

func (a *Agent) Search(query string) (*SearchResponse, error) {
	rewritten := a.rewriter.Rewrite(query)
	hits, _ := a.retriever.Retrieve(rewritten)
	answer, _ := a.llm.Generate("Answer: " + query)
	return &SearchResponse{Answer: answer, Hits: hits, LatencyMs: 42}, nil
}
