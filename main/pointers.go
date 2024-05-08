package main

import "fmt"

// Go 支持指针，允许您传递对程序中的值和记录的引用。

// zeroval 有一个 int 参数，因此参数将按值传递给它。 zeroval 将获得 ival 的副本，该副本与调用函数中的副本不同。
func zeroval(ival int) {
	ival = 0
}

// zeroptr有一个*int参数，这是一个int指针, zeroptr会更改变量i的值，因为它具有对该变量的内存地址的引用。
func zeroptrt(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i 语法给出 i 的内存地址，即指向 i 的指针。
	zeroptrt(&i)
	fmt.Println("zeroptr:", i)

	// 也可以打印指针
	fmt.Println("pointer:", &i)
}
