package main

import (
	"encoding/xml"
	"fmt"
)

/*
Go 通过 encoding.xml 包提供对 XML 和类 XML 格式的内置支持
*/

// XMLName 字段名称指示表示该结构的 XML 元素的名称；
// id,attr 表示 Id 字段是 XML 属性而不是嵌套元素。
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

// 复写了String()方法
func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	// 生成xml
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	fmt.Println("---------------")

	// 要将通用 XML 标头添加到输出中，请显式附加它
	fmt.Println(xml.Header + string(out))
	fmt.Println("---------------")

	var p Plant
	// 使用 Unmarshal 将带有 XML 的字节流解析为数据结构。
	// 如果 XML 格式错误或无法映射到 Plant，将返回描述性错误。
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
	fmt.Println("---------------")

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		// parent>child>plant 字段标记告诉编码器将所有 plant 嵌套在 <parent><child>... 下
		Plants []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))

}
