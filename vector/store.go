package vector

import "context"

// Document is one piece of source content that the RAG system can index and
// later retrieve as context for an LLM.
//
// A source file is usually split into smaller chunks before it becomes a
// Document. Smaller chunks make semantic search more precise and keep the
// amount of context sent to the LLM manageable. The store keeps the original
// text together with its embedding and metadata so a search result is useful
// both to the model and to the person tracing where an answer came from.
type Document struct {
	// ID gives this chunk a stable identity. It allows ingestion to update the
	// same chunk later and lets callers delete it without removing other chunks.
	ID string

	// Content is the human-readable text returned as candidate context. The LLM
	// reads this field when it creates an answer grounded in the source material.
	Content string

	// Metadata records where the chunk came from, such as a file, title, URL, or
	// page number. It does not usually guide the answer directly, but it makes
	// results traceable and can support filtering or citations.
	Metadata map[string]string

	// Embedding is the vector representation of Content produced by an embedding
	// model. pgvector compares this vector with a query vector to find chunks
	// that are semantically related, even when they use different wording.
	Embedding []float32
}

// Result is a Document found during semantic search, plus the similarity score
// calculated by the vector database. RAG uses that score to rank the chunks so
// the most relevant context can be sent to the LLM first.
type Result struct {
	Document
	// Score describes how closely the document embedding matches the query
	// embedding. Its exact meaning depends on the distance metric, but the
	// retrieval layer uses it to compare and limit candidate context.
	Score float32
}

// Store is the small contract the RAG pipeline needs from a vector database.
// Keeping this interface independent of pgvector makes indexing and retrieval easier to test
// and leaves the rest of the application free from database-specific details.
type Store interface {
	// Upsert adds new document chunks or replaces chunks with the same IDs. This
	// makes ingestion repeatable: re-indexing changed source content does not
	// create duplicate chunks in the RAG knowledge base.
	Upsert(ctx context.Context, docs []Document) error

	// Query finds up to topK chunks whose embeddings are closest to the query
	// embedding. The application then uses their Content and Metadata as
	// candidate context when building the LLM prompt.
	Query(ctx context.Context, embedding []float32, topK int) ([]Result, error)

	// Delete removes specific chunks by ID. This is useful when a source is
	// deleted or when selected chunks are no longer valid.
	Delete(ctx context.Context, ids []string) error

	// DeleteBySource removes every chunk belonging to one source before it is
	// re-indexed. Without this step, chunks from an older version could remain
	// searchable and cause the LLM to answer with stale information.
	DeleteBySource(ctx context.Context, source string) error

	// Close releases database connections and other resources when the
	// application shuts down, allowing the store to finish cleanly.
	Close() error
}
