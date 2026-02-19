package models

type SearchReq struct {
	Query     string   `json:"q"`
	Providers []string `json:"providers"`
	Limit     int      `json:"limit"`
}

type SearchResp struct {
	Hits      []Hit  `json:"hits"`
	Query     string `json:"query"`
	LatencyUs int64  `json:"latency_us"`
	Provider  string `json:"provider"`
}

type Hit struct {
	ID       string      `json:"id"`
	Score    float64     `json:"score"`
	Provider string      `json:"provider"`
}
