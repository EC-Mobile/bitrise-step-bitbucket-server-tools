package bitrise

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// API Ref; https://devcenter.bitrise.io/en/api/api-reference.html
// /apps/{app-slug}/builds/{build-slug}/abort
func AbortBuild(buildSlug string, reason string, success bool, skipNotifications bool) bool {
	url := GetBuildsUrl() + buildSlug + "/abort"
	data := fmt.Sprintf(`{ "abort_reason": "%s", "abort_with_success": %t, "skip_notifications": %t }`, reason, success, skipNotifications)
	fmt.Println("Aborting Build: " + url)
	fmt.Println("Data: " + data)

	reqBody := []byte(data)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", GetToken())
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	_, readError := io.ReadAll(response.Body)
	if readError != nil {
		fmt.Println(readError.Error())
		return false
	}

	if response.StatusCode >= 200 && response.StatusCode < 210 {
		return true
	}

	fmt.Printf("Invalid status code: %d", response.StatusCode)
	fmt.Println()
	return false
}

// #!/bin/bash
// set -e
// set -o pipefail
// # set -x

// echo "-------------------------------------------"
// echo "App Slug: $PR_NOTIFICATION_TEAMS_GROUP"
// echo "Build Slug: $BITRISE_BUILD_SLUG"
// echo "Git Message: $BITRISE_GIT_MESSAGE"
// echo "-------------------------------------------"
// echo "\n"

// echo "Checking SV tag in Git Message...."
// svRegex="\[sv\]"
// skipVerification=`echo $BITRISE_GIT_MESSAGE | tr "[:upper:]" "[:lower:]" | { grep "$svRegex" -o || true; }`

// if [ -z $skipVerification ]
// then
//   envman add --key SKIP_VERIFICATION --value false
//   echo "Skip Verification tag not found."
//   echo "Lets verify...."
// else
//   echo "-------------------------------------------"

// https://api.bitrise.io/v0.1/apps/$BITRISE_APP_SLUG/builds/$BITRISE_BUILD_SLUG/abort
//   API=""
//   echo "Abort API Link: $API"
//   echo "-------------------------------------------"
//   echo "\n"

//   curl -X POST "$API" \
//     -H 'accept: application/json' \
//     -H "Authorization: $BITRISE_API_ACCESS_TOKEN" \
//     -H 'Content-Type: application/json' \
//     -d "{
//       \"abort_reason\": \"Marked as [SV], so skipping verification.\",
//       \"abort_with_success\": true,
//       \"skip_notifications\": false
//   }"
//   envman add --key SKIP_VERIFICATION --value true
// fi
