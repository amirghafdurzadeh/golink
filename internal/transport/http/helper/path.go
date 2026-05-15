package helper

import (
	"fmt"
	"net/http"
)

func MustPathValue(r *http.Request, key string) (string, error) {
	v := r.PathValue(key)
	if v == "" {
		return "", fmt.Errorf("%s is required", key)
	}
	return v, nil
}
