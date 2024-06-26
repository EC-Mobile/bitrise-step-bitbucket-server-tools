package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	BITBUCKET_ACCESS_TOKEN string
	BITBUCKET_DOMAIN       string
	BITBUCKET_PROJECT_KEY  string
	BITBUCKET_REPO_SLUG    string

	BITRISE_API_ACCESS_TOKEN string
	BITRISE_APP_SLUG         string
	BITRISE_BUILD_SLUG       string

	FUNCTION    string
	COMMIT_HASH string
	PR_TITLE    string
	PR_ID       string
)

func LoadEnvironment() {
	fmt.Println("Loading env variables...")
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("Env variable load error: %s", err)
		fmt.Println()
	}

	BITBUCKET_ACCESS_TOKEN = os.Getenv("BITBUCKET_ACCESS_TOKEN")
	BITBUCKET_DOMAIN = os.Getenv("BITBUCKET_DOMAIN")
	BITBUCKET_PROJECT_KEY = os.Getenv("BITBUCKET_PROJECT_KEY")
	BITBUCKET_REPO_SLUG = os.Getenv("BITBUCKET_REPO_SLUG")

	BITRISE_API_ACCESS_TOKEN = os.Getenv("BITRISE_API_ACCESS_TOKEN")
	BITRISE_APP_SLUG = os.Getenv("BITRISE_APP_SLUG")
	BITRISE_BUILD_SLUG = os.Getenv("BITRISE_BUILD_SLUG")

	FUNCTION = os.Getenv("FUNCTION")
	COMMIT_HASH = os.Getenv("COMMIT_HASH")
	PR_TITLE = os.Getenv("PR_TITLE")
	PR_ID = os.Getenv("PR_ID")

	DumpEnvironment()
	fmt.Println("Loading env finished !!")
}

func DumpEnvironment() {
	fmt.Println("-----------------------------------------")
	fmt.Println("Environment: ")
	fmt.Println("Bitbucket Base Url: " + BITBUCKET_DOMAIN)
	fmt.Println("Bitbucket Project Key: " + BITBUCKET_PROJECT_KEY)
	fmt.Println("Bitbucket Repo Slug: " + BITBUCKET_REPO_SLUG)
	fmt.Printf("Bitbucket Token Provided: %t", len(BITBUCKET_ACCESS_TOKEN) != 0)
	fmt.Println()
	fmt.Println()

	fmt.Println("Bitrise App Slug: " + BITRISE_APP_SLUG)
	fmt.Println("Bitrise Build Slug: " + BITRISE_BUILD_SLUG)
	fmt.Printf("Bitrise Token Provided: %t", len(BITRISE_API_ACCESS_TOKEN) != 0)
	fmt.Println()
	fmt.Println()

	fmt.Println("Selected Function: " + FUNCTION)
	fmt.Println("Commit Hash: " + COMMIT_HASH)
	fmt.Println("Pull Request Title: " + PR_TITLE)
	fmt.Println("Pull Request Id: " + PR_ID)
	fmt.Println("-----------------------------------------")
	fmt.Println("")
}
