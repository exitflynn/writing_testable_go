package shortener

import (
	"errors"
	"hash/fnv"
	"net/url"
	"strconv"
	"time"
)

type URLMapping struct {
	ShortCode string
	LongURL   string
	CreatedAt time.Time
}

var (
	ErrInvalidURL    = errors.New("invalid URL")
	ErrInvalidScheme = errors.New("URL must use http or https")
)

func ValidateURL(rawURL string) error {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return ErrInvalidURL
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return ErrInvalidScheme
	}

	if parsed.Host == "" {
		return ErrInvalidURL
	}

	return nil
}

func GenerateShortCode(inputURL string) string {
	h := fnv.New64a()
	h.Write([]byte(inputURL))
	return strconv.FormatUint(h.Sum64(), 36)
}

func CreateMapping(longURL string) (*URLMapping, error) {
	if err := ValidateURL(longURL); err != nil {
		return nil, err
	}

	return &URLMapping{
		ShortCode: GenerateShortCode(longURL),
		LongURL:   longURL,
		CreatedAt: time.Now(),
	}, nil
}
