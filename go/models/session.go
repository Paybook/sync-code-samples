package models

type (
	// Session API response
	Session struct {
		IV       string `json:"iv"`
		Key      string `json:"key"`
		Token    string `json:"token"`
		DTCreate int    `json:"dt_create"`
		DTModify int    `json:"dt_modify"`
	}
)
