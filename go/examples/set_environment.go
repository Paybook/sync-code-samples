package examples

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"paybook.com/sync-code-samples/models"
)

var (
	// BaseURL for all API requests
	BaseURL string

	// APIkey for all API requests
	APIkey string

	// Credentials used on requests
	Credentials string

	// IDUser for requests
	IDUser string

	// UserName for requests
	UserName string

	// IDSite for requests
	IDSite string

	// IDCredential for requests
	IDCredential string

	// IDAccount for requests
	IDAccount string

	// IDTransaction for requests
	IDTransaction string

	// TransactionLimit for requests
	TransactionLimit int

	// Token used on requests
	Token string

	// Debug option on request
	Debug bool

	response models.Sync
	error    models.Sync
)

// Setup environment
func Setup(baseURL string, apiKey string) {
	re := regexp.MustCompile("^[a-f0-9]{32}$")
	errAPIkey := re.MatchString(apiKey)
	if !errAPIkey {
		log.Fatal("Invalid API key format")
	}

	APIkey = apiKey

	re = regexp.MustCompile("^(http://www.|https://www.|http://|https://)?[a-z0-9]+([-.]{1}[a-z0-9]+)*.[a-z]{2,5}(:[0-9]{1,5})?(/.*)?$")
	errBASEURL := re.MatchString(baseURL)
	if !errBASEURL {
		log.Fatal("Invalid BASE URL")
	}

	BaseURL = baseURL
}

// SetToken for requests
func SetToken(token string) {
	re := regexp.MustCompile("^[a-f0-9]{64}$")
	errToken := re.MatchString(token)
	if !errToken {
		log.Fatal("Invalid token")
	}

	Token = token
}

// SetCredentials for requests
func SetCredentials(credentials string) {
	var userCredentials models.UserCredentials

	err := json.Unmarshal([]byte(credentials), &userCredentials)
	if err != nil {
		log.Fatal("Invalid credentials", err)
	}
	Credentials = credentials
}

// SetIDUser for requests
func SetIDUser(idUser string) {
	re := regexp.MustCompile("^[a-f0-9]{24}$")
	errIDUser := re.MatchString(idUser)
	if !errIDUser {
		log.Fatal("Invalid id_user")
	}

	IDUser = idUser
}

// SetUserName for requests
func SetUserName(userName string) {
	if len(userName) < 5 {
		log.Fatal("Invalid username must have at least 5 chars")
	}
	UserName = userName
}

// SetIDSite for requests
func SetIDSite(idSite string) {
	re := regexp.MustCompile("^[a-f0-9]{24}$")
	errIDSite := re.MatchString(idSite)
	if !errIDSite {
		log.Fatal("Invalid id_site")
	}

	IDSite = idSite
}

// SetIDCredential for requests
func SetIDCredential(idCredential string) {
	re := regexp.MustCompile("^[a-f0-9]{24}$")
	errIDCredential := re.MatchString(idCredential)
	if !errIDCredential {
		log.Fatal("Invalid id_credential")
	}

	IDCredential = idCredential
}

// SetIDAccount for requests
func SetIDAccount(idAccount string) {
	re := regexp.MustCompile("^[a-f0-9]{24}$")
	errIDAccount := re.MatchString(idAccount)
	if !errIDAccount {
		log.Fatal("Invalid id_account")
	}

	IDAccount = idAccount
}

// SetTransactionLimit for requests
func SetTransactionLimit(transactionLimit int) {
	if transactionLimit <= 0 {
		fmt.Println("Transaction limit set to 10 ")
		transactionLimit = 10
	}

	TransactionLimit = transactionLimit
}

// SetIDTransaction for requests
func SetIDTransaction(idTransaction string) {
	re := regexp.MustCompile("^[a-f0-9]{24}$")
	errIDTransaction := re.MatchString(idTransaction)
	if !errIDTransaction {
		log.Fatal("Invalid id_transaction")
	}

	IDTransaction = idTransaction
}

// SetDebug for requests
func SetDebug(debug bool) {
	Debug = debug
}
