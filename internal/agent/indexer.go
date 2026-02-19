package agent

type Indexer struct {
	embedder Embedder
}

func NewIndexer(embedder Embedder) *Indexer {
	return &Indexer{embedder: embedder}
}

func (i *Indexer) Index(docs []Document) (map[string][]float32, error) {
	embeddings := make(map[string][]float32)
	
	for _, doc := range docs {
		emb, err := i.embedder.Embed(doc.Content)
		if err != nil {
			return nil, err
		}
		embeddings[doc.Content] = emb
	}
	
	return embeddings, nil
}

func (i *Indexer) IndexSingle(doc Document) ([]float32, error) {
	return i.embedder.Embed(doc.Content)
}
