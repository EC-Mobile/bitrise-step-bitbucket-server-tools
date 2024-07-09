package function

import (
	env "bitbucket-tools/Source/Environment"
	network "bitbucket-tools/Source/Network"
	"fmt"
	"strings"
)

func PerformApiCall() {
	fmt.Println()
	fmt.Println("Reseting Export Values....")
	fmt.Println("........................................")
	env.Setenv(env.CA_RESPONSE_BODY, "")

	fmt.Println()
	fmt.Println("Preparing headers....")
	headers := prepareHeaders(env.CA_HEADERS)

	fmt.Println("Calling Api....")
	fmt.Println("........................................")
	success, data := network.CallApi(env.CA_METHOD, env.CA_URL, headers, env.CA_BODY)
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
	env.Setenv(env.CA_RESPONSE_BODY, response)
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
