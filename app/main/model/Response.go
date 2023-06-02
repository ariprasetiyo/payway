package model

type Response struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	AdditionalInfo  any    `json:"additionalInfo"`
}
