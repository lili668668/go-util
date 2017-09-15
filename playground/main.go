package main

import "fmt"
import "strconv"

type Object struct {
	type_name string
	value     string
}

func (this Object) Value(value interface{}) {
	switch value.(type) {
	case *string:
		*value = this.value
	case *int:
		i, _ := strconv.ParseInt(this.value, 10, 64)
		*value = i
	case *float32:
		f, _ := strconv.ParseFloat(this.value, 32)
		*value = f
	}
}

func main() {

	arr := []Object{Object{"you"}, Object{"1"}, Object{"2.3"}}

	str := ""
	num := 0
	flo := 1.1

	arr[0].Value(&str)
	arr[1].Value(&num)
	arr[2].Value(&flo)

	fmt.Println(num)
	fmt.Println(str)
	fmt.Println(flo)

}
