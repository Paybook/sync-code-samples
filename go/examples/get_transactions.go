package examples

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	resty "gopkg.in/resty.v1"
	"paybook.com/sync-code-samples/models"
)

// GetTransactions from API
func GetTransactions() {
	var transactions []models.Transaction
	baseURL := BaseURL + "transactions?pretty=1&id_credential=" + IDCredential + "&id_account=" + IDAccount + "&limit=" + strconv.Itoa(TransactionLimit)

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

	err = json.Unmarshal([]byte(response.Response), &transactions)
	if err != nil {
		log.Fatal("Unmarshal transactions", err)
	}

	fmt.Println("RID: ", response.Rid)
	fmt.Println("Code: ", response.Code)
	fmt.Println("Errors: ", response.Errors)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Message: ", response.Message)

	fmt.Println("Transactions:")
	for key, value := range transactions {

		fmt.Println(key, ".- Description: ", value.Description)
		fmt.Println(key, ".- Amount: ", value.Amount)
		fmt.Println(key, ".- Date: ", time.Unix(value.DTTransaction, 0))
		fmt.Println("----------------------------")
	}
}
