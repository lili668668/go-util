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
	index_start := strings.IndexAny(str, start)
	index_end := strings.LastIndexAny(str, end)

	if index_start == -1 || index_end == -1 {
		return str
	}

	new_str := str[index_start+1 : index_end]
	return new_str
}

func CaptureFrom(str string, start string, end string, from int) string {
	substr := str[from:]
	new_str := Capture(substr, start, end)
	return new_str
}

func TrimAll(origin string) string {
	str := strings.Replace(origin, " ", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\f", "", -1)
	return str
}
