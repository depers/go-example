package main

import (
	"fmt"
	"time"
)

/**
  switch语句
*/

func main() {

	// 这是一个基本的switch语句
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 使用逗号分隔同一 case 语句中的多个表达式。在此示例中，我们也使用可选的 default 情况。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekeed")
	default:
		fmt.Println("It's a weekday")
	}

	// 不带表达式的 switch 是表达 if/else 逻辑的另一种方式。这里我们还展示了 case 表达式如何可以是非常量
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// 类型 switch 比较类型而不是值。您可以使用它来发现接口值的类型。在此示例中，变量 t 将具有与其子句相对应的类型。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Println("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
