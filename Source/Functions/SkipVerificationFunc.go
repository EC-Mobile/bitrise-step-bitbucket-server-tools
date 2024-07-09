package function

import (
	bitbucket "bitbucket-tools/Source/Bitbucket"
	bitrise "bitbucket-tools/Source/Bitrise"
	env "bitbucket-tools/Source/Environment"
	"fmt"
	"strconv"
	"strings"
)

func PerformSkipVerification() {
	fmt.Println()
	fmt.Println("Reseting Export Values....")
	fmt.Println("........................................")
	env.Setenv(env.SV_SKIPPED_VERIFICATION, strconv.FormatBool(false))

	fmt.Println()
	fmt.Println("Fetching the PR info....")
	fmt.Println("........................................")
	if len(env.SV_PR_ID) == 0 {
		fmt.Println("Pull Request id is not provided !")
		return
	}
	pullRequest := bitbucket.GetPullRequest(env.SV_PR_ID)
	fmt.Println("Fetched pull request: " + pullRequest.Title)

	if len(pullRequest.Title) == 0 {
		fmt.Println("Pull Request title is invalid !")
		return
	}
	skipVerification := strings.Contains(strings.ToLower(pullRequest.Title), "[sv]")
	fmt.Printf("PR marked as skip verification: %t", skipVerification)
	fmt.Println()

	if skipVerification {
		fmt.Println()
		fmt.Println("Adding Build Status....")
		fmt.Println("........................................")
		fmt.Println("Fetched pull request Latest Commit: " + pullRequest.FromRef.LatestCommit)
		status := bitbucket.BuildStatusValue{
			State:       "FAILED",
			Key:         "[SV] Build Status for PR " + env.SV_PR_ID,
			Name:        "[SV] Build Status for PR " + env.SV_PR_ID,
			Url:         bitrise.GetBuildUrl(env.SV_BITRISE_BUILD_SLUG),
			Description: "Build maked FAILED due to [SV] tag.",
		}
		statusAdded := bitbucket.SetBuildStatus(pullRequest.FromRef.LatestCommit, status)
		fmt.Printf("Pull Request build status added: %t", statusAdded)
		fmt.Println()

		fmt.Println()
		fmt.Println("Aborting Bitrise Build....")
		fmt.Println("........................................")
		buildAborted := bitrise.AbortBuild(env.SV_BITRISE_BUILD_SLUG, "PR is marked as [SV], so skipping verification", true, true)
		fmt.Printf("Build aborted added: %t", buildAborted)
		fmt.Println()

		fmt.Println()
		fmt.Println("Exporting Results....")
		fmt.Println("........................................")
		env.Setenv(env.SV_SKIPPED_VERIFICATION, strconv.FormatBool(buildAborted))
		fmt.Println("Exported !!!")
	}
}
