package examples

import (
	"encoding/json"
	"fmt"
	"log"

	resty "gopkg.in/resty.v1"
	"paybook.com/sync-code-samples/models"
)

// DownloadAttachments from API
func DownloadAttachments() {
	var attachments []models.Attachment
	baseURL := BaseURL + "attachments?pretty=1&id_credential=" + IDCredential + "&id_account=" + IDAccount + "&id_transaction=" + IDTransaction

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

	err = json.Unmarshal([]byte(response.Response), &attachments)
	if err != nil {
		log.Fatal("Unmarshal attachments", err)
	}

	fmt.Println("RID: ", response.Rid)
	fmt.Println("Code: ", response.Code)
	fmt.Println("Errors: ", response.Errors)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Message: ", response.Message)

	fmt.Println("Attachments:")
	for key, value := range attachments {

		fmt.Println(key, ".- IDAttachment: ", value.IDAttachment)
		fmt.Println(key, ".- File: ", value.File)
		fmt.Println(key, ".- Mime: ", value.Mime)
		fmt.Println(key, ".- URL: ", value.URL)
		DownloadFile(value.File, value.URL)
		fmt.Println("----------------------------")
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) {
	resp, err := resty.SetDebug(Debug).R().
		SetHeader("Content-type", "application/json").
		SetHeader("Authorization", "TOKEN token="+Token).
		SetError(&error).
		SetOutput(filepath).
		Get(BaseURL + url)

	if err != nil {
		log.Fatal("Request download error ", err, resp)
	}
}
