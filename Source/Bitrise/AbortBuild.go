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
