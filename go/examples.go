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
	BaseURL      string `required:"true"`
	APIkey       string `required:"true"`
	Token        string
	Credentials  string
	IDUser       string
	UserName     string
	IDSite       string
	IDCredential string
	IDAccount    string
}

var (
	users        = flag.String("users", "", "{create|get|delete|modify}")
	catalogs     = flag.String("catalogs", "", "{sites}")
	sessions     = flag.String("sessions", "", "{get_token|verify_token}")
	credentials  = flag.String("credentials", "", "{create|get|delete}")
	accounts     = flag.String("accounts", "", "{get}")
	transactions = flag.String("transactions", "", "{get}")
	attachments  = flag.String("attachments", "", "{get}")
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
		// Create user
		examples.SetUserName(config.UserName)
		examples.CreateUser()
		break
	case "get":
		// Get ACME id_site
		examples.GetUsers()
		break
	case "modify":
		examples.SetIDUser(config.IDUser)
		examples.SetUserName(config.UserName)
		examples.ModifyUser()
		break
	case "delete":
		examples.SetIDUser(config.IDUser)
		examples.DeleteUser()
		break
	}

	switch *sessions {
	case "get_token":
		// Create session token
		examples.SetIDUser(config.IDUser)
		examples.CreateSession()
		break
	case "verify_token":
		// Verify session
		examples.SetToken(config.Token)
		examples.VerifySessions()
		break
	}

	switch *catalogs {
	case "sites":
		// Get ACME id_site
		examples.SetToken(config.Token)
		examples.GetSitesByOrganization()
		break
	}

	switch *credentials {
	case "create":
		examples.SetToken(config.Token)
		examples.SetCredentials(config.Credentials)
		examples.SetIDSite(config.IDSite)
		examples.CreateCredentials()
		break
	case "get":
		examples.SetToken(config.Token)
		examples.GetCredentials()
		break
	case "delete":
		examples.SetToken(config.Token)
		examples.SetIDCredential(config.IDCredential)
		examples.DeleteCredential()
		break
	}

	switch *accounts {
	case "get":
		examples.SetToken(config.Token)
		examples.SetIDCredential(config.IDCredential)
		examples.GetAccounts()
		break
	}

	switch *transactions {
	case "get":
		examples.SetToken(config.Token)
		examples.SetIDCredential(config.IDCredential)
		examples.SetIDAccount(config.IDAccount)
		examples.GetTransactions()
		break
	}

	switch *attachments {
	case "get":
		fmt.Println("Attachments")
		examples.SetToken(config.Token)
		examples.SetIDCredential(config.IDCredential)
		examples.SetIDAccount(config.IDAccount)
		examples.GetAttachments()
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
