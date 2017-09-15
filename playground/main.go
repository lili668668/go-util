package main

import "fmt"

type Object struct {
	value interface{}
}

func main() {
	ob1 := Object{1}
	ob2 := Object{"love"}
	ob3 := Object{true}
	ob4 := Object{4.3}

	fmt.Println(ob1.value)
	fmt.Println(ob2.value)
	fmt.Println(ob3.value)
	fmt.Println(ob4.value)

	var arr []interface{}
	arr = append(arr, true)
	arr = append(arr, nil)

	for _, e := range arr {
		fmt.Println(e)
	}
}
