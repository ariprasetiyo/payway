package model

import (
	"payway/app/main/model"
)

type TokenResponse struct {
	model.Response
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   string `json:"expiresIn"`
}
