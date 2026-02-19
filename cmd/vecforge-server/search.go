package main

import (
	"encoding/json"
	"net/http"
	"text/template"
)

func handleSearchResults(w http.ResponseWriter, r *http.Request) {
	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	results := []map[string]interface{}{
		{"id": "doc1", "score": 0.95, "provider": "qdrant", "content": "Result 1"},
		{"id": "doc2", "score": 0.87, "provider": "weaviate", "content": "Result 2"},
	}
	
	tmpl := `<div class="hit" data-id="{{.id}}">
		<div class="flex justify-between">
			<span class="font-medium">{{.id}}</span>
			<span class="text-green-600">{{.score}}</span>
		</div>
		<p class="text-sm text-gray-500">{{.provider}}</p>
	</div>`
	
	w.Header().Set("Content-Type", "text/html")
	template.Must(template.New("result").Parse(tmpl)).Execute(w, results[0])
}
