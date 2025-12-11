package main

import(
	"fmt"
	//"errors"
)

func main(){
	sum,err:=add(-1,2)
	if cm,ok:=err.(*commonError);ok{
		fmt.Println("错误代码为：",cm.errorCode,"，错误信息为：",cm.errorMsg)
	}else{
		fmt.Println(sum)
	}
}

func add(a,b int) (int,error){
	if a<0 || b<0 {
		return 0,&commonError{
			errorCode:1,
			errorMsg:"a或者b不能为负数"}
	}else {
		return a+b,nil
	}
}

type commonError struct{
	errorCode int //错误码
	errorMsg string//错误信息
}

func (ce *commonError) Error() string{
	return ce.errorMsg
}