package examples

import (
	"encoding/json"
	"fmt"
	"log"

	resty "gopkg.in/resty.v1"
	"paybook.com/sync-code-samples/models"
)

// CreateSession user on Sync
func CreateSession() {
	var session models.Session
	baseURL := BaseURL + "sessions?pretty=1"

	resp, err := resty.SetDebug(Debug).R().
		SetHeader("Content-type", "application/json").
		SetHeader("Authorization", "API_KEY api_key="+APIkey).
		SetBody(`{"id_user":"` + IDUser + `"}`).
		SetResult(&response).
		SetError(&error).
		Post(baseURL)

	if err != nil {
		log.Fatal("Request error ", err, resp)
	}

	if error.Code != 0 {
		log.Fatal("API error code: ", error.Code, " Message: ", error.Message)
	}

	err = json.Unmarshal([]byte(response.Response), &session)
	if err != nil {
		log.Fatal("Unmarshal user session ", err)
	}

	fmt.Println("RID: ", response.Rid)
	fmt.Println("Code: ", response.Code)
	fmt.Println("Errors: ", response.Errors)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Message: ", response.Message)
	fmt.Println("Token: ", session.Token)
	fmt.Println("SYNC_TOKEN=" + session.Token)
}
