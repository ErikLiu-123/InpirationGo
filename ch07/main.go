package main

import "fmt"

func main() {
	array := [...]string{"a", "b", "c", "d", "e"}
	array1 := [5]string{1: "b", 3: "d"}
	fmt.Println("array[2]的值：", array[2], "array1[2]的值", array1[2])

	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, array[i])
	}
}
