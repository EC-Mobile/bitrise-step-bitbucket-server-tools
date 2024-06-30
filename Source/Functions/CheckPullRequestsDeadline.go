package function

import (
	bitbucket "bitbucket-tools/Source/Bitbucket"
	env "bitbucket-tools/Source/Environment"
	msc "bitbucket-tools/Source/Msc"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func PerformCheckPullRequestsDeadline() {
	fmt.Println()
	fmt.Println("Reseting Export Values....")
	fmt.Println("........................................")
	os.Setenv(env.PULL_REQUESTS_DEADLINE_NEAR, strconv.FormatBool(false))
	os.Setenv(env.PULL_REQUESTS_DEADLINE, "")

	fmt.Println()
	fmt.Println("Fetching all Open PRs info....")
	fmt.Println("........................................")
	pullRequests := bitbucket.GetPullRequests(bitbucket.PR_STATE_OPEN)
	fmt.Println("Fetched open pull requests: ")
	printPullRequests(pullRequests)

	if len(pullRequests) > 0 {
		fmt.Println()
		fmt.Println("Filtering by, not approved, title and desired author...")
		fmt.Println("........................................")
		pullRequests = msc.Filter(pullRequests, func(pullRequest bitbucket.PullRequest) bool {
			return !isApproved(pullRequest) && isFromDesiredAuthor(pullRequest) && matchesTitle(pullRequest) && matchesId(pullRequest)
		})
		printPullRequests(pullRequests)

		fmt.Println()
		fmt.Println("Checking deadline...")
		fmt.Println("........................................")
		isDeadlinesNear := isDeadlinesNear(pullRequests)
		fmt.Printf("Is deadlines near: %t\n", isDeadlinesNear)

		fmt.Println()
		fmt.Println("Preparing output...")
		fmt.Println("........................................")
		output := prepareOutput(pullRequests)
		outputJson := msc.ConvertToJson(output, true)
		fmt.Println("Output prepared !!")

		fmt.Println()
		fmt.Println("Exporting Results....")
		fmt.Println("........................................")
		os.Setenv(env.PULL_REQUESTS_DEADLINE_NEAR, strconv.FormatBool(isDeadlinesNear))
		os.Setenv(env.PULL_REQUESTS_DEADLINE, outputJson)
		fmt.Println("Exported !!!")
	}
}

func printPullRequests(pullRequests []bitbucket.PullRequest) {
	for index, pullRequest := range pullRequests {
		fmt.Printf("Pull Request: %d\nAuthor: %s\nApproved: %t\nTitle: %s\n", index, pullRequest.Author.User.EmailAddress, isApproved(pullRequest), pullRequest.Title)
		fmt.Println()
	}
}

func isApproved(pullRequest bitbucket.PullRequest) bool {
	approvalCount := 0
	for _, reviewer := range pullRequest.Reviewers {
		if reviewer.Approved {
			approvalCount++
		}
	}
	return approvalCount >= env.REQUIRED_APPROVAL_COUNT
}

func isFromDesiredAuthor(pullRequest bitbucket.PullRequest) bool {
	return strings.Contains(env.EMAILS, pullRequest.Author.User.EmailAddress)
}

func matchesId(pullRequest bitbucket.PullRequest) bool {
	id := strconv.Itoa(pullRequest.Id)
	return msc.Matches(env.PR_ID, id)
}

func matchesTitle(pullRequest bitbucket.PullRequest) bool {
	title := strings.ToLower(pullRequest.Title)
	return msc.Matches(env.TITLE, title)
}

func isDeadlinesNear(pullRequests []bitbucket.PullRequest) bool {
	for index, pullRequest := range pullRequests {
		fmt.Printf("%d: %s\n", index, pullRequest.Title)
		if isDeadlineNear(pullRequest) {
			return true
		}
	}
	return false
}

func isDeadlineNear(pullRequest bitbucket.PullRequest) bool {
	deadlineRegex, _ := regexp.Compile(`[deadline [0-9]+\/[0-9]+]`)
	dayRegex, _ := regexp.Compile(`[0-9]+\/[0-9]+`)
	title := strings.ToLower(pullRequest.Title)
	if deadlineRegex.MatchString(title) {
		dayValues := dayRegex.FindAllString(title, -1)
		if len(dayValues) > 0 {
			deadline := fmt.Sprintf("%s/%d", dayValues[0], time.Now().Year())
			deadlineDate, _ := time.Parse("2/1/2006", deadline)
			since := time.Since(deadlineDate)
			fmt.Printf("---> %s - %s(%f)\n", deadlineDate, since, -since.Hours())
			if -since.Hours() < 24 {
				return true
			}
		}
	}
	return false
}

type pullRequestsInfo struct {
	Title          string `json:"title"`
	Url            string `json:"url"`
	IsDeadlineNear bool   `json:"isDeadlineNear"`
}

type pullRequestsDeadlineOutup struct {
	PullRequests []pullRequestsInfo `json:"pullRequests"`
}

func prepareOutput(pullRequests []bitbucket.PullRequest) pullRequestsDeadlineOutup {
	output := pullRequestsDeadlineOutup{}
	output.PullRequests = []pullRequestsInfo{}

	for _, pullRequest := range pullRequests {
		info := pullRequestsInfo{}
		info.Title = pullRequest.Title
		info.IsDeadlineNear = isDeadlineNear(pullRequest)
		info.Url = ""
		if len(pullRequest.Links.Self) > 0 {
			info.Url = pullRequest.Links.Self[0].HRef
		}
		output.PullRequests = append(output.PullRequests, info)
	}

	return output
}
