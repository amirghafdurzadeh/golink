package link

import "strings"

func buildCode(customCode string, length int) string {
	if customCode != "" {
		return customCode
	}

	return generateShortCode(length)
}

func buildCreateResponse(baseURL string, link Link) CreateResponse {
	shortURL := strings.TrimRight(baseURL, "/") + "/r/" + link.Code

	return CreateResponse{
		Code:      link.Code,
		ShortURL:  shortURL,
		TargetURL: link.TargetURL,
	}
}
