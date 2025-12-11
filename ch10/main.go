package main

import "fmt"

func main() {
	nameAgeMap:=make(map[string]int)
	nameAgeMap["飞雪无情"]=20

	nameAgeMap1:=map[string]int{"飞雪无情":20}

	//添加键值对或者更新对应 Key 的 Value
	nameAgeMap1["飞雪无情"]=30

	//获取指定 Key 对应的 Value
	age:=nameAgeMap["飞雪无情"]

	nameAgeMap2:=make(map[string]int)
	nameAgeMap2["飞雪无情"]=20

	age,ok:=nameAgeMap["飞雪无情"]
	if ok{
		fmt.Println(age)
	}

}
