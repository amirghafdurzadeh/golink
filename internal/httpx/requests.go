package httpx

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var ErrInvalidRequest = errors.New("invalid request body")

func DecodeJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return ErrInvalidRequest
	}

	if decoder.Decode(&struct{}{}) != io.EOF {
		return ErrInvalidRequest
	}

	return nil
}
