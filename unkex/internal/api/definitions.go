package api

type SubscribeRequest struct {
	Email string `json:"email"`
}

type RateResponse struct {
	"base" string  `json:"base"`
	"usd"  float32 `json:"usd"`
	"uah"  float32 `json:"uah"`
}