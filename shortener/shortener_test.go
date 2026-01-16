package shortener

import "testing"

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		// TODO: add test cases
		// hint: test valid URLs, invalid schemes, missing host
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
	// test that same input gives same output
	t.Run("deterministic", func(t *testing.T) {
		url := "https://example.com/test"
		code1 := GenerateShortCode(url)
		code2 := GenerateShortCode(url)
		if code1 != code2 {
			t.Errorf("same URL gave different codes: %s vs %s", code1, code2)
		}
	})
	// TODO: test that different inputs give different outputs
}

func TestCreateMapping(t *testing.T) {
	// TODO: test invalid URL returns error
	t.Run("invalid URL", func(t *testing.T) {
		_, err := CreateMapping("not-a-url")
		if err == nil {
			t.Error("expected error for invalid URL")
		}
	})
}
