package main

import (
	"fmt"
	"slices"
)

/*
Go 的 slices 包实现了内置类型的排序
*/

func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted:", s)
}
