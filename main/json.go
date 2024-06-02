package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
Go 提供对 JSON 编码和解码的内置支持，包括内置和自定义数据类型的传入和传出。
*/

type response1 struct {
	Page   int
	Fruits []string
}

// 只有导出的字段才会以 JSON 进行编码/解码。要导出的字段必须以大写字母开头。
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// -----------------------------------------------将go的值编码为JSON数据
	// 布尔值编码为json字符串
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	// 整型编码为json字符串
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	// 浮点型编码为json字符串
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	// 字符串编码为json字符串
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 切片编码为json字符串
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	// map映射编码为json字符串
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// 将自定义类型编码为json字符串
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 可以在结构体字段声明上使用标签来自定义编码的 JSON 键名称
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// -----------------------------------------------将JSON数据解码为Go值
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// map[string]interface{} 将保存字符串到任意数据类型的映射。
	var dat map[string]interface{}

	// 进行解码，并检查错误
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 将解码map中的值转换为适当的类型
	num := dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套类型
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 将 JSON 解码为自定义数据类型
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 在上面的示例中，我们始终使用字节和字符串作为标准输出上数据和 JSON 表示之间的中间体。
	// 我们还可以将 JSON 编码直接流式传输到 os.Writer ，例如 os.Stdout 甚至 HTTP 响应主体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
