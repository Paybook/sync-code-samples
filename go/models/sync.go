package models

import (
	"encoding/json"
)

type (
	// Sync API response
	Sync struct {
		Rid      string          `json:"rid"`
		Code     int             `json:"code"`
		Status   bool            `json:"status"`
		Errors   string          `json:"errors"`
		Message  string          `json:"message"`
		Response json.RawMessage `json:"response"`
	}
)
