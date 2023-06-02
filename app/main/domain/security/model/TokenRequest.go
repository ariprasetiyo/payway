package model

type TokenRequest struct {
	GrantType      string         `json:"grantType"`
	AdditionalInfo AdditionalInfo `json:"additionalInfo"`
}

type AdditionalInfo struct {
}
