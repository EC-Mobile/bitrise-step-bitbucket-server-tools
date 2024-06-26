package function

import (
	bitbucket "bitbucket-tools/Source/Bitbucket"
	env "bitbucket-tools/Source/Environment"
	"fmt"
	"strings"
)

func PerformSkipVerification() {
	if len(env.PR_ID) == 0 {
		fmt.Println("Pull Request id is not provided !")
		return
	}
	pullRequest := bitbucket.GetPullRequest(env.PR_ID)
	fmt.Println("Fetched pull request: " + pullRequest.Title)

	if len(pullRequest.Title) == 0 {
		fmt.Println("Pull Request title is invalid !")
		return
	}
	skipVerification := strings.Contains(strings.ToLower(pullRequest.Title), "[sv]")
	fmt.Printf("PR marked as skip verification: %t", skipVerification)
	fmt.Println()

	if skipVerification {
		fmt.Println("Fetched pull request Latest Commit: " + pullRequest.FromRef.LatestCommit)
	}
}
