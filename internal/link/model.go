package link

import (
	"time"
)

type Link struct {
	Code      string    `json:"code"`
	TargetURL string    `json:"target_url"`
	CreatedAt time.Time `json:"created_at"`
}
