package utils

import "testing"

func TestValidateTitle(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		wantErr bool
	}{
		{
			name:    "valid title",
			title:   "My first article",
			wantErr: false,
		},
		{
			name:    "empty title",
			title:   "",
			wantErr: true,
		},
		{
			name:    "too long title",
			title:   string(make([]byte, 101)), // 101 chars
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTitle(tt.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTitle(%q) error = %v, wantErr = %v", tt.title, err, tt.wantErr)
			}
		})
	}
}

func TestValidateContent(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name:    "valid content",
			content: "This is article content",
			wantErr: false,
		},
		{
			name:    "empty content",
			content: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateContent(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateContent(%q) error = %v, wantErr = %v", tt.content, err, tt.wantErr)
			}
		})
	}
}

func TestValidatePublished(t *testing.T) {
	tests := []struct {
		name      string
		published string
		wantErr   bool
	}{
		{
			name:      "valid date",
			published: "2025-02-17",
			wantErr:   false,
		},
		{
			name:      "empty date",
			published: "",
			wantErr:   true,
		},
		{
			name:      "wrong format",
			published: "17-02-2025",
			wantErr:   true,
		},
		{
			name:      "random string",
			published: "hello",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePublished(tt.published)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePublished(%q) error = %v, wantErr = %v", tt.published, err, tt.wantErr)
			}
		})
	}
}
