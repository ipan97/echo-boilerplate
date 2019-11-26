package response

import "time"

type LoginResponse struct {
	AccessToken string    `json:"access_token"`
	Expire      time.Time `json:"expire"`
}
