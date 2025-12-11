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
	printString(p)
	printString(p.addr)
}

func printString(s fmt.Stringer){
	fmt.Println(s.String())
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d",p.name,p.age)
}

func (addr address) String() string{
	return fmt.Sprintf("the addr is %s%s", addr.province,addr.city)
}

