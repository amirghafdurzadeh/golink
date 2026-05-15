package link

type CreateLinkResponse struct {
	Code      string `json:"code"`
	ShortURL  string `json:"short_url"`
	TargetURL string `json:"target_url"`
}

type GetLinkResponse struct {
	Code      string `json:"code"`
	TargetURL string `json:"target_url"`
}
