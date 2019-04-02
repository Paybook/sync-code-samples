package models

import "encoding/json"

type (
	// Site catalog
	Site struct {
		IDSite      string            `json:"id_site"`
		IDSiteType  string            `json:"id_site_type"`
		ISBusiness  int               `json:"is_business"`
		ISPersonal  int               `json:"is_personal"`
		Version     int               `json:"version"`
		Name        string            `json:"name"`
		Credentials []json.RawMessage `json:"credentials"`
		Input       []json.RawMessage `json:"input"`
		Endpoint    string            `json:"endpoint"`
	}

	// Sites catalog
	Sites struct {
		IDSiteOrganization     string `json:"id_site_organization"`
		IDSiteOrganizationType string `json:"id_site_organization_type"`
		IDCountry              string `json:"id_country"`
		Name                   string `json:"name"`
		Avatar                 string `json:"avatar"`
		SmallCover             string `json:"small_cover"`
		Cover                  string `json:"cover"`
		Sites                  []Site `json:"sites"`
	}
)
