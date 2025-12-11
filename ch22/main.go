package main

import "fmt"

type person struct{
	name string 
	age uint

	addr address
}

type address struct{
	province string
	city string
}

func main() {
	p:=person{
		age:30,
		name:"飞雪无情",
		addr:address{
			province: "北京",
			city:     "北京",
		},
	}
	fmt.Println(p.name,p.age,p.addr.province)
}


