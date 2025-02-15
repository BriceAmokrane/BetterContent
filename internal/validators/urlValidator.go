package validators

import (
	"fmt"
	"net/http"
	"regexp"
)

type URLValidator interface {
	ValidateURL(url string) error
	IsAlive(url string) error
}

type DefaultURLValidator struct{}

func NewURLValidator() URLValidator {
	return &DefaultURLValidator{}
}

func (v *DefaultURLValidator) ValidateURL(url string) error {
	if !v.isValidFormat(url) {
		return fmt.Errorf("invalid URL format: %s", url)
	}
	return nil
}

func (v *DefaultURLValidator) isValidFormat(url string) bool {
	patterns := []string{
		"^http://\\w+(\\.\\w+)+.*$",
		"^https://\\w+(\\.\\w+)+.*$",
	}

	for _, pattern := range patterns {
		matched, err := regexp.Match(pattern, []byte(url))
		if err == nil && matched {
			return true
		}
	}
	return false
}

func (v *DefaultURLValidator) IsAlive(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to reach URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("URL returned status code: %d", resp.StatusCode)
	}
	return nil
}
