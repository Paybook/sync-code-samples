package examples

import (
	"fmt"
	"log"

	resty "gopkg.in/resty.v1"
)

// DeleteUser user on Sync
func DeleteUser() {
	// var credentialsResponse models.Credential
	baseURL := BaseURL + "users/" + IDUser + "?pretty=1"

	resp, err := resty.SetDebug(Debug).R().
		SetHeader("Content-type", "application/json").
		SetHeader("Authorization", "API_KEY api_key="+APIkey).
		SetResult(&response).
		SetError(&error).
		Delete(baseURL)

	if err != nil {
		log.Fatal("Request error ", err, resp)
	}

	if error.Code != 0 {
		log.Fatal("API error code: ", error.Code, " Message: ", error.Message)
	}

	fmt.Println("RID: ", response.Rid)
	fmt.Println("Code: ", response.Code)
	fmt.Println("Errors: ", response.Errors)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Message: ", response.Message)
}
