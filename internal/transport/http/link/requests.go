package link

import (
	"errors"
	"net/url"
)

var (
	ErrTargetURLIsRequired = errors.New("target_url is required")
	ErrInvalidTargetURL    = errors.New("invalid target_url")
	ErrBaseURLIsRequired   = errors.New("base_url is required")
	ErrCodeIsRequired      = errors.New("code is required")
)

type CreateLinkRequest struct {
	TargetURL  string `json:"target_url"`
	BaseURL    string `json:"base_url"`
	CustomCode string `json:"custom_code,omitempty"`
}

func (r CreateLinkRequest) Validate() error {
	if r.TargetURL == "" {
		return ErrTargetURLIsRequired
	}

	if r.BaseURL == "" {
		return ErrBaseURLIsRequired
	}

	parsedURL, err := url.ParseRequestURI(r.TargetURL)
	if err != nil ||
		parsedURL.Scheme == "" ||
		parsedURL.Host == "" {

		return ErrInvalidTargetURL
	}

	return nil
}
