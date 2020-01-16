package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"paybook.com/sync-code-samples/examples"
)

// Config app
type Config struct {
	BaseURL          string `required:"true"`
	APIkey           string `required:"true"`
	Token            string
	Credentials      string
	IDUser           string
	UserName         string
	IDSite           string
	IDCredential     string
	IDAccount        string
	IDTransaction    string
	TransactionLimit int
}

var (
	users        = flag.String("users", "", "{create|get|delete|modify}")
	catalogs     = flag.String("catalogs", "", "{sites}")
	sessions     = flag.String("sessions", "", "{get_token|verify_token}")
	credentials  = flag.String("credentials", "", "{create|get|delete}")
	accounts     = flag.String("accounts", "", "{get}")
	transactions = flag.String("transactions", "", "{get}")
	attachments  = flag.String("attachments", "", "{get|download}")
	test         = flag.Bool("test", false, "true")
	config       Config
)

func main() {
	flag.Parse()

	err := envconfig.Process("sync", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Setup and validate environment
	examples.Setup(config.BaseURL, config.APIkey)
	examples.SetDebug(false)

	switch *users {
	case "create":
		fmt.Println("Create user")
		// Create user
		examples.SetUserName(config.UserName)
		examples.CreateUser()
		break
	case "get":
		fmt.Println("Get users")
		// Get ACME id_site
		examples.GetUsers()
		break
	case "modify":
		fmt.Println("Modify user")
		examples.SetIDUser(config.IDUser)
		examples.SetUserName(config.UserName)
		examples.ModifyUser()
		break
	case "delete":
		fmt.Println("Delete user")
		examples.SetIDUser(config.IDUser)
		examples.DeleteUser()
		break
	}

	switch *sessions {
	case "get_token":
		// Create session token
		fmt.Println("Create token")
		examples.SetIDUser(config.IDUser)
		examples.CreateSession()
		break
	case "verify_token":
		// Verify session
		fmt.Println("Verify token")
		examples.SetToken(config.Token)
		examples.VerifySessions()
		break
	}

	switch *catalogs {
	case "sites":
		fmt.Println("Get sites")
		examples.SetToken(config.Token)
		examples.GetSitesByOrganization()
		break
	}

	switch *credentials {
	case "create":
		fmt.Println("Create credential")
		examples.SetToken(config.Token)
		examples.SetCredentials(config.Credentials)
		examples.SetIDSite(config.IDSite)
		examples.CreateCredentials()
		break
	case "get":
		fmt.Println("Get credentials")
		examples.SetToken(config.Token)
		examples.GetCredentials()
		break
	case "delete":
		fmt.Println("Delete credential")
		examples.SetToken(config.Token)
		examples.SetIDCredential(config.IDCredential)
		examples.DeleteCredential()
		break
	}

	switch *accounts {
	case "get":
		fmt.Println("Get accounts")
		examples.SetToken(config.Token)
		examples.SetIDCredential(config.IDCredential)
		examples.GetAccounts()
		break
	}

	switch *transactions {
	case "get":
		fmt.Println("Get transactions")
		examples.SetToken(config.Token)
		examples.SetIDCredential(config.IDCredential)
		examples.SetIDAccount(config.IDAccount)
		examples.SetTransactionLimit(config.TransactionLimit)
		examples.GetTransactions()
		break
	}

	switch *attachments {
	case "get":
		fmt.Println("Get attachments")
		examples.SetToken(config.Token)
		examples.SetIDCredential(config.IDCredential)
		examples.SetIDAccount(config.IDAccount)
		examples.SetIDTransaction(config.IDTransaction)
		examples.GetAttachments()
		break
	case "download":
		fmt.Println("Downloading attachments")
		examples.SetToken(config.Token)
		examples.SetIDCredential(config.IDCredential)
		examples.SetIDAccount(config.IDAccount)
		examples.SetIDTransaction(config.IDTransaction)
		examples.DownloadAttachments()
		break
	}

	if *test {
		Test()
	}
}

// Test samples
func Test() {
	fmt.Println("Test Go samples:")
	fmt.Println("SYNC_BASEURL=", config.BaseURL)
}
