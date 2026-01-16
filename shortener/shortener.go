package shortener

import (
	"hash/fnv"
	"strconv"
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
func GenerateShortCode(inputURL string) string {
	h := fnv.New64a()
	h.Write([]byte(inputURL))
	return strconv.FormatUint(h.Sum64(), 36)
}

// CreateMapping validates a URL and creates a mapping.
// Returns error if URL is invalid.
func CreateMapping(longURL string) (*URLMapping, error) {
	// TODO: validate the URL first
	// TODO: generate short code
	// TODO: return URLMapping with current time
	panic("not implemented")
}
