package model

import "time"

type RequestModel struct {
	RequestId   string `json:"requestId"`
	Type        string `json:"type"`
	RequestTime int64  `json:"requestTime"`
}

func (rm *RequestModel) RequestModelDefault() {
	if rm.RequestTime == 0 {
		rm.RequestTime = time.Now().UnixMilli()
	}
}
