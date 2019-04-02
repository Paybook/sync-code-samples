package models

type (
	// User API response
	User struct {
		IDUser     string `json:"id_user"`
		IDExternal string `json:"id_external"`
		Name       string `json:"name"`
		DTCreate   int    `json:"dt_create"`
		DTModify   int    `json:"dt_modify"`
	}
)
