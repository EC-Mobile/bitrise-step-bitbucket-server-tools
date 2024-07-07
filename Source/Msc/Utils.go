package msc

import (
	"encoding/json"
	"fmt"
	"image/color"
	"regexp"
	"strings"
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

func Matches(regex string, value string) bool {
	if len(regex) <= 0 {
		return true
	}
	matchFlag := true
	splits := strings.Split(regex, " ")
	if len(splits) >= 2 && splits[0] == "NRR" {
		matchFlag = false
		regex = strings.Replace(regex, "NRR ", "", 1)
	}

	valueRegex, _ := regexp.Compile(regex)
	matchFound := len(valueRegex.FindAllString(value, 1)) > 0

	return matchFound == matchFlag
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}
