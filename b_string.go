package b_string

import (
	"strings"
)

func IndexFrom(str string, find string, from int) int {
	substr := str[from:]
	index := strings.IndexAny(substr, find)
	final_index := from + index
	return final_index
}

func Capture(str, start, end string) string {
	return ""
}

func TrimAll(origin string) string {
	str := strings.Replace(origin, " ", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\f", "", -1)
	return str
}
