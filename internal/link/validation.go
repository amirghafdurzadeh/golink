package link

import (
	"errors"
	"net/url"
)

var (
	errTargetURLIsRequired = errors.New("target_url is required")
	errInvalidTargetURL    = errors.New("invalid target_url")
	errBaseURLIsRequired   = errors.New("base_url is required")
	errCodeIsRequired      = errors.New("code is required")
)

func (r CreateRequest) Validate() error {
	if r.TargetURL == "" {
		return errTargetURLIsRequired
	}

	if r.BaseURL == "" {
		return errBaseURLIsRequired
	}

	parsedURL, err := url.ParseRequestURI(r.TargetURL)
	if err != nil ||
		parsedURL.Scheme == "" ||
		parsedURL.Host == "" {

		return errInvalidTargetURL
	}

	return nil
}
