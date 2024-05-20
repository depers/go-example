package main

import (
	"fmt"
)

// Go 支持在结构类型上定义的方法
type rect struct {
	width, height int
}

// 结构体的指针接收器，area 方法的接收器类型为 *rect
func (r *rect) area() int {
	return r.width * r.height
}

// 结构体的值接收器，perim方法的接收器类型为 rect
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// 调用为我们的结构定义的 2 个方法
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
