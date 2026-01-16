package bad_example

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// This is BAD CODE - it's tightly coupled and hard to test.
// Can you spot the problems?

type URLService struct {
	// nothing here - we create clients inside methods!
}

func NewURLService() *URLService {
	return &URLService{}
}

func (s *URLService) SaveURL(shortCode, longURL string) error {
	// Problem 1: creating redis client inside the method
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	// Problem 2: hardcoded configuration
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Problem 3: no way to inject a different implementation
	return client.Set(ctx, shortCode, longURL, 24*time.Hour).Err()
}

func (s *URLService) GetURL(shortCode string) (string, error) {
	// Same problems repeated
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Get(ctx, shortCode).Result()
}

// How would you test this?
// - Need a real Redis running
// - Can't mock the redis client
// - Can't control what it returns
// - Tests are slow and flaky
