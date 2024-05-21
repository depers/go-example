package main

import (
	"errors"
	"fmt"
)

// 在 Go 中，通过显式的、单独的返回值来传达错误是惯用的。
// 按照约定，错误是最后一个返回值，并且具有类型 error ，这是一个内置接口。
func f(arg int) (int, error) {
	if arg == 42 {
		// errors.New 使用给定的错误消息构造基本的 error 值。
		return -1, errors.New("can't work with 42")
	}

	// 错误位置的 nil 值表示没有错误。
	return arg + 3, nil
}

// 哨兵错误是一个预先声明的变量，用于表示特定的错误情况。
// 用更高级别的错误来包装错误以添加上下文。
// 最简单的方法是使用 fmt.Errorf 中的 %w 动词。包装错误创建一个逻辑链（A 包装 B，B 包装 C 等），可以使用 errors.Is 和 errors.As 等函数进行查询。
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrorPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrorPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		// 在 if 行中使用内联错误检查是很常见的。
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			// errors.Is 检查给定错误（或其链中的任何错误）是否与特定错误值匹配
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrorPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Println("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}

}
