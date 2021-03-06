package examples

import (
	"encoding/json"
	"fmt"
	"log"

	resty "gopkg.in/resty.v1"
	"paybook.com/sync-code-samples/models"
)

// GetAccounts from API
func GetAccounts() {
	var accounts []models.Account
	baseURL := BaseURL + "accounts?pretty=1&id_credential=" + IDCredential

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

	err = json.Unmarshal([]byte(response.Response), &accounts)
	if err != nil {
		log.Fatal("Unmarshal accounts", err)
	}

	fmt.Println("RID: ", response.Rid)
	fmt.Println("Code: ", response.Code)
	fmt.Println("Errors: ", response.Errors)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Message: ", response.Message)

	fmt.Println("Accounts:")
	for key, value := range accounts {
		fmt.Println(key, ".- ID Account: ", value.IDAccount)
		fmt.Println(key, ".- Name: ", value.Name)
		fmt.Println(key, ".- Number: ", value.Number)
		fmt.Println(key, ".- Balance: ", value.Balance)
		fmt.Println(key, ".- SYNC_IDACCOUNT="+value.IDAccount)
		fmt.Println("----------------------------")
	}
}
