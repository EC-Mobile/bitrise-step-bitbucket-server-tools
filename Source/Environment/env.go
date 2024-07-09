package env

import (
	functionname "bitbucket-tools/Source/FunctionNames"
	msc "bitbucket-tools/Source/Msc"
	"fmt"
	"image/color"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	// Bitbucket Auth
	BITBUCKET_API_ACCESS_TOKEN string
	BITBUCKET_DOMAIN           string
	BITBUCKET_PROJECT_KEY      string
	BITBUCKET_REPO_SLUG        string

	// Bitrise Auth
	BITRISE_API_ACCESS_TOKEN string
	BITRISE_APP_SLUG         string

	// Selected Function
	SELECTED_FUNCTION string

	// Skip Verification
	SV_PR_ID              string
	SV_BITRISE_BUILD_SLUG string

	// Check Pull Requests Deadline
	CPRD_PR_ID                   string
	CPRD_AUTHOR_EMAILS           string
	CPRD_REQUIRED_APPROVAL_COUNT int
	CPRD_TITLE                   string

	// Call Api
	CA_METHOD  string
	CA_URL     string
	CA_HEADERS string
	CA_BODY    string

	// Generate Icons
	GI_BASE_ICONS_SET          string
	GI_DESTINATION_ICONS_SET   string
	GI_ICON_PRIMARY_COLOR      color.Color
	GI_ICON_SECONDARY_COLOR    color.Color
	GI_ICON_LABEL_COLOR        color.Color
	GI_ICON_SCALER             int
	GI_ICON_OVERLAY_RESOLUTION int
	GI_ICON_OVERLAY_SCALE      int
	GI_ICON_OVERLAY_FONT_SIZE  int
	GI_APP_VERSION_NUMBER      string
	GI_APP_BUILD_NUMBER        string
	GI_APP_BUILD_TYPE          string

	// Terminate Build
	TB_BITRISE_BUILD_SLUG string
	TB_REGEX              string
	TB_VALUE              string
	TB_REASON             string
)

func LoadEnvironment() {
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("Env variable load error: %s", err)
		fmt.Println()
	}

	// Bitbucket Auth
	BITBUCKET_API_ACCESS_TOKEN = os.Getenv("bitbucket_api_access_token")
	BITBUCKET_DOMAIN = os.Getenv("bitbucket_domain")
	BITBUCKET_PROJECT_KEY = os.Getenv("bitbucket_project_key")
	BITBUCKET_REPO_SLUG = os.Getenv("bitbucket_repo_slug")

	// Bitbucket Auth
	BITRISE_API_ACCESS_TOKEN = os.Getenv("bitrise_api_access_token")
	BITRISE_APP_SLUG = os.Getenv("bitrise_app_slug")

	// Function to Perform
	SELECTED_FUNCTION = os.Getenv("selected_function")

	// Skip Verification
	SV_PR_ID = os.Getenv("sv_pr_id")
	SV_BITRISE_BUILD_SLUG = os.Getenv("sv_bitrise_build_slug")

	// Check Pull Requests Deadline
	CPRD_PR_ID = os.Getenv("cprd_pr_id")
	CPRD_AUTHOR_EMAILS = os.Getenv("cprd_author_emails")
	CPRD_REQUIRED_APPROVAL_COUNT, _ = strconv.Atoi(os.Getenv("cprd_required_approval_count"))
	CPRD_TITLE = os.Getenv("cprd_title")

	// Call Api
	CA_METHOD = os.Getenv("ca_method")
	CA_URL = os.Getenv("ca_url")
	CA_HEADERS = os.Getenv("ca_headers")
	CA_BODY = os.Getenv("ca_body")

	// Generate Icons
	GI_BASE_ICONS_SET = os.Getenv("gi_base_icons_set")
	GI_DESTINATION_ICONS_SET = os.Getenv("gi_destination_icons_set")
	GI_ICON_PRIMARY_COLOR, _ = msc.ParseHexColor(os.Getenv("gi_icon_primary_color"))
	GI_ICON_SECONDARY_COLOR, _ = msc.ParseHexColor(os.Getenv("gi_icon_secondary_color"))
	GI_ICON_LABEL_COLOR, _ = msc.ParseHexColor(os.Getenv("gi_icon_label_color"))
	GI_ICON_SCALER, _ = strconv.Atoi(os.Getenv("gi_icon_scaler"))
	GI_ICON_OVERLAY_RESOLUTION, _ = strconv.Atoi(os.Getenv("gi_icon_overlay_resolution"))
	GI_ICON_OVERLAY_SCALE, _ = strconv.Atoi(os.Getenv("gi_icon_overlay_scale"))
	GI_ICON_OVERLAY_FONT_SIZE, _ = strconv.Atoi(os.Getenv("gi_icon_overlay_font_size"))
	GI_APP_VERSION_NUMBER = os.Getenv("gi_app_version_number")
	GI_APP_BUILD_NUMBER = os.Getenv("gi_app_build_number")
	GI_APP_BUILD_TYPE = os.Getenv("gi_app_build_type")

	// Skip Verification
	TB_BITRISE_BUILD_SLUG = os.Getenv("tb_bitrise_build_slug")
	TB_REGEX = os.Getenv("tb_regex")
	TB_VALUE = os.Getenv("tb_value")
	TB_REASON = os.Getenv("tb_reason")
}

