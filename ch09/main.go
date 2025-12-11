package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[2:5]
	fmt.Println(slice)
	slice[1] = "f"
	fmt.Println(array)

	slice1:=make([]string,4,8)
	fmt.Println(slice1)

	slice2:=[]string{"a","b","c","d","e"}
	fmt.Println(len(slice2),cap(slice2))
	//追加一个元素
	slice2=append(slice2,"f")
	fmt.Println(slice2)
	fmt.Println(len(slice2),cap(slice2))
	//追加多个元素
	slice2=append(slice2,"f","g")
	fmt.Println(slice2)
	fmt.Println(len(slice2),cap(slice2))
	//追加另一个切片
	slice2=append(slice2,slice...)
	fmt.Println(slice2)
	fmt.Println(len(slice2),cap(slice2))

	for i,v:=range slice2{
		fmt.Printf("数组索引：%d,对应值：%s\n", i, v)
	}
}
