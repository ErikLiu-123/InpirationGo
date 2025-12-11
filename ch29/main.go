package main

import "fmt"
import "errors"

func main() {
	sum,err:=add(-1,2)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(sum)
	}
}

func add(a,b int) (int,error){
	if a<0 || b<0 {
		return 0,errors.New("a或者b不能为负数")
	}else {
		return a+b,nil
	}
}