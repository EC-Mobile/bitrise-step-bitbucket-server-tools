package bitbucket

import (
	env "bitbucket-tools/Source/Environment"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	PR_STATE_OPEN = "OPEN"
)

type GetPullRequestsResponse struct {
	Values []PullRequest `json:"values"`
}

func GetPullRequests(state string) []PullRequest {
	url := GetUrl() + "/rest/api/1.0/projects/" + env.BITBUCKET_PROJECT_KEY + "/repos/" + env.BITBUCKET_REPO_SLUG + "/pull-requests?state=" + state
	fmt.Println("Getting PRs from: " + url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", GetToken())
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return []PullRequest{}
	}

	responseData, readError := io.ReadAll(response.Body)
	if readError != nil {
		fmt.Println(readError.Error())
		return []PullRequest{}
	}

	if response.StatusCode < 200 || response.StatusCode > 210 {
		fmt.Printf("Invalid status code: %d", response.StatusCode)
		fmt.Println()
		return []PullRequest{}
	}

	var reponse GetPullRequestsResponse
	json.Unmarshal(responseData, &reponse)
	return reponse.Values
}
