package main

import "fmt"

func main() {
	fmt.Println(WithReverse("katak"))
	fmt.Println(WithReverse("kasur rusak"))
}

func WithReverse(value string) bool {
	tempValue := ""

	for i := len(value) - 1; i >= 0; i-- {
		tempValue += string(value[i])
	}

	return value == tempValue
}
