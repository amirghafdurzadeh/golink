package link

import (
	"net/http"
)

func getCodeFromRequest(r *http.Request) (string, error) {
	code := r.PathValue("code")

	if code == "" {
		return "", errCodeIsRequired
	}

	return code, nil
}
