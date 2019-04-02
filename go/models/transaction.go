package models

type (
	// Transaction catalog
	Transaction struct {
		IDTransaction          string  `json:"id_transaction"`
		IDAccount              string  `json:"id_account"`
		IDAccountType          string  `json:"id_account_type"`
		IDCurrency             string  `json:"id_currency"`
		IDUser                 string  `json:"id_user"`
		IDExternal             string  `json:"id_external"`
		IDCredential           string  `json:"id_credential"`
		IDSite                 string  `json:"id_site"`
		IDSiteOrganization     string  `json:"id_site_organization"`
		IDSiteOrganizationType string  `json:"id_site_organization_type"`
		ISAccountDisable       int     `json:"is_account_disable"`
		ISDisable              int     `json:"is_disable"`
		ISDeleted              int     `json:"is_deleted"`
		ISPending              int     `json:"is_pending"`
		Description            string  `json:"description"`
		Amount                 float64 `json:"amount"`
		Currency               string  `json:"currency"`
		AccountType            string  `json:"account_type"`
		DTTransaction          int64   `json:"dt_transaction"`
		DTRefresh              int     `json:"dt_refresh"`
	}
)
