package models

type (
	// UserCredentials API response
	UserCredentials struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}

	// Credential from API
	Credential struct {
		IDCredential           string `json:"id_credential"`
		IDUser                 string `json:"id_user"`
		IDEnvironment          string `json:"id_environment"`
		IDSite                 string `json:"id_site"`
		IDSiteOrganization     string `json:"id_site_organization"`
		IDSiteOrganizationType string `json:"id_site_organization_type"`
		IsAuthorized           int    `json:"is_authorized"`
		IsLocked               int    `json:"is_locked"`
		IsTwofa                int    `json:"is_twofa"`
		CanSync                int    `json:"can_sync"`
		Username               string `json:"username"`
		Code                   int    `json:"code"`
		DTAuthorized           int64  `json:"dt_authorized"`
		DTRefresh              int64  `json:"dt_refresh"`
		DTExecute              int64  `json:"dt_execute"`
		DTLastSync             int64  `json:"dt_last_sync"`
	}

	// CredentialJob information from API
	CredentialJob struct {
		IDCredential string `json:"id_credential"`
		IDJobUUID    string `json:"id_job_uuid"`
		IDJob        string `json:"id_job"`
		Username     string `json:"username"`
		WS           string `json:"ws"`
		Status       string `json:"status"`
		Twofa        string `json:"twofa"`
	}

	// Job credentials from API
	Job struct {
		Code    int     `json:"code"`
		Address string  `json:"address"`
		Twofa   []Twofa `json:"twofa"`
	}

	// Twofa job from API
	Twofa struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Label string `json:"label"`
	}
)
