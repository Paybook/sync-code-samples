package models

type (
	// Attachment catalog
	Attachment struct {
		IDAttachment     string `json:"id_attachment"`
		IDAccount        string `json:"id_account"`
		IDUser           string `json:"id_user"`
		IDExternal       string `json:"id_external"`
		IDTransaction    string `json:"id_transaction"`
		IDAttachmentType string `json:"id_attachment_type"`
		IsValid          int    `json:"is_valid"`
		File             string `json:"file"`
		Mime             string `json:"mime"`
		URL              string `json:"url"`
		DTRefresh        int    `json:"dt_refresh"`
	}
)
