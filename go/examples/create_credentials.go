package examples

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	resty "gopkg.in/resty.v1"
	"paybook.com/sync-code-samples/models"
)

var (
	isTokenSet bool
)

// CreateCredentials user on Sync
func CreateCredentials() {
	var credentialJob models.CredentialJob
	baseURL := BaseURL + "credentials/pulls?pretty=1"

	resp, err := resty.SetDebug(Debug).R().
		SetHeader("Content-type", "application/json").
		SetHeader("Authorization", "TOKEN token="+Token).
		SetBody(`{
			"id_site":"` + IDSite + `",
			"credentials":` + Credentials + `
			}`).
		SetResult(&response).
		SetError(&error).
		Post(baseURL)

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

	err = json.Unmarshal([]byte(response.Response), &credentialJob)
	if err != nil {
		log.Fatal("Unmarshal credentialJob error ", err)
	}

	fmt.Println("StatusURL: ", credentialJob.Status)
	if len(credentialJob.Twofa) > 1 {
		fmt.Println("Twofa: ", credentialJob.Twofa)
	}

	getJobStatus(credentialJob.Status)
}

func getJobStatus(statusURL string) {
	var jobs []models.Job
	time.Sleep(5 * time.Second)
	respStatus, errStatus := resty.SetDebug(Debug).R().
		SetHeader("Content-type", "application/json").
		SetHeader("Authorization", "TOKEN token="+Token).
		SetBody(`{
			"id_site":"` + IDSite + `",
			"credentials":` + Credentials + `
			}`).
		SetResult(&response).
		SetError(&error).
		Get(statusURL)

	if errStatus != nil {
		log.Fatal("Request error ", errStatus, respStatus)
	}

	if error.Code != 0 {
		log.Fatal("API error code: ", error.Code, " Message: ", error.Message)
	}

	err := json.Unmarshal([]byte(response.Response), &jobs)
	if err != nil {
		log.Fatal("Unmarshal jobs error ", err)
	}

	for _, value := range jobs {
		switch value.Code {
		case 100:
			fmt.Println("Status : Register")
			fmt.Println("Message : API register a new job")
			break
		case 101:
			fmt.Println("Status : Starting")
			fmt.Println("Message : Sync got job information to start working")
			break
		case 102:
			fmt.Println("Status : Running")
			fmt.Println("Message : Sync is running (login successful)")
			break
		case 103:
			fmt.Println("Status : TokenReceived")
			fmt.Println("Message : Sync received the token")
			break
		case 200:
			fmt.Println("Status : Finish")
			log.Fatal(response.Message)
			break
		case 201:
			fmt.Println("Status : Pending")
			log.Fatal(response.Message)
			break
		case 202:
			fmt.Println("Status : NoTransactions")
			log.Fatal(response.Message)
			break
		case 203:
			fmt.Println("Status : PartialTransactions")
			log.Fatal(response.Message)
			break
		case 204:
			fmt.Println("Status : Incomplete")
			log.Fatal(response.Message)
			break
		case 206:
			fmt.Println("Status : NoAccounts")
			log.Fatal(response.Message)
			break
		case 400:
			fmt.Println("Status : Credentials error")
			log.Fatal(response.Message)
			break
		case 401:
			fmt.Println("Status : Unauthorized")
			log.Fatal(response.Message)
			break
		case 403:
			fmt.Println("Status : Invalidtoken")
			log.Fatal(response.Message)
			break
		case 405:
			fmt.Println("Status : Locked")
			log.Fatal(response.Message)
			break
		case 406:
			fmt.Println("Status : Conflict")
			log.Fatal(response.Message)
			break
		case 408:
			fmt.Println("Status : UserAction")
			log.Fatal(response.Message)
			break
		case 409:
			fmt.Println("Status : WrongSite")
			log.Fatal(response.Message)
			break
		case 410:
			fmt.Println("Status : Waiting")
			fmt.Println("Message : Waiting for two-fa")
			if !isTokenSet {
				getTwofa(value.Address, value.Twofa)
				getJobStatus(statusURL)
			}
			break
		case 411:
			fmt.Println("Status : TwofaTimeout")
			log.Fatal(response.Message)
			break
		case 500:
			fmt.Println("Status : Error")
			log.Fatal(response.Message)
			break
		case 501:
			fmt.Println("Status : Unavailable")
			log.Fatal(response.Message)
			break
		case 504:
			fmt.Println("Status : ScriptTimeout")
			log.Fatal(response.Message)
			break
		case 509:
			fmt.Println("Status : UndergoingMaintenance")
			log.Fatal(response.Message)
			break
		default:
			fmt.Println(response.Message)
			log.Fatal("Unknown status ", response.Code)
			break
		}
		fmt.Println("----------------------------")
	}
}

func getTwofa(address string, twofa []models.Twofa) {
	for _, value := range twofa {
		fmt.Println(value.Label)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		token := `{"twofa":{"` + value.Name + `":"` + text + `"}}`

		resp, err := resty.SetDebug(Debug).R().
			SetHeader("Content-type", "application/json").
			SetHeader("Authorization", "TOKEN token="+Token).
			SetBody(token).
			SetResult(&response).
			SetError(&error).
			Post(address)

		if err != nil {
			log.Fatal("Request error ", err, resp)
		}

		if error.Code != 0 {
			log.Fatal("API error code: ", error.Code, " Message: ", error.Message)
		}

		isTokenSet = true
	}
}
