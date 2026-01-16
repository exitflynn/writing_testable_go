package shortener

import (
	"context"
	"errors"
	"time"
)

var ErrCacheMiss = errors.New("cache miss")

// MockCacheClient is a hand-written mock for testing.
type MockCacheClient struct {
	data     map[string]string
	SetErr   error
	GetErr   error
	SetCalls int
	GetCalls int
}

func NewMockCacheClient() *MockCacheClient {
	return &MockCacheClient{data: make(map[string]string)}
}

func (m *MockCacheClient) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	m.SetCalls++
	if m.SetErr != nil {
		return m.SetErr
	}
	m.data[key] = value
	return nil
}

func (m *MockCacheClient) Get(ctx context.Context, key string) (string, error) {
	m.GetCalls++
	if m.GetErr != nil {
		return "", m.GetErr
	}
	val, ok := m.data[key]
	if !ok {
		return "", ErrCacheMiss
	}
	return val, nil
}
