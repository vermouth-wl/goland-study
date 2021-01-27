package main

import (
	"fmt"
	"unsafe"
)

/* 数据类型：byte、rune类型 */

func main() {
	// byte类型
	var a byte = 65
	// 八进制写法: var a byte = '\101'
	// 十六进制写法: var a byte = '\x41'

	var b uint = 66
	fmt.Printf("a 的值: %c\nb 的值: %c", a, b)

	var a1 byte = 'A'
	var b1 byte = 'B'
	fmt.Printf("\na1 的值: %c\nb1 的值: %c", a1, b1)

	// rune类型
	var a2 byte = 'A'
	var b2 rune = 'B'
	fmt.Printf("\na2 占用 %d 个字节数\nb2 占用 %d 个字节数", unsafe.Sizeof(a2), unsafe.Sizeof(b2))
}
