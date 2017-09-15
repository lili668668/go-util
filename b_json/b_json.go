package b_json

import "strconv"
import "errors"
import "strings"
import "../b_string"

type Object struct {
	type_name string
	value     string
}

func (this Object) String() string {
	return this.value
}

func (this Object) Int() (int, error) {
	i, err := strconv.ParseInt(this.value, 10, 64)
	return i, err
}

func (this Object) Float() (float32, error) {
	f, err := strconv.ParseFloat(this.value, 32)
	return f, err
}

func (this Object) Bool() (bool, error) {
	b, err := strconv.ParseBool(this.value)
	return b, err
}
