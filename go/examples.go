package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"paybook.com/sync-code-samples/examples"
)

const (
	get = "get"
)

// Config app
type Config struct {
	APIkey           string `required:"true"`
	BaseURL          string `required:"true"`
	Credentials      string
	UserName         string
	IDAccount        string
	IDCredential     string
	IDJob            string
	IDSite           string
	IDTransaction    string
	IDUser           string
	Token            string
	TransactionLimit int
}

var (
	accounts     = flag.String("accounts", "", "{get}")
	attachments  = flag.String("attachments", "", "{get|download}")
	catalogs     = flag.String("catalogs", "", "{sites}")
	credentials  = flag.String("credentials", "", "{create|get|delete}")
	documents    = flag.String("documents", "", "{get}")
	sessions     = flag.String("sessions", "", "{get_token|verify_token}")
	transactions = flag.String("transactions", "", "{get}")
	test         = flag.Bool("test", false, "true")
	users        = flag.String("users", "", "{create|get|delete|modify}")
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

	switch *accounts {
	case get:
		fmt.Println("Get accounts")
		examples.SetToken(config.Token)
		examples.SetIDCredential(config.IDCredential)
		examples.GetAccounts()
		break
	}

	switch *attachments {
	case get:
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
	case get:
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

	switch *documents {
	case get:
		fmt.Println("Get documents")
		examples.SetToken(config.Token)
		examples.GetDocuments()
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

	switch *transactions {
	case get:
		fmt.Println("Get transactions")
		examples.SetToken(config.Token)
		examples.SetIDCredential(config.IDCredential)
		examples.SetIDAccount(config.IDAccount)
		examples.SetTransactionLimit(config.TransactionLimit)
		examples.GetTransactions()
		break
	}

	switch *users {
	case "create":
		fmt.Println("Create user")
		// Create user
		examples.SetUserName(config.UserName)
		examples.CreateUser()
		break
	case get:
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

	if *test {
		Test()
	}
}

// Test samples
func Test() {
	fmt.Println("Test Go samples:")
	fmt.Println("SYNC_BASEURL=", config.BaseURL)
}
