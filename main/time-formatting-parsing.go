package main

import (
	"fmt"
	"time"
)

/*
Go 通过基于模式的布局支持时间格式化和解析。
*/
func main() {
	p := fmt.Println

	t := time.Now()
	// 根据 RFC3339 使用相应的布局常量格式化时间的基本示例
	p(t.Format(time.RFC3339))
	p("------------")

	// 时间解析使用与 Format 相同的布局值。
	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	if e != nil {
		panic(e)
	}
	p(t1)
	p("------------")

	// 使用Format方法中的布局参数表示时间
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	// 使用Parse方法中的布局参数表示时间，第一个参数是布局参数，第二个是要表示的时间
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)
	p("------------")

	// 可以使用标准字符串格式提取时间
	fmt.Println("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	p("------------")

	// Parse 将在格式错误的输入上返回错误，解释解析问题。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
