package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}

	for i, v := range array {
		fmt.Printf("数组索引：%d,对应值：%s\n", i, v)
	}

	for _, v := range array {
		fmt.Printf("对应值:%s\n:", v)
	}
}
