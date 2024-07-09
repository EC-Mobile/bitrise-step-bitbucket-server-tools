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
	if len(env.TB_VALUE) > 0 && len(env.TB_REGEX) > 0 {
		regex, _ := regexp.Compile(env.TB_REGEX)
		shouldTerminate = len(regex.FindAllString(env.TB_VALUE, 1)) > 0
	}
	fmt.Printf("Should terminate: %t\n", shouldTerminate)

	if shouldTerminate {
		fmt.Println()
		fmt.Println("Aborting Bitrise Build....")
		fmt.Println("........................................")
		buildAborted := bitrise.AbortBuild(env.TB_BITRISE_BUILD_SLUG, env.TB_REASON, true, true)
		fmt.Printf("Build aborted: %t", buildAborted)
		fmt.Println()
	}
}
