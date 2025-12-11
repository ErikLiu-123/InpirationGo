package main

import "fmt"

type person struct {
	name string
	age uint
	address
}

type address struct {
	province string
	city string
}

type WalkRun interface {
	Walk()
	Run()
}

func main() {
	p:= person{
		age:20,
		name:"飞雪无情",
		address:address{
			province: "北京",
			city: "北京",
		},
	}
	printWalkRun(&p)
	println(p.province,p.city)
}

func printWalkRun(s WalkRun){
	s.Walk()
	s.Run()
}

func (p *person) Walk() {
	fmt.Printf("%s 能走 \n", p.name)
}

func (p *person) Run() {
	fmt.Printf("%s 能跑 \n", p.name)
}