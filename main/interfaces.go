package main

import (
	"fmt"
	"math"
)

// 接口是方法签名的命名集合
type geometry interface {
	area() float64
	perim() float64
}

// 在 rect 和 circle 类型上实现此接口
type rect struct {
	width, height float64
}

type circle struct {
	redius float64
}

// 要在Go中实现一个接口，我们只需要实现该接口中的所有方法即可
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.redius
}

// 使用接口方法
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// circle 和 rect 结构类型都实现 geometry 接口，因此我们可以使用这些结构的实例作为 measure 的参数。
func main() {
	r := rect{width: 3, height: 4}
	c := circle{redius: 5}

	measure(r)
	measure(c)
}
