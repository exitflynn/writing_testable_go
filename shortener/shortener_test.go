package shortener

import "testing"

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{"valid https", "https://example.com", false},
		{"valid http", "http://example.com/path", false},
		{"valid with port", "https://localhost:8080/api", false},
		{"ftp scheme", "ftp://files.example.com", true},
		{"no scheme", "example.com", true},
		// Missing: empty host case like "https://"
		// This is the bug we'll discover later!
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateURL(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateURL(%q) error = %v, wantErr %v", tt.url, err, tt.wantErr)
			}
		})
	}
}

func TestGenerateShortCode(t *testing.T) {
	t.Run("deterministic", func(t *testing.T) {
		url := "https://example.com/test"
		code1 := GenerateShortCode(url)
		code2 := GenerateShortCode(url)
		if code1 != code2 {
			t.Errorf("same URL gave different codes: %s vs %s", code1, code2)
		}
	})

	t.Run("different URLs give different codes", func(t *testing.T) {
		code1 := GenerateShortCode("https://example.com/a")
		code2 := GenerateShortCode("https://example.com/b")
		if code1 == code2 {
			t.Error("different URLs gave same code")
		}
	})
}

func TestCreateMapping(t *testing.T) {
	t.Run("valid URL", func(t *testing.T) {
		mapping, err := CreateMapping("https://example.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if mapping.ShortCode == "" {
			t.Error("short code is empty")
		}
		if mapping.LongURL != "https://example.com" {
			t.Errorf("long URL = %q, want %q", mapping.LongURL, "https://example.com")
		}
	})

	t.Run("invalid URL", func(t *testing.T) {
		_, err := CreateMapping("not-a-url")
		if err == nil {
			t.Error("expected error for invalid URL")
		}
	})
}
