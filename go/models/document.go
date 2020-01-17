package models

type (
	// Document catalog
	Document struct {
		IDDocument             string `json:"id_document"`
		IDDocumentStatus       string `json:"id_document_status"`
		IDDocumentType         string `json:"id_document_type"`
		IDExternal             string `json:"id_external"`
		IDSite                 string `json:"id_site"`
		IDSiteOrganization     string `json:"id_site_organization"`
		IDSiteOrganizationType string `json:"id_site_organization_type"`
		File                   string `json:"file"`
		Content                string `json:"content"`
	}
)
