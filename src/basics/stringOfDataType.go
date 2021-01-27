package main

import "fmt"

/* 数据类型：字符串 */
func main() {
	// 定义字符串
	var a string = "hello"
	var b [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("a: %s\n", a)
	fmt.Printf("b: %s\n", b)

	// 转义字符
	var c string = "\\r\\n"
	fmt.Printf("c: %s\n", c)
	var d string = `\r\n`
	fmt.Printf("d: %s\n", d)
	var e string = `\r\n`
	fmt.Printf("e-\\r\\n的解释性字符串: %q\n", e)

	// 多行字符串
	var f string = `秘密
		使女人更女人`
	fmt.Printf("f: %s\n", f)
}
