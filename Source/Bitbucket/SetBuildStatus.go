package bitbucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BuildStatusValue struct {
	State       string `json:"state"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

// API Ref: https://docs.atlassian.com/bitbucket-server/rest/5.16.0/bitbucket-rest.html
// /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}
func SetBuildStatus(commitId string, status BuildStatusValue) bool {
	url := GetUrl() + "/rest/build-status/1.0/commits/" + commitId
	fmt.Println("Setting build Status: " + url)
	fmt.Println(status)

	reqBody, _ := json.Marshal(status)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", GetToken())
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
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
