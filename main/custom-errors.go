package main

import (
	"errors"
	"fmt"
)

// 自定义错误类型通常带有后缀“Error”。
type argError struct {
	arg     int
	message string
}

// 通过在自定义类型上实现 Error() 方法，可以将自定义类型用作 error
// 添加此 Error 方法使 argError 实现 error 接口。
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		// 返回我们的自定义错误。
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	var ae *argError
	// errors.As 是 errors.Is 的更高级版本。它检查给定错误（或其链中的任何错误）是否与特定错误类型匹配，并转换为该类型的值，返回 true 。如果没有匹配，则返回 false 。
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
