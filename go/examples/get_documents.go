package examples

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	resty "gopkg.in/resty.v1"
	"paybook.com/sync-code-samples/models"
)

// GetDocuments from API
func GetDocuments() {
	var documents []models.Document
	baseURL := BaseURL + "documents?pretty=1"

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

	err = json.Unmarshal([]byte(response.Response), &documents)
	if err != nil {
		log.Fatal("Unmarshal documents", err)
	}

	fmt.Println("RID: ", response.Rid)
	fmt.Println("Code: ", response.Code)
	fmt.Println("Errors: ", response.Errors)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Message: ", response.Message)

	fmt.Println("Documents:")
	for key, value := range documents {
		fmt.Println(key, ".- IDDocument: ", value.IDDocument)
		fmt.Println(key, ".- File name: ", value.File)
		fmt.Println(key, ".- SYNC_IDDocument="+value.IDDocument)
		fmt.Println("Writing file")
		writeFile(value.File, value.Content)
		fmt.Println("----------------------------")
	}
}

func writeFile(filename, content string) {
	data, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		log.Fatal("Error on decoding base64 string:", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal("Error on write file:", err)
	}
}
