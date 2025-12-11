package main

import "fmt"

func main(){
	cacheCh:=make(chan int, 5)
	cacheCh <- 2
	cacheCh <- 3
	fmt.Println("cacheCh容量为：",cap(cacheCh),",元素个数为：",len(cacheCh))
	close(cacheCh)

	//onlySend:=make(chan<- int)
	//onlyReceive:=make(<-chan int)
}

func counter(out chan<- int){
	//函数体内使用变量out，只能进行发送操作
}