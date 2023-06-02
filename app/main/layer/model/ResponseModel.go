package model

type ResponseModel struct {
	RequestId     string `json:"requestId"`
	Type          string `json:"type"`
	Status        int    `json:"status"`
	StatusMessage string `json:"statusMessage"`
	Body          any    `json:"body"`
}
