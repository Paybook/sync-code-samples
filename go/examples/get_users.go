package examples

import (
	"encoding/json"
	"fmt"
	"log"

	resty "gopkg.in/resty.v1"
	"paybook.com/sync-code-samples/models"
)

// GetUsers from API
func GetUsers() {
	var users []models.User
	baseURL := BaseURL + "users?pretty=1"

	resp, err := resty.SetDebug(Debug).R().
		SetHeader("Content-type", "application/json").
		SetHeader("Authorization", "API_KEY api_key="+APIkey).
		SetResult(&response).
		SetError(&error).
		Get(baseURL)

	if err != nil {
		log.Fatal("Request error ", err, resp)
	}

	if error.Code != 0 {
		log.Fatal("API error code: ", error.Code, " Message: ", error.Message)
	}

	err = json.Unmarshal([]byte(response.Response), &users)
	if err != nil {
		log.Fatal("Unmarshal users", err)
	}

	fmt.Println("RID: ", response.Rid)
	fmt.Println("Code: ", response.Code)
	fmt.Println("Errors: ", response.Errors)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Message: ", response.Message)

	fmt.Println("Users:")
	for key, value := range users {
		fmt.Println(key, ".- ID User: ", value.IDUser)
		fmt.Println(key, ".- ID External: ", value.IDExternal)
		fmt.Println(key, ".- Name: ", value.Name)
		fmt.Println(key, ".- SYNC_IDUSER="+value.IDUser)
		fmt.Println("----------------------------")
	}
}
