package utils

import (
	"errors"
	"time"
)

func ValidateTitle(title string) error {
	if title == "" || len(title) > 100 {
		return errors.New("title required and <100 chars")
	}
	return nil
}

func ValidateContent(content string) error {
	if content == "" {
		return errors.New("content required")
	}
	return nil
}

func ValidatePublished(published string) error {
	if _, err := time.Parse("2006-01-02", published); err != nil {
		return errors.New("invalid date format YYYY-MM-DD")
	}
	return nil
}
