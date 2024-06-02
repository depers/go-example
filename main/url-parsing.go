package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	p := fmt.Println
	// 我们将解析这个示例 URL，其中包括方案、身份验证信息、主机、端口、路径、查询参数和查询片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析 URL 并确保没有错误。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	// 访问方案
	p(u.Scheme)
	p("---")

	// User 包含所有身份验证信息
	p(u.User)
	p(u.User.Username())
	ps, _ := u.User.Password()
	p(ps)
	p("---")

	// Host 包含主机名和端口
	p(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	p(host)
	p(port)
	p("---")

	// 提取 path 和 # 之后的片段
	p(u.Path)
	p(u.Fragment)
	p("---")

	// 要获取 k=v 格式字符串的查询参数，请使用 RawQuery
	p(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"][0])

}
