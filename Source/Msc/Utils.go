package msc

import (
	"encoding/json"
	"fmt"
)

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func ConvertToJson[T any](data T, pretty bool) string {
	bytes := &data
	if pretty {
		jsonString, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return string(jsonString)
	} else {
		jsonString, err := json.Marshal(bytes)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return string(jsonString)
	}

}
