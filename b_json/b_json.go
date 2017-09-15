package b_json

import "strconv"
import "errors"
import "strings"
import "../b_string"

func ParseToObject(str string) interface{} {
	obj := make(map[string]interface{})
	var arr []interface{}
	var rbc, lbc, rsc, lsc int
	for pointer := 0; pointer < str; {
		switch str[pointer] {
		case "{":
			substr := b_string.Capture(str, "{", "}")
			return ParseToObject(substr)
		case "[":
			substr := b_string.Capture(str, "[", "]")
			return ParseToObject(substr)
		default:
			mark, index = b_string.MinIndex(str, ":", ",")
			switch mark {
			case ":":
				key := str[pointer:index]
				value_pointer := index + 1
				switch mark := str[value_pointer] {
				case "{" || "[":
					substr := str[value_pointer:]
					value_str := FindEndAndCapture(substr, mark)
					value := ParseToObject(value_str)
				default:

				}

			case ",":
			default:
				return nil
			}
		}

	}
}

func FindEndAndCapture(str string, mark string) string {

}
