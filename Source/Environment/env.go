package env

import (
	msc "bitbucket-tools/Source/Msc"
	"fmt"
	"image/color"
	"os"
	"os/exec"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	BITBUCKET_API_ACCESS_TOKEN string
	BITBUCKET_DOMAIN           string
	BITBUCKET_PROJECT_KEY      string
	BITBUCKET_REPO_SLUG        string

	BITRISE_API_ACCESS_TOKEN string
	BITRISE_APP_SLUG         string

	FUNCTION                string
	PR_ID                   string
	BITRISE_BUILD_SLUG      string
	EMAILS                  string
	REQUIRED_APPROVAL_COUNT int
	TITLE                   string
	DESCRIPTION             string
	METHOD                  string
	URL                     string
	HEADERS                 string
	BODY                    string

	BASE_ICONS_SET          string
	DESTINATION_ICONS_SET   string
	ICON_PRIMARY_COLOR      color.Color
	ICON_SECONDARY_COLOR    color.Color
	ICON_LABEL_COLOR        color.Color
	ICON_SCALER             int
	ICON_OVERLAY_RESOLUTION int
	ICON_OVERLAY_SCALE      int
	ICON_OVERLAY_FONT_SIZE  int
	APP_VERSION_NUMBER      string
	APP_BUILD_NUMBER        string
	APP_BUILD_TYPE          string
)

func LoadEnvironment() {
	fmt.Println("Loading env variables...")
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("Env variable load error: %s", err)
		fmt.Println()
	}

	BITBUCKET_API_ACCESS_TOKEN = os.Getenv("bitbucket_api_access_token")
	BITBUCKET_DOMAIN = os.Getenv("bitbucket_domain")
	BITBUCKET_PROJECT_KEY = os.Getenv("bitbucket_project_key")
	BITBUCKET_REPO_SLUG = os.Getenv("bitbucket_repo_slug")

	BITRISE_API_ACCESS_TOKEN = os.Getenv("bitrise_api_access_token")
	BITRISE_APP_SLUG = os.Getenv("bitrise_app_slug")

	FUNCTION = os.Getenv("function")
	PR_ID = os.Getenv("pr_id")
	BITRISE_BUILD_SLUG = os.Getenv("bitrise_build_slug")
	EMAILS = os.Getenv("emails")
	REQUIRED_APPROVAL_COUNT, _ = strconv.Atoi(os.Getenv("required_approval_count"))
	TITLE = os.Getenv("title")
	DESCRIPTION = os.Getenv("description")
	METHOD = os.Getenv("method")
	URL = os.Getenv("url")
	HEADERS = os.Getenv("headers")
	BODY = os.Getenv("body")

	BASE_ICONS_SET = os.Getenv("base_icons_set")
	DESTINATION_ICONS_SET = os.Getenv("destination_icons_set")
	ICON_PRIMARY_COLOR, _ = msc.ParseHexColor(os.Getenv("icon_primary_color"))
	ICON_SECONDARY_COLOR, _ = msc.ParseHexColor(os.Getenv("icon_secondary_color"))
	ICON_LABEL_COLOR, _ = msc.ParseHexColor(os.Getenv("icon_label_color"))
	ICON_SCALER, _ = strconv.Atoi(os.Getenv("icon_scaler"))
	ICON_OVERLAY_RESOLUTION, _ = strconv.Atoi(os.Getenv("icon_overlay_resolution"))
	ICON_OVERLAY_SCALE, _ = strconv.Atoi(os.Getenv("icon_overlay_scale"))
	ICON_OVERLAY_FONT_SIZE, _ = strconv.Atoi(os.Getenv("icon_overlay_font_size"))
	APP_VERSION_NUMBER = os.Getenv("app_version_number")
	APP_BUILD_NUMBER = os.Getenv("app_build_number")
	APP_BUILD_TYPE = os.Getenv("app_build_type")

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
	fmt.Println("Emails: " + EMAILS)
	fmt.Printf("Required Approvals Count: %d", REQUIRED_APPROVAL_COUNT)
	fmt.Println("Title: " + TITLE)
	fmt.Println("Method: " + METHOD)
	fmt.Println("Url: " + URL)
	fmt.Println("Headers: " + HEADERS)
	fmt.Println("Body: " + BODY)

	fmt.Println()
	fmt.Println("Base Icons Set: " + BASE_ICONS_SET)
	fmt.Println("Destination Icons Set: " + DESTINATION_ICONS_SET)
	fmt.Printf("Icon Primary Color: %s\n", ICON_PRIMARY_COLOR)
	fmt.Printf("Icon Secondary Color: %s\n", ICON_SECONDARY_COLOR)
	fmt.Printf("Icon Label Color: %s\n", ICON_LABEL_COLOR)

	fmt.Printf("Icon Scalar: %d\n", ICON_SCALER)
	fmt.Printf("Icon Resolution: %d\n", ICON_OVERLAY_RESOLUTION)
	fmt.Printf("Icon Scale: %d\n", ICON_OVERLAY_SCALE)
	fmt.Printf("Icon Font Size: %d\n", ICON_OVERLAY_FONT_SIZE)

	fmt.Println("App Version Number: " + APP_VERSION_NUMBER)
	fmt.Println("App Build Number: " + APP_BUILD_NUMBER)
	fmt.Println("App Build Type: " + APP_BUILD_TYPE)

	fmt.Println()
	fmt.Println("-----------------------------------------")
	fmt.Println("")
}

const (
	SKIPPED_VERIFICATION = "SKIPPED_VERIFICATION"

	PULL_REQUESTS_DEADLINE_NEAR = "PULL_REQUESTS_DEADLINE_NEAR"
	PULL_REQUESTS_DEADLINE      = "PULL_REQUESTS_DEADLINE"
	RESPONSE_BODY               = "RESPONSE_BODY"
)

func DumpOutputs() {
	fmt.Println()
	fmt.Println("........................................")
	fmt.Println("Outputs: ")
	fmt.Println("Verification Skipped: " + os.Getenv(SKIPPED_VERIFICATION))

	fmt.Printf("Pull Requests Deadline Near: %s\n", os.Getenv(PULL_REQUESTS_DEADLINE_NEAR))
	fmt.Println("Pull Requests Deadline Info: \n" + os.Getenv(PULL_REQUESTS_DEADLINE))
	fmt.Println("Response Body: \n" + os.Getenv(RESPONSE_BODY))
}

func Setenv(key string, value string) {
	c := exec.Command("envman", "add", "--key", key, "--value", value)
	err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Setenv(key, value)
}
