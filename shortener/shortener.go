package shortener

import (
	"time"
)

type URLMapping struct {
	ShortCode string
	LongURL   string
	CreatedAt time.Time
}

// ValidateURL checks if the given URL is valid.
// Returns an error if invalid.
func ValidateURL(rawURL string) error {
	// TODO: check if URL has valid scheme (http/https)
	// TODO: check if URL can be parsed
	panic("not implemented")
}

// GenerateShortCode creates a short code from a URL.
// The same URL should always produce the same code.
func GenerateShortCode(url string) string {
	// TODO: use hash/fnv to generate a hash
	// TODO: convert to base62 string (a-z, A-Z, 0-9)
	panic("not implemented")
}

// CreateMapping validates a URL and creates a mapping.
// Returns error if URL is invalid.
func CreateMapping(longURL string) (*URLMapping, error) {
	// TODO: validate the URL first
	// TODO: generate short code
	// TODO: return URLMapping with current time
	panic("not implemented")
}
