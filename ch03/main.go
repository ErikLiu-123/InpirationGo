package main

import "fmt"

func main() {
	switch j := 1; j {
	case 1:
		fallthrough
	case 2:
		fmt.Println("1")
	case 3:
		fmt.Println("没有匹配")
	}
}
