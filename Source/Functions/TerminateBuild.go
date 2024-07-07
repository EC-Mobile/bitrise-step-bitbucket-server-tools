package function

import (
	bitrise "bitbucket-tools/Source/Bitrise"
	env "bitbucket-tools/Source/Environment"
	"fmt"
	"regexp"
)

func PerformTerminateBuild() {

	fmt.Println()
	fmt.Println("Checking Pre-Conditions....")
	fmt.Println("........................................")
	shouldTerminate := true
	if len(env.BODY) > 0 && len(env.TITLE) > 0 {
		regex, _ := regexp.Compile(env.TITLE)
		shouldTerminate = len(regex.FindAllString(env.BODY, 1)) > 0
	}
	fmt.Printf("Should terminate: %t\n", shouldTerminate)

	if shouldTerminate {
		fmt.Println()
		fmt.Println("Aborting Bitrise Build....")
		fmt.Println("........................................")
		buildAborted := bitrise.AbortBuild(env.BITRISE_BUILD_SLUG, env.DESCRIPTION, true, true)
		fmt.Printf("Build aborted: %t", buildAborted)
		fmt.Println()
	}
}
