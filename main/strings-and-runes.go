package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// s 是 string 分配的字面值，Go 字符串文字是 UTF-8 编码的文本。
	const s = "中国"

	// 由于字符串相当于 []byte ，这将产生存储在其中的原始字节的长度。
	fmt.Println("len:", len(s))

	// 对字符串进行索引会在每个索引处生成原始字节值。此循环生成构成 s 中代码点的所有字节的十六进制值。
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()

	// 使用utf8包计算字符串的具体长度
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// range 循环专门处理字符串，并对每个 rune 及其在字符串中的偏移量进行解码。
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// 以通过显式使用 utf8.DecodeRuneInString 函数来实现与range相同的迭代
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// 将 rune 值传递给函数
		examineRune(runeValue)
	}

}

// 用单引号括起来的值是符文文字。我们可以直接将 rune 值与符文文字进行比较。
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
