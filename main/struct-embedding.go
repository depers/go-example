package main

import "fmt"

/**
Go 支持结构体和接口的嵌入，以表达更无缝的类型组合。
*/

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container 嵌入 base 。嵌入看起来就像一个没有名称的字段。
type container struct {
	base
	str string
}

func main() {

	// 当使用文字创建结构体时，我们必须显式初始化嵌入；这里嵌入的类型用作字段名称。
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// 直接在 co 上访问base的字段，例如 co.num
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// 可以使用嵌入的类型名称拼出完整路径
	fmt.Println("also num:", co.base.num)

	// 由于 container 嵌入 base ，因此 base 的方法也成为 container 的方法。在这里，我们调用直接从 base 嵌入 co 的方法。
	fmt.Println("describe:", co.describe())

	// 在这里我们看到 container 现在实现了 describer 接口，因为它嵌入了 base 。
	type describer interface {
		describe() string
	}

	// 在这里我们看到 container 现在实现了 describer 接口，因为它嵌入了 base 。
	var d describer = co
	fmt.Println("describer:", d.describe())
}
