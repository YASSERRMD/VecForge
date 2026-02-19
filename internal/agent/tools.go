package agent

import "context"

type Tool interface {
	Name() string
	Description() string
	Execute(ctx context.Context, args map[string]interface{}) (interface{}, error)
}

type SearchTool struct {
	pool *dbPool
}

func NewSearchTool(p *dbPool) *SearchTool {
	return &SearchTool{pool: p}
}

func (t *SearchTool) Name() string        { return "search" }
func (t *SearchTool) Description() string { return "Search vector database" }

func (t *SearchTool) Execute(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	query, _ := args["query"].(string)
	topK, _ := args["top_k"].(int)
	if topK == 0 {
		topK = 10
	}
	return []Hit{}, nil
}

type dbPool struct{}

type ToolRegistry struct {
	tools map[string]Tool
}

func NewToolRegistry() *ToolRegistry {
	return &ToolRegistry{tools: make(map[string]Tool)}
}

func (r *ToolRegistry) Register(t Tool) {
	r.tools[t.Name()] = t
}

func (r *ToolRegistry) Get(name string) Tool {
	return r.tools[name]
}

func (r *ToolRegistry) List() []Tool {
	var list []Tool
	for _, t := range r.tools {
		list = append(list, t)
	}
	return list
}
