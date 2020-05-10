package model

type URL struct {
	ShortURL string `json:"short"`
	LongURL string `json:"long"`
	IP string `json:"ip"`
	Count int `json:"count"`
}
