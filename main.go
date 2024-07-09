package main

import (
	env "bitbucket-tools/Source/Environment"
	functionname "bitbucket-tools/Source/FunctionNames"
	function "bitbucket-tools/Source/Functions"
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Loading env variables...")
	println("-----------------------------------------")
	env.LoadEnvironment()
	env.DumpInputs()
	println("-----------------------------------------")
	fmt.Println("Loading env finished !!")

	fmt.Println()
	fmt.Println("Applying function....")
	println("-----------------------------------------")
	switch env.SELECTED_FUNCTION {
	case functionname.SKIP_VERIFICATION:
		function.PerformSkipVerification()
	case functionname.CHECK_PULL_REQUESTS_DEADLINE:
		function.PerformCheckPullRequestsDeadline()
	case functionname.CALL_API:
		function.PerformApiCall()
	case functionname.GENERATE_ICONS:
		function.PerformGenerateIcons()
	case functionname.TERMINATE_BUILD:
		function.PerformTerminateBuild()
	default:
		fmt.Println("Invalid function selected !!")
	}
	fmt.Println("-----------------------------------------")
	fmt.Println("Function completed !!")

	fmt.Println()
	fmt.Println("Outputs:")
	fmt.Println("-----------------------------------------")
	env.DumpOutputs()
	println("-----------------------------------------")

	fmt.Println("Tools signing off :)")
}
