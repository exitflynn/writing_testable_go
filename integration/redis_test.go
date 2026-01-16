//go:build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	tcredis "github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/workshop/writing-testable-go/shortener"
)

func TestWithRealRedis(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()

	// Start a real Redis container
	redisC, err := tcredis.Run(ctx, "redis:7-alpine")
	if err != nil {
		t.Fatalf("failed to start redis: %v", err)
	}
	defer func() {
		if err := testcontainers.TerminateContainer(redisC); err != nil {
			t.Logf("failed to terminate: %v", err)
		}
	}()

	// Get connection string
	endpoint, err := redisC.Endpoint(ctx, "")
	if err != nil {
		t.Fatalf("failed to get endpoint: %v", err)
	}

	// Create real Redis cache
	cache := shortener.NewRedisCache(endpoint)
	svc := shortener.NewURLServiceV2(cache)

	// Test with real Redis
	err = svc.SaveURL(ctx, "test123", "https://example.com")
	if err != nil {
		t.Fatalf("SaveURL failed: %v", err)
	}

	got, err := svc.GetURL(ctx, "test123")
	if err != nil {
		t.Fatalf("GetURL failed: %v", err)
	}

	if got != "https://example.com" {
		t.Errorf("got %q, want %q", got, "https://example.com")
	}

	// Verify directly with redis client
	client := redis.NewClient(&redis.Options{Addr: endpoint})
	ttl, _ := client.TTL(ctx, "test123").Result()
	if ttl < 23*time.Hour {
		t.Errorf("TTL = %v, want ~24h", ttl)
	}
}
