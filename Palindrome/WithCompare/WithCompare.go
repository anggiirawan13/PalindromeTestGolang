package main

import "fmt"

func main() {
	fmt.Println(WithCompare("katak"))
	fmt.Println(WithCompare("kasur rusak"))

	fmt.Println(WithCompareV2("katak"))
	fmt.Println(WithCompareV2("kasur rusak"))
}

func WithCompare(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}

	return true
}

func WithCompareV2(value string) bool {
	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}

	return true
}
