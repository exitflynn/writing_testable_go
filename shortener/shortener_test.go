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
	// TODO: test that same input gives same output
	// TODO: test that different inputs give different outputs
}

func TestCreateMapping(t *testing.T) {
	// TODO: test valid URL returns mapping
	// TODO: test invalid URL returns error
}
