package main

import (
	"os"
	"text/template"
)

/*
go语言通过text/template包向用户提供动态创建内容的功能
*/
func main() {
	t1 := template.New("t1")
	// {{.}}用于动态的插入内容
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// 通过使用 template.Must，可以简化错误处理，它会在发生错误时引发 panic
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// 使用附中函数创建模板
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// 如果数据是结构体或是map
	t2 := Create("t2", "Name:{{.Name}}\n")

	// 如果数据是结构体
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// 如果数据是map
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// f/else 为模板提供条件执行，下面的'-'表示空
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// 范围块让我们循环遍历切片、数组、映射或通道。范围块内的 {{.}} 设置为迭代的当前项
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

}
