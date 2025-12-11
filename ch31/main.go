package main

import(
	"fmt"
	"errors"
)

type MyError struct {
	err error
	msg string
}

func (e *MyError) Error() string {
	return e.err.Error() + e.msg
}

func main() {
	//err是一个存在的错误，可以从另外一个函数返回
	//newErr:= MyError(err, "数据上传问题")
	e:=errors.New("原始错误e")
	w:=fmt.Errorf("Wrap了一个错误：%w",e)
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w))
	fmt.Println(errors.Is(w,e))
}
