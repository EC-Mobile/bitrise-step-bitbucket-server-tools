package network

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type KeyValue struct {
	Key   string
	Value string
}

const (
	GET  = "GET"
	POST = "POST"
)

func CallApi(method string, url string, headers []KeyValue, body string) (bool, []byte) {
	fmt.Println("Method: " + method)
	fmt.Println("Headers: ")
	for index, header := range headers {
		fmt.Printf("%d: %s: %s\n", index, header.Key, header.Value)
	}
	fmt.Println("Calling Api: " + url)

	fmt.Println("Body: ")
	fmt.Println(body)

	data := []byte(body)
	if len(body) == 0 {
		data = nil
	}

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(data))
	for _, header := range headers {
		req.Header.Set(header.Key, header.Value)
	}
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return false, []byte{}
	}
	fmt.Printf("Api call status: %d\n", response.StatusCode)

	responseData, readError := io.ReadAll(response.Body)
	if readError != nil {
		fmt.Println(readError.Error())
		return false, []byte{}
	}

	if response.StatusCode < 200 || response.StatusCode > 210 {
		return false, []byte{}
	}

	return true, responseData
}
