package bitbucket

import (
	env "bitbucket-tools/Source/Environment"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Reviewer struct {
	Approved bool `json:"approved"`
}

type User struct {
	EmailAddress string `json:"emailAddress"`
}

type Ref struct {
	LatestCommit string `json:"latestCommit"`
}

type PullRequest struct {
	Id        int        `json:id`
	Title     string     `json:"title"`
	Open      bool       `json:"open"`
	Closed    bool       `json:"closed"`
	FromRef   Ref        `json:"fromRef"`
	Reviewers []Reviewer `json:"reviewers"`
	Links     struct {
		Self []struct {
			HRef string `json:"href"`
		} `json:"self"`
	} `json:"links"`
	Author struct {
		User User `json:"user"`
	} `json:"author"`
}

func GetPullRequest(id string) PullRequest {
	url := GetUrl() + "/rest/api/1.0/projects/" + env.BITBUCKET_PROJECT_KEY + "/repos/" + env.BITBUCKET_REPO_SLUG + "/pull-requests/" + id
	fmt.Println("Getting PR from: " + url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", GetToken())
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return PullRequest{}
	}

	responseData, readError := io.ReadAll(response.Body)
	if readError != nil {
		fmt.Println(readError.Error())
		return PullRequest{}
	}

	if response.StatusCode < 200 || response.StatusCode > 210 {
		fmt.Printf("Invalid status code: %d", response.StatusCode)
		fmt.Println()
		return PullRequest{}
	}

	var pullRequest PullRequest
	json.Unmarshal(responseData, &pullRequest)
	return pullRequest
}
