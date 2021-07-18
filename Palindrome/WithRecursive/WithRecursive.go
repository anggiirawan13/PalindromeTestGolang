package main

import "fmt"

func main() {
	fmt.Println(WithRecursive("katak", 0))
	fmt.Println(WithRecursive("kasur rusak", 0))

	fmt.Println(WithRecursiveV2("katak", 0))
	fmt.Println(WithRecursiveV2("kasur rusak", 0))
}

func WithRecursive(value string, indexAwal int) bool {
	if indexAwal < len(value) {
		if value[indexAwal] != value[len(value)-indexAwal-1] {
			return false
		} else {
			return WithRecursive(value, indexAwal+1)
		}
	} else {
		return true
	}
}

func WithRecursiveV2(value string, indexAwal int) bool {
	if indexAwal < len(value)/2 {
		if value[indexAwal] != value[len(value)-indexAwal-1] {
			return false
		} else {
			return WithRecursiveV2(value, indexAwal+1)
		}
	} else {
		return true
	}
}
