package model

import "payway/app/main/layer/status"

type CeberusAggregatorResponse struct {
	Username string                 `json:"username"`
	Status   status.CERBERUS_STATUS `json:"status"`
}
