package function

import (
	env "bitbucket-tools/Source/Environment"
	network "bitbucket-tools/Source/Network"
	"fmt"
	"os"
	"strings"
)

func PerformApiCall() {
	fmt.Println()
	fmt.Println("Reseting Export Values....")
	fmt.Println("........................................")
	os.Setenv(env.RESPONSE_BODY, "")

	fmt.Println()
	fmt.Println("Preparing headers....")
	headers := prepareHeaders(env.HEADERS)

	fmt.Println("Calling Api....")
	fmt.Println("........................................")
	success, data := network.CallApi(env.METHOD, env.URL, headers, env.BODY)
	fmt.Printf("Api Success: %t\n", success)
	response := ""
	if success {
		response = string(data)
		fmt.Println("Response is: ")
		fmt.Println(response)
	}

	fmt.Println()
	fmt.Println("Exporting Results....")
	fmt.Println("........................................")
	os.Setenv(env.RESPONSE_BODY, response)
	fmt.Println("Exported !!!")
}

func prepareHeaders(keyValueString string) []network.KeyValue {
	headers := []network.KeyValue{}
	splits := strings.Split(keyValueString, "|")
	index := 1
	totalSplits := len(splits)

	for index < totalSplits {
		headers = append(headers, network.KeyValue{
			Key:   splits[index-1],
			Value: splits[index],
		})
		index += 2
	}
	return headers
}
