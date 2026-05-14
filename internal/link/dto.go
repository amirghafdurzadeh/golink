package link

type CreateRequest struct {
	TargetURL  string `json:"target_url"`
	BaseURL    string `json:"base_url"`
	CustomCode string `json:"custom_code,omitempty"`
}

type CreateResponse struct {
	Code      string `json:"code"`
	ShortURL  string `json:"short_url"`
	TargetURL string `json:"target_url"`
}

type GetResponse struct {
	Code      string `json:"code"`
	TargetURL string `json:"target_url"`
}

type GetStatsResponse struct {
	Code   string `json:"code"`
	Clicks int64  `json:"clicks"`
}
