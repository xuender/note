package main

import "fmt"

const (
	T1 = 1 << iota
	T2
	T3
)

func main() {
	tag := 0
	// 增加标签
	tag |= T1
	tag |= T2
	fmt.Println("增加标签后", tag)
	// 判断是否有标签
	fmt.Println(tag&T1 > 0)
	fmt.Println(tag&T2 > 0)
	fmt.Println(tag&T3 > 0)
	// 去掉标签
	tag ^= T1
	fmt.Println("去掉标签后", tag)
	fmt.Println(tag&T1 > 0)
	fmt.Println(tag&T2 > 0)
	fmt.Println(tag&T3 > 0)

	tag2 := 0
	tag2 |= T2
	tag2 |= T3
	fmt.Println("另一个标前", tag2)
	fmt.Println("有交集", tag&tag2, tag&tag2 > 0)
}
