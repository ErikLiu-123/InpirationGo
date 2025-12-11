package main

import (
	"fmt"
	//"time"
	"sync"
)

//共享的资源
var (
	sum int
	mutex sync.RWMutex
)

func main() {
	run()
}

func run(){
	var wg sync.WaitGroup
	//因为要监控110个协程，所以设置计数器为110
	wg.Add(110)
	//开启100个协程让sum+10
	for i:=0;i<100;i++{
		go func(){
			//计数器值减1
			defer wg.Done()
			add(10)
		}()
	}
	for i:=0;i<10;i++{
		go func(){
			//计数器值减1
			defer wg.Done()
			fmt.Println("和为：",readSum())
		}()
	}
	//一直等待，直到计数器值为0
	wg.Wait()
	
	//防止提前退出
	//time.Sleep(2 * time.Second)
}

func add(i int){
	mutex.Lock()
	defer mutex.Unlock()
	sum+=i
}

func readSum() int{
	//只获取读锁
	mutex.RLock()
	defer mutex.RUnlock()
	b:=sum
	return b
}