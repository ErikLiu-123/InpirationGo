package main

import(
	"fmt"
	"time"
)

func main() {
	vegetabelsCh:=washVagetables() //洗菜
	waterCh:=boilWater() //烧水

	fmt.Println("已经安排洗菜和烧水了，我先眯一会")
	time.Sleep(2*time.Second)

	fmt.Println("要做火锅了，看看菜和水好了吗")
	vegetabels:=<-vegetabelsCh
	water:=<-waterCh
	fmt.Println("准备好了，可以做火锅了:",vegetabels,water)
}

//洗菜
func washVagetables() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5*time.Second)
		vegetables<-"洗好的菜"
	}()
	return vegetables
}

//烧水
func boilWater() <-chan string {
	water:=make(chan string)
	go func() {
		time.Sleep(5*time.Second)
		water <- "烧开的水"
	}()
	return water
}