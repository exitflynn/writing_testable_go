package shortener

import (
	"context"
	"errors"
	"testing"
)

func TestURLServiceV2_SaveAndGet(t *testing.T) {
	mock := NewMockCacheClient()
	svc := NewURLServiceV2(mock)
	ctx := context.Background()

	err := svc.SaveURL(ctx, "abc", "https://example.com")
	if err != nil {
		t.Fatalf("SaveURL failed: %v", err)
	}

	got, err := svc.GetURL(ctx, "abc")
	if err != nil {
		t.Fatalf("GetURL failed: %v", err)
	}

	if got != "https://example.com" {
		t.Errorf("got %q, want %q", got, "https://example.com")
	}

	if mock.SetCalls != 1 {
		t.Errorf("SetCalls = %d, want 1", mock.SetCalls)
	}
	if mock.GetCalls != 1 {
		t.Errorf("GetCalls = %d, want 1", mock.GetCalls)
	}
}

func TestURLServiceV2_CacheError(t *testing.T) {
	mock := NewMockCacheClient()
	mock.SetErr = errors.New("connection refused")

	svc := NewURLServiceV2(mock)
	ctx := context.Background()

	err := svc.SaveURL(ctx, "abc", "https://example.com")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
