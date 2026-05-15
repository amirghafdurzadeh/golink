package link

import (
	"errors"
	"net/url"
)

var (
	ErrTargetURLIsRequired = errors.New("target_url is required")
	ErrInvalidTargetURL    = errors.New("invalid target_url")
	ErrCodeIsRequired      = errors.New("code is required")
)

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

	return nil
}
