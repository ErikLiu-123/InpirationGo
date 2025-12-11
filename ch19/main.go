package main

import (
	"fmt"
	"sync"
)

func main() {
	containers:=puchase(50) //采购50台iPad
	//七班人同时运营50台iPad
	images1:=mark(containers)
	images2:=mark(containers)
	images3:=mark(containers)
	images4:=mark(containers)
	images5:=mark(containers)
	images6:=mark(containers)
	images7:=mark(containers)
	//汇聚七个channel成一个
	images:=gather(images1,images2,images3,images4,images5,images6,images7)
	targets:=target(images)//运营它们以便呈现

	//输出测试，看看效果
	for p:=range targets{
		fmt.Println(p)
	}
}

//工序 1 采购
func puchase(n int) <-chan string{
	out:=make(chan string)
	go func(){
		defer close(out)
		for i:=1;i<=n;i++{
			out<-fmt.Sprint("镜像",i)
		}
	}()
	return out
}

//工序 2 运营
func mark(in <-chan string) <-chan string{
	out:=make (chan string)
	go func(){
		defer close(out)
		for c:=range in{
			out<-"运营["+c+"]"
		}
	}()
	return out
	
}

//工序 3 组装成iPad
func target(in <-chan string) <-chan string {
	out:=make(chan string)
	go func() {
		defer close(out)
		for c:=range in{
			out<-"iPad["+c+"]"
		}
	}()
	return out
}

//扇入函数（组件），把多个channel中的数据发送到一个channel中
func gather(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out:=make(chan string)

	//把一个channel中的数据发送到out中
	p:=func(in <-chan string){
		defer wg.Done()
		for c:=range in{
			out <-c
		}
	}

	wg.Add(len(ins))

	//扇入，需要启动多个goroutine用于处理多个channel中的数据
	for _,cs:=range ins{
		go p(cs)
	}

	//等待所有输入的数据ins处理完，再关闭输出out
	go func(){
		wg.Wait()
		close(out)
	}()

	return out
}