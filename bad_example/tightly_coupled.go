package bad_example

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type URLService struct {
	// initialise with a redis client interface
}

func NewURLService() *URLService {
	return &URLService{}
}

func (s *URLService) SaveURL(shortCode, longURL string) error {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Set(ctx, shortCode, longURL, 24*time.Hour).Err()
}

func (s *URLService) GetURL(shortCode string) (string, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Get(ctx, shortCode).Result()
}
