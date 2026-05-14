package link

type CreateResponse struct {
	Code      string `json:"code"`
	ShortURL  string `json:"short_url"`
	TargetURL string `json:"target_url"`
}

type GetResponse struct {
	Code      string `json:"code"`
	TargetURL string `json:"target_url"`
}