func DumpInputs() {
	values := map[string]string{}

	// Bitbucket Auth
	values["Bitbucket Token Provided:"] = strconv.FormatBool(len(BITBUCKET_API_ACCESS_TOKEN) != 0)
	values["Bitbucket Base Url:"] = BITBUCKET_DOMAIN
	values["Bitbucket Project Key:"] = BITBUCKET_PROJECT_KEY
	values["Bitbucket Repo Slug:"] = BITBUCKET_REPO_SLUG

	// Bitrise Auth
	values["Bitrise Token Provided:"] = strconv.FormatBool(len(BITRISE_API_ACCESS_TOKEN) != 0)
	values["Bitrise App Slug:"] = BITRISE_APP_SLUG

	print(values)
	fmt.Println()
	values = map[string]string{}

	// Selected Function
	values["Selected Function:"] = SELECTED_FUNCTION

	print(values)
	fmt.Println("........................................")
	values = map[string]string{}

	// Skip Verification
	values[functionname.SKIP_VERIFICATION+"For Pull Request Id:"] = SV_PR_ID
	values[functionname.SKIP_VERIFICATION+"Bitrise Build Slug:"] = SV_BITRISE_BUILD_SLUG

	// Check Pull Requests Deadline
	values[functionname.CHECK_PULL_REQUESTS_DEADLINE+"Filter by Pull Request ID:"] = CPRD_PR_ID
	values[functionname.CHECK_PULL_REQUESTS_DEADLINE+"Filter by Author Emails:"] = CPRD_AUTHOR_EMAILS
	values[functionname.CHECK_PULL_REQUESTS_DEADLINE+"Required Approvals Count:"] = strconv.Itoa(CPRD_REQUIRED_APPROVAL_COUNT)
	values[functionname.CHECK_PULL_REQUESTS_DEADLINE+"Filter by Title Regex:"] = CPRD_TITLE

	// Api Call
	values[functionname.CALL_API+"Method:"] = CA_METHOD
	values[functionname.CALL_API+"URL:"] = CA_URL
	values[functionname.CALL_API+"Headers:"] = CA_HEADERS
	values[functionname.CALL_API+"Body:"] = CA_BODY

	// Generate Icons
	values[functionname.GENERATE_ICONS+"Base Icons Set:"] = GI_BASE_ICONS_SET
	values[functionname.GENERATE_ICONS+"Destination Icons Set:"] = GI_DESTINATION_ICONS_SET
	// values[functionname.GENERATE_ICONS + "Icon Primary Color:"] = GI_ICON_PRIMARY_COLOR
	// values[functionname.GENERATE_ICONS + "Icon Secondary Color:"] = GI_ICON_SECONDARY_COLOR
	// values[functionname.GENERATE_ICONS + "Icon Label Color:"] = GI_ICON_LABEL_COLOR
	values[functionname.GENERATE_ICONS+"Icon Scalar:"] = strconv.Itoa(GI_ICON_SCALER)
	values[functionname.GENERATE_ICONS+"Icon Resolution:"] = strconv.Itoa(GI_ICON_OVERLAY_RESOLUTION)
	values[functionname.GENERATE_ICONS+"Icon Scale:"] = strconv.Itoa(GI_ICON_OVERLAY_SCALE)
	values[functionname.GENERATE_ICONS+"Icon Font Size:"] = strconv.Itoa(GI_ICON_OVERLAY_FONT_SIZE)
	values[functionname.GENERATE_ICONS+"App Version Number:"] = GI_APP_VERSION_NUMBER
	values[functionname.GENERATE_ICONS+"App Build Number:"] = GI_APP_BUILD_NUMBER
	values[functionname.GENERATE_ICONS+"App Build Type:"] = GI_APP_BUILD_TYPE

	// Terminate Build
	values[functionname.TERMINATE_BUILD+"Terminate Build Slug:"] = TB_BITRISE_BUILD_SLUG
	values[functionname.TERMINATE_BUILD+"Condition:"] = TB_REGEX
	values[functionname.TERMINATE_BUILD+"Value:"] = TB_VALUE
	values[functionname.TERMINATE_BUILD+"Reason of termination:"] = TB_REASON

	printKV(values, SELECTED_FUNCTION)
}

