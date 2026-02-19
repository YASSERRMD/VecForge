package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Templates struct {
	layout    *template.Template
	dashboard *template.Template
	metrics   *template.Template
}

func LoadTemplates() (*Templates, error) {
	layout := template.Must(template.ParseFiles("web/templates/layout.html"))
	dash := template.Must(layout.ParseFiles("web/templates/dashboard.html"))
	met := template.Must(layout.ParseFiles("web/templates/metrics.html"))
	
	return &Templates{
		layout:    layout,
		dashboard: dash,
		metrics:   met,
	}, nil
}

func (t *Templates) RenderDashboard(w http.ResponseWriter, r *http.Request) {
	t.dashboard.Execute(w, nil)
}

func (t *Templates) RenderMetrics(w http.ResponseWriter, r *http.Request) {
	t.metrics.Execute(w, nil)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("web/static", r.URL.Path[1:])
	http.ServeFile(w, r, path)
}
