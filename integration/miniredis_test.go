package integration

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/workshop/writing-testable-go/shortener"
)

func TestWithMiniredis(t *testing.T) {
	// Start miniredis - no Docker needed!
	s := miniredis.RunT(t)

	// Create cache pointing to miniredis
	cache := shortener.NewRedisCache(s.Addr())
	svc := shortener.NewURLServiceV2(cache)
	ctx := context.Background()

	// Test
	err := svc.SaveURL(ctx, "mini", "https://miniredis.io")
	if err != nil {
		t.Fatalf("SaveURL failed: %v", err)
	}

	got, err := svc.GetURL(ctx, "mini")
	if err != nil {
		t.Fatalf("GetURL failed: %v", err)
	}

	if got != "https://miniredis.io" {
		t.Errorf("got %q, want %q", got, "https://miniredis.io")
	}

	// Miniredis lets you inspect state directly
	val, _ := s.Get("mini")
	if val != "https://miniredis.io" {
		t.Errorf("miniredis has %q, want %q", val, "https://miniredis.io")
	}
}

func TestMiniredis_CanSimulateErrors(t *testing.T) {
	s := miniredis.RunT(t)

	// Save something
	client := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client.Set(context.Background(), "key", "value", 0)

	// Now close the server to simulate failure
	s.Close()

	// Next call should fail
	_, err := client.Get(context.Background(), "key").Result()
	if err == nil {
		t.Error("expected error after closing miniredis")
	}
}
