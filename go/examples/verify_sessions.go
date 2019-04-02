package examples

import (
	"fmt"
	"log"

	resty "gopkg.in/resty.v1"
)

// VerifySessions from API
func VerifySessions() {
	if len(Token) == 0 {
		log.Fatal("Token not set")
	}
	baseURL := BaseURL + "sessions/" + Token + "/verify?pretty=1"

	resp, err := resty.SetDebug(Debug).R().
		SetHeader("Content-type", "application/json").
		SetResult(&response).
		SetError(&error).
		Get(baseURL)

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
