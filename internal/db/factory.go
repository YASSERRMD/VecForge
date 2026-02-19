package db

type ProviderFactory func(url string) (Provider, error)

var providers = make(map[string]ProviderFactory)

func RegisterProvider(name string, factory ProviderFactory) {
	providers[name] = factory
}

func CreateProvider(name, url string) (Provider, error) {
	factory, ok := providers[name]
	if !ok {
		return nil, ErrUnknownProvider
	}
	return factory(url)
}

type ErrUnknownProviderType struct{}

func (e *ErrUnknownProviderType) Error() string {
	return "unknown provider type"
}

var ErrUnknownProvider = &ErrUnknownProviderType{}

func init() {
	RegisterProvider("qdrant", func(url string) (Provider, error) {
		return NewQdrantClient(url), nil
	})
	RegisterProvider("weaviate", func(url string) (Provider, error) {
		return NewWeaviateClient(url), nil
	})
	RegisterProvider("milvus", func(url string) (Provider, error) {
		return NewMilvusClient(url), nil
	})
}
