package shortener

import (
	"context"
	"time"
)

// CacheClient is an interface for cache operations.
// This allows us to swap implementations for testing.
type CacheClient interface {
	Set(ctx context.Context, key, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

// URLServiceV2 is the improved, testable version.
// Dependencies are injected via constructor.
type URLServiceV2 struct {
	cache CacheClient
}

func NewURLServiceV2(cache CacheClient) *URLServiceV2 {
	return &URLServiceV2{cache: cache}
}

func (s *URLServiceV2) SaveURL(ctx context.Context, shortCode, longURL string) error {
	return s.cache.Set(ctx, shortCode, longURL, 24*time.Hour)
}

func (s *URLServiceV2) GetURL(ctx context.Context, shortCode string) (string, error) {
	return s.cache.Get(ctx, shortCode)
}
