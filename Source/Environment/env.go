package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	BITBUCKET_API_ACCESS_TOKEN string
	BITBUCKET_DOMAIN           string
	BITBUCKET_PROJECT_KEY      string
	BITBUCKET_REPO_SLUG        string

	BITRISE_API_ACCESS_TOKEN string
	BITRISE_APP_SLUG         string

	FUNCTION           string
	PR_ID              string
	BITRISE_BUILD_SLUG string
)

func LoadEnvironment() {
	fmt.Println("Loading env variables...")
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("Env variable load error: %s", err)
		fmt.Println()
	}

	BITBUCKET_API_ACCESS_TOKEN = os.Getenv("BITBUCKET_API_ACCESS_TOKEN")
	BITBUCKET_DOMAIN = os.Getenv("BITBUCKET_DOMAIN")
	BITBUCKET_PROJECT_KEY = os.Getenv("BITBUCKET_PROJECT_KEY")
	BITBUCKET_REPO_SLUG = os.Getenv("BITBUCKET_REPO_SLUG")

	BITRISE_API_ACCESS_TOKEN = os.Getenv("BITRISE_API_ACCESS_TOKEN")
	BITRISE_APP_SLUG = os.Getenv("BITRISE_APP_SLUG")

	FUNCTION = os.Getenv("FUNCTION")
	PR_ID = os.Getenv("PR_ID")
	BITRISE_BUILD_SLUG = os.Getenv("BITRISE_BUILD_SLUG")

	DumpInputs()
	fmt.Println("Loading env finished !!")
}

func DumpInputs() {
	fmt.Println("-----------------------------------------")
	fmt.Println("Inputs: ")
	fmt.Println("Bitbucket Base Url: " + BITBUCKET_DOMAIN)
	fmt.Println("Bitbucket Project Key: " + BITBUCKET_PROJECT_KEY)
	fmt.Println("Bitbucket Repo Slug: " + BITBUCKET_REPO_SLUG)
	fmt.Printf("Bitbucket Token Provided: %t", len(BITBUCKET_API_ACCESS_TOKEN) != 0)
	fmt.Println()
	fmt.Println()

	fmt.Println("Bitrise App Slug: " + BITRISE_APP_SLUG)
	fmt.Printf("Bitrise Token Provided: %t", len(BITRISE_API_ACCESS_TOKEN) != 0)
	fmt.Println()
	fmt.Println()

	fmt.Println("Selected Function: " + FUNCTION)
	fmt.Println("Pull Request Id: " + PR_ID)
	fmt.Println("Bitrise Build Slug: " + BITRISE_BUILD_SLUG)
	fmt.Println("-----------------------------------------")
	fmt.Println("")
}

const (
	SKIPPED_VERIFICATION = "SKIPPED_VERIFICATION"
)

func DumpOutputs() {
	fmt.Println()
	fmt.Println("........................................")
	fmt.Println("Outputs: ")
	fmt.Println("Verification Skipped: " + SKIPPED_VERIFICATION)
}
