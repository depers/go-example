package main

import (
	"crypto/sha256"
	"fmt"
)

/*
SHA256 哈希经常用于计算二进制或文本 blob 的短标识。
例如，TLS/SSL 证书使用 SHA256 来计算证书的签名。
*/
func main() {
	s := "sha256 this string"

	h := sha256.New()

	// Write 需要字节。如果您有字符串 s ，请使用 []byte(s) 将其强制转换为字节。
	h.Write([]byte(s))

	// 得到最终的哈希结果作为字节切片
	bs := h.Sum(nil)

	fmt.Println(s)
	// 运行该程序会计算哈希值并以人类可读的十六进制格式打印它
	fmt.Printf("%x\n", bs)
}
