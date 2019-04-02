package models

type (
	// Account catalog
	Account struct {
		IDAccount              string  `json:"id_account"`
		IDAccountType          string  `json:"id_account_type"`
		IDCurrency             string  `json:"id_currency"`
		IDUser                 string  `json:"id_user"`
		IDExternal             string  `json:"id_external"`
		IDCredential           string  `json:"id_credential"`
		IDSite                 string  `json:"id_site"`
		IDSiteOrganization     string  `json:"id_site_organization"`
		IDSiteOrganizationType string  `json:"id_site_organization_type"`
		ISDisable              int     `json:"is_disable"`
		Name                   string  `json:"name"`
		Number                 string  `json:"number"`
		Balance                float64 `json:"balance"`
		Currency               string  `json:"currency"`
		AccountType            string  `json:"account_type"`
		DTRefresh              int     `json:"dt_refresh"`
	}
)
