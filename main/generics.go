package main

import "fmt"

// 泛型方法
/**
MapKeys 接受任何类型的映射并返回其键的切片
该函数有两个类型参数 - K 和 V ； K 具有 comparable 约束，这意味着我们可以将此类型的值与 == 和 != 运算符进行比较。这是 Go 中的映射键所必需的。
V 具有 any 约束，这意味着它不受任何方式的限制（ any 是 interface{} 的别名）。
*/
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}

	return r
}

// 作为泛型类型的示例， List 是一个包含任意类型值的单链表。
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// 泛型类型的结构体方法
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	// 在调用 MapKeys 时，我们不必指定 K 和 V 的类型 - 编译器会自动推断它们。
	// Go 中没有定义映射键的迭代顺序，因此不同的调用可能会导致不同的顺序。
	fmt.Println("keys:", MapKeys(m))

	// 也可以明确指定它们
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list", lst.GetAll())
}
