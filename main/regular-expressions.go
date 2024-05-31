package main

import (
	"bytes"
	"fmt"
	"regexp"
)

/*
正则表达式
*/

func main() {
	// 测试 模式是否与字符串匹配
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Compile 优化的 Regexp 结构
	r, _ := regexp.Compile("p([a-z]+)ch")
	// 匹配测试
	fmt.Println(r.MatchString("patch"))

	// 找到正则表达式的匹配项
	fmt.Println(r.FindString("peach punch"))

	// 找到第一个匹配项，但返回匹配项的开始和结束索引，而不是匹配的文本
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// Submatch 变体包括有关整个模式匹配和这些匹配中的子匹配的信息。例如，这将返回 p([a-z]+)ch 和 ([a-z]+) 的信息。
	// 也就是能够额外返回([a-z]+) 子匹配的信息
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 也就是能够额外返回([a-z]+) 子匹配的信息，下标索引
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 查找正则表达式的所有匹配项
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 查找正则表达式的所有匹配项，子项；还有他们的下标索引
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// 提供非负整数作为这些函数的第二个参数将限制匹配的数量
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 可以提供 []byte 参数并从函数名称中删除 String
	fmt.Println(r.Match([]byte("peach")))

	// 可以使用 Compile 的 MustCompile 变体。
	//MustCompile 发生panic而不是返回错误，这使得使用全局变量更安全。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// regexp 包还可用于将字符串子集替换为其他值
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func 变体允许您使用给定函数转换匹配的文本
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
