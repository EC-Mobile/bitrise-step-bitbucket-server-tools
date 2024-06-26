package bitbucket

import (
	env "bitbucket-tools/Source/Environment"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Reviewer struct {
	Approved bool `json:"approved"`
}

type Ref struct {
	LatestCommit string `json:"latestCommit"`
}

type PullRequest struct {
	Title     string     `json:"title"`
	Open      bool       `json:"open"`
	Closed    bool       `json:"closed"`
	FromRef   Ref        `json:"fromRef"`
	Reviewers []Reviewer `json:"reviewers"`
}

func GetPullRequest(id string) PullRequest {
	url := GetUrl() + "/rest/api/1.0/projects/" + env.BITBUCKET_PROJECT_KEY + "/repos/" + env.BITBUCKET_REPO_SLUG + "/pull-requests/" + id
	fmt.Println("Getting PR from: " + url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", GetToken())
	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
		return PullRequest{}
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return PullRequest{}
	}

	var pullRequest PullRequest
	json.Unmarshal(responseData, &pullRequest)
	return pullRequest
}
