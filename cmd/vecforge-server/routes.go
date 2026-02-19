package main

import (
	"net/http"
)

func setupRoutes(tmpl *Templates) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			tmpl.RenderDashboard(w, r)
		} else if r.URL.Path == "/metrics" {
			tmpl.RenderMetrics(w, r)
		}
	})
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	
	http.HandleFunc("/v1/search", handleSearchResults)
	http.HandleFunc("/v1/metrics", handleMetrics)
}
