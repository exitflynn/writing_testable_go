package shortener

import "errors"

var ErrNotFound = errors.New("mapping not found")

// URLStore defines how we persist URL mappings.
// Different implementations: memory, redis, postgres, etc.
type URLStore interface {
	Save(mapping *URLMapping) error
	Get(shortCode string) (*URLMapping, error)
}

// MemoryStore is a simple in-memory implementation.
type MemoryStore struct {
	data map[string]*URLMapping
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[string]*URLMapping)}
}

func (m *MemoryStore) Save(mapping *URLMapping) error {
	m.data[mapping.ShortCode] = mapping
	return nil
}

func (m *MemoryStore) Get(shortCode string) (*URLMapping, error) {
	mapping, ok := m.data[shortCode]
	if !ok {
		return nil, ErrNotFound
	}
	return mapping, nil
}
