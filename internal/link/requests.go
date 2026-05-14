package link

type CreateRequest struct {
	TargetURL  string `json:"target_url"`
	BaseURL    string `json:"base_url"`
	CustomCode string `json:"custom_code,omitempty"`
}
