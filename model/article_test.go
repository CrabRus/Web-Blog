package model

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestArticleJSONMarshalUnmarshal(t *testing.T) {
	original := Article{
		ID:        1,
		Title:     "Test title",
		Content:   "Test content",
		Published: "true",
		Author:    "admin",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var decoded Article
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if decoded != original {
		t.Fatalf("expected %+v, got %+v", original, decoded)
	}
}

func TestArticleAuthorOmitEmpty(t *testing.T) {
	a := Article{
		ID:        1,
		Title:     "Test",
		Content:   "Content",
		Published: "true",
		Author:    "",
	}

	data, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	jsonStr := string(data)

	if strings.Contains(jsonStr, `"author"`) {
		t.Fatal("author field should be omitted when empty (omitempty)")
	}
}

func TestArticleJSONTable(t *testing.T) {
	tests := []struct {
		name    string
		article Article
	}{
		{
			name: "full article",
			article: Article{
				ID:        1,
				Title:     "Title",
				Content:   "Content",
				Published: "true",
				Author:    "admin",
			},
		},
		{
			name: "without author",
			article: Article{
				ID:        2,
				Title:     "Title",
				Content:   "Content",
				Published: "false",
				Author:    "",
			},
		},
		{
			name: "empty content",
			article: Article{
				ID:        3,
				Title:     "Only title",
				Content:   "",
				Published: "false",
				Author:    "user",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.article)
			if err != nil {
				t.Fatalf("marshal error: %v", err)
			}

			var decoded Article
			err = json.Unmarshal(data, &decoded)
			if err != nil {
				t.Fatalf("unmarshal error: %v", err)
			}

			if decoded.ID != tt.article.ID {
				t.Errorf("ID mismatch: %d != %d", decoded.ID, tt.article.ID)
			}

			if decoded.Title != tt.article.Title {
				t.Errorf("Title mismatch: %s != %s", decoded.Title, tt.article.Title)
			}

			if decoded.Content != tt.article.Content {
				t.Errorf("Content mismatch: %s != %s", decoded.Content, tt.article.Content)
			}

			if decoded.Published != tt.article.Published {
				t.Errorf("Published mismatch: %s != %s", decoded.Published, tt.article.Published)
			}

			if decoded.Author != tt.article.Author {
				t.Errorf("Author mismatch: %s != %s", decoded.Author, tt.article.Author)
			}
		})
	}
}

func TestArticleInvalidJSON(t *testing.T) {
	invalidJSON := `{"id":1,"title":123,"content":true}`

	var a Article
	err := json.Unmarshal([]byte(invalidJSON), &a)

	if err == nil {
		t.Fatal("expected unmarshal error for invalid JSON, got nil")
	}
}

func TestArticleZeroValue(t *testing.T) {
	var a Article

	if a.ID != 0 {
		t.Error("expected ID = 0")
	}

	if a.Title != "" {
		t.Error("expected empty Title")
	}

	if a.Content != "" {
		t.Error("expected empty Content")
	}

	if a.Published != "" {
		t.Error("expected empty Published")
	}

	if a.Author != "" {
		t.Error("expected empty Author")
	}
}
