package examples

import (
	"encoding/json"
	"fmt"
	"log"

	resty "gopkg.in/resty.v1"
	"paybook.com/sync-code-samples/models"
)

// CreateUser user on Sync
func CreateUser() {
	var user models.User
	baseURL := BaseURL + "users?pretty=1"

	resp, err := resty.SetDebug(Debug).R().
		SetHeader("Content-type", "application/json").
		SetHeader("Authorization", "API_KEY api_key="+APIkey).
		SetBody(`{"name":"` + UserName + `"}`).
		SetResult(&response).
		SetError(&error).
		Post(baseURL)

	if err != nil {
		log.Fatal("Request error ", err, resp)
	}

	if error.Code != 0 {
		log.Fatal("API error code: ", error.Code, " Message: ", error.Message)
	}

	err = json.Unmarshal([]byte(response.Response), &user)
	if err != nil {
		log.Fatal("Unmarshal user error ", err)
	}

	fmt.Println("RID: ", response.Rid)
	fmt.Println("Code: ", response.Code)
	fmt.Println("Errors: ", response.Errors)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Message: ", response.Message)
	fmt.Println("User name: ", user.Name)
	fmt.Println("User id: ", user.IDUser)
}
