package app

import "fmt"

func NewTesting(str1 string) {
	for _, ch := range str1 {
		fmt.Println(string(ch))
	}
}
