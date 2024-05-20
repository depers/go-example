package main

import "fmt"

// 枚举是一种具有固定数量的可能值的类型，每个值都有不同的名称。
// Go 没有枚举类型作为独特的语言功能，但使用现有的语言习惯用法很容易实现枚举。

// 枚举类型 ServerState 具有底层 int 类型
type ServerState int

// ServerState 的可能值被定义为常量。特殊关键字iota自动生成连续的常量值；在本例中为 0、1、2 等等。
const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

// 构建一个enum对应map
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// 通过实现 fmt.Stringer 接口，可以打印出 ServerState 的值或将其转换为字符串
// 如果枚举中包含更多的值，就要使用其他方法了
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// 转换模拟服务器的状态转换；它获取现有状态并返回一个新状态。
// 假设我们在这里检查一些谓词来确定下一个状态
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknow state: %s", s))
	}
	return StateConnected
}
