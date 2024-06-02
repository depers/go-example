package main

import (
	"fmt"
	"math/rand/v2"
)

/*
Go 的 math/rand/v2 包提供伪随机数生成。
*/
func main() {
	p := fmt.Print
	p2 := fmt.Println

	//  rand.IntN 返回随机 int n, 0 <= n < 100
	p(rand.IntN(100), ",")
	p(rand.IntN(100))
	p2()

	// rand.Float64 返回 float64 f 、 0.0 <= f < 1.0
	p(rand.Float64())
	p2()

	// 这可用于生成其他范围内的随机浮点数，例如 5.0 <= f' < 10.0 。
	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	p2()

	// 创建随机种子，生成的随机数是一样的
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	p(r2.IntN(100), ",")
	p(r2.IntN(100))
	p2()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	p(r3.IntN(100), ",")
	p(r3.IntN(100))
}
