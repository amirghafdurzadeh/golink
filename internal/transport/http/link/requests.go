package link

import (
	"errors"
	"net/url"
	"regexp"
)

var (
	ErrTargetURLIsRequired = errors.New("target_url is required")
	ErrInvalidTargetURL    = errors.New("invalid target_url")
	ErrInvalidCustomCode   = errors.New("custom_code must be alphanumeric and between 3 and 10 characters")
)

var codeRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,}$`)

type CreateLinkRequest struct {
	TargetURL  string `json:"target_url"`
	CustomCode string `json:"custom_code,omitempty"`
}

func (r CreateLinkRequest) Validate() error {
	if r.TargetURL == "" {
		return ErrTargetURLIsRequired
	}

	parsedURL, err := url.ParseRequestURI(r.TargetURL)
	if err != nil ||
		parsedURL.Scheme == "" ||
		parsedURL.Host == "" {

		return ErrInvalidTargetURL
	}

	if r.CustomCode != "" && !codeRegex.MatchString(r.CustomCode) {
		return ErrInvalidCustomCode
	}

	return nil
}
