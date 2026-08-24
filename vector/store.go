package vector

import "context"

// The actual document - pgvector will store
type Document struct {
	ID        string
	Content   string
	Metadata  map[string]string
	Embedding []float32
}

// Result returned by querying the pgvector db object
type Result struct {
	Document
	Score float32 //How close they on scale of 0 to 1 (above 0.8 great and less than 0.4 noise)
}

// Contract defined for db operations
type Store interface {
	Upsert(ctx context.Context, docs []Document) error
	Query(ctx context.Context, embedding []float32, topK int) ([]Result, error)
	Delete(ctx context.Context, ids []string) error
	DeleteBySource(ctx context.Context, source string) error
	Close() error // resource cleanup

}
