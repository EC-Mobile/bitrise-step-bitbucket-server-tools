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

	values := map[string]string{}
	values["Bitbucket Token Provided:"] = strconv.FormatBool(len(BITBUCKET_API_ACCESS_TOKEN) != 0)
	values["Bitbucket Base Url:"] = BITBUCKET_DOMAIN
	values["ucket Project Key:"] = BITBUCKET_PROJECT_KEY
	values["Bitbucket Repo Slug:"] = BITBUCKET_REPO_SLUG

	values["Bitrise Token Provided:"] = strconv.FormatBool(len(BITRISE_API_ACCESS_TOKEN) != 0)
	values["Bitrise App Slug:"] = BITRISE_APP_SLUG

	values["Selected Function:"] = FUNCTION
	values["Pull Request Id:"] = PR_ID
	values["Bitrise Build Slug:"] = BITRISE_BUILD_SLUG
	values["Emails:"] = EMAILS
	values["Required Approvals Count:"] = strconv.Itoa(REQUIRED_APPROVAL_COUNT)
	values["Title:"] = TITLE
	values["Description:"] = DESCRIPTION
	values["Method:"] = METHOD
	values["URL:"] = URL
	values["Headers:"] = HEADERS
	values["Body:"] = BODY

	values["Base Icons Set:"] = BASE_ICONS_SET
	values["Destination Icons Set:"] = DESTINATION_ICONS_SET
	// values["Icon Primary Color:"] = ICON_PRIMARY_COLOR
	// values["Icon Secondary Color:"] = ICON_SECONDARY_COLOR
	// values["Icon Label Color:"] = ICON_LABEL_COLOR
	values["Icon Scalar:"] = strconv.Itoa(ICON_SCALER)
	values["Icon Resolution:"] = strconv.Itoa(ICON_OVERLAY_RESOLUTION)
	values["Icon Scale:"] = strconv.Itoa(ICON_OVERLAY_SCALE)
	values["Icon Font Size:"] = strconv.Itoa(ICON_OVERLAY_FONT_SIZE)

	values["App Version Number:"] = APP_VERSION_NUMBER
	values["App Build Number:"] = APP_BUILD_NUMBER
	values["App Build Type:"] = APP_BUILD_TYPE

	fmt.Println("-----------------------------------------")
	fmt.Println("Inputs: ")
	for title, value := range values {
		if len(value) > 0 {
			fmt.Println(title + " " + value)
		}
	}
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
	values := map[string]string{}
	values["Verification Skipped:"] = os.Getenv(SKIPPED_VERIFICATION)
	values["Pull Requests Deadline Near:"] = os.Getenv(PULL_REQUESTS_DEADLINE_NEAR)
	values["Pull Requests Deadline Info:"] = os.Getenv(PULL_REQUESTS_DEADLINE)
	values["Response Body:"] = os.Getenv(RESPONSE_BODY)

	fmt.Println("-----------------------------------------")
	fmt.Println("Outputs: ")
	for title, value := range values {
		if len(value) > 0 {
			fmt.Println(title + " " + value)
		}
	}
	fmt.Println("-----------------------------------------")
	fmt.Println("")
}

func Setenv(key string, value string) {
	c := exec.Command("envman", "add", "--key", key, "--value", value)
	err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Setenv(key, value)
}