// Outputs

const (
	// Skip Verification
	SV_SKIPPED_VERIFICATION = "SV_SKIPPED_VERIFICATION"

	// Check Pull Requests Deadline
	CPRD_PULL_REQUESTS_DEADLINE_NEAR = "CPRD_PULL_REQUESTS_DEADLINE_NEAR"
	CPRD_PULL_REQUESTS_DEADLINE      = "CPRD_PULL_REQUESTS_DEADLINE"

	// Call Api
	CA_RESPONSE_BODY = "CA_RESPONSE_BODY"
)

func DumpOutputs() {
	values := map[string]string{}

	// Skip Verification
	values[functionname.SKIP_VERIFICATION+"Verification Skipped:"] = os.Getenv(SV_SKIPPED_VERIFICATION)

	// Check Pull Requests Deadline
	values[functionname.CHECK_PULL_REQUESTS_DEADLINE+"Pull Requests Deadline Near:"] = os.Getenv(CPRD_PULL_REQUESTS_DEADLINE_NEAR)
	values[functionname.CHECK_PULL_REQUESTS_DEADLINE+"Pull Requests Deadline Info:"] = os.Getenv(CPRD_PULL_REQUESTS_DEADLINE)

	// Call Api
	values[functionname.CALL_API+"Response Body:"] = os.Getenv(CA_RESPONSE_BODY)

	printKV(values, SELECTED_FUNCTION)
}

func print(values map[string]string) {
	for title, value := range values {
		if len(value) > 0 {
			fmt.Println(title + " " + value)
		}
	}
}

func printKV(values map[string]string, function string) {
	for title, value := range values {
		if len(function) > 0 {
			if strings.Contains(title, function) {
				newTitle := strings.Replace(title, function, "", 1)
				fmt.Println(newTitle + " " + value)
			}
		}
	}
}

func printK(values map[string]string, function string) {
	for title, value := range values {
		if len(value) > 0 && len(function) > 0 {
			if strings.Contains(title, function) {
				newTitle := strings.Replace(title, function, "", 1)
				fmt.Println(newTitle + " " + value)
			}
		}
	}
}

func Setenv(key string, value string) {
	c := exec.Command("envman", "add", "--key", key, "--value", value)
	err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Setenv(key, value)
}
