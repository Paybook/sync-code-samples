package examples

import (
	"encoding/json"
	"fmt"
	"log"

	resty "gopkg.in/resty.v1"
	"paybook.com/sync-code-samples/models"
)

// GetSitesByOrganization from API
func GetSitesByOrganization() {
	var sites []models.Sites
	baseURL := BaseURL + "catalogues/organizations/sites?pretty=1"

	if len(Token) == 0 {
		log.Fatal("Token not set")
	}

	resp, err := resty.SetDebug(Debug).R().
		SetHeader("Content-type", "application/json").
		SetHeader("Authorization", "TOKEN token="+Token).
		SetResult(&response).
		SetError(&error).
		Get(baseURL)

	if err != nil {
		log.Fatal("Request error ", err, resp)
	}

	if error.Code != 0 {
		log.Fatal("API error code: ", error.Code, " Message: ", error.Message)
	}

	err = json.Unmarshal([]byte(response.Response), &sites)
	if err != nil {
		log.Fatal("Unmarshal sites catalog", err)
	}

	fmt.Println("RID: ", response.Rid)
	fmt.Println("Code: ", response.Code)
	fmt.Println("Errors: ", response.Errors)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Message: ", response.Message)

	fmt.Println("Sites:")
	for _, value := range sites {
		for keySite, valueSite := range value.Sites {
			fmt.Println(keySite, ".- ID Site: ", valueSite.IDSite)
			fmt.Println(keySite, ".- Name: ", valueSite.Name)
			fmt.Println(keySite, ".- Version: ", valueSite.Version)
			fmt.Printf("%d .- Credentials: \n%s\n", keySite, valueSite.Credentials)
			fmt.Println(keySite, ".- SYNC_IDSITE="+valueSite.IDSite)
			fmt.Println("----------------------------")
		}
	}
}
