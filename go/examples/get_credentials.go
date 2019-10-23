package examples

import (
	"encoding/json"
	"fmt"
	"log"

	resty "gopkg.in/resty.v1"
	"paybook.com/sync-code-samples/models"
)

// GetCredentials from API
func GetCredentials() {
	var credentials []models.Credential
	baseURL := BaseURL + "credentials?pretty=1"

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

	err = json.Unmarshal([]byte(response.Response), &credentials)
	if err != nil {
		log.Fatal("Unmarshal users", err)
	}

	fmt.Println("RID: ", response.Rid)
	fmt.Println("Code: ", response.Code)
	fmt.Println("Errors: ", response.Errors)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Message: ", response.Message)

	fmt.Println("Credentials:")
	for key, value := range credentials {
		fmt.Println(key, ".- SYNC_IDCREDENTIAL="+value.IDCredential)
		fmt.Println(key, ".- ID User: ", value.IDUser)
		fmt.Println(key, ".- Username: ", value.Username)
		fmt.Println("----------------------------")
	}
}
