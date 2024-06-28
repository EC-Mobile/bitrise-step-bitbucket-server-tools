package main

import (
	env "bitbucket-tools/Source/Environment"
	function "bitbucket-tools/Source/Functions"
	"fmt"
)

func main() {
	env.LoadEnvironment()

	fmt.Println("Applying Function....")
	println("-----------------------------------------")
	switch env.FUNCTION {
	case function.SKIP_VERIFICATION:
		function.PerformSkipVerification()
	default:
		fmt.Println("Invalid function selected !!")
	}
	env.DumpOutputs()
	println("-----------------------------------------")

	fmt.Println()
	fmt.Println("Function is completed !")
	fmt.Println("Tools signing off :)")
}
