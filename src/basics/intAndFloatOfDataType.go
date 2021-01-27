package main

import (
	"fmt"
)

/* 数据类型：整型、浮点型 */

func main() {
	// 整型
	var num01 int = 0b1100 // 二进制
	var num02 int = 0o14   // 八进制
	var num03 int = 0xC    // 十六进制

	fmt.Printf("二进制数 %b 表示的是 %d\n", num01, num01)
	fmt.Printf("八进制数 %o 表示的是 %d\n", num02, num02)
	fmt.Printf("十六进制数 %X 表示的是 %d\n", num03, num03)

	// 浮点型-满足精度
	var myfloat float32 = 10000018
	fmt.Println("myfloat: ", myfloat)
	fmt.Println("myfloat: ", myfloat+1)

	// 浮点型-不满足精度
	var myfloat01 float32 = 100000182
	var myfloat02 float32 = 100000187
	fmt.Println("myfloat01: ", myfloat01)
	fmt.Println("myfloat01 + 5: ", myfloat01+5)
	fmt.Println(myfloat01+5 == myfloat02)
	fmt.Println("myfload01 addr: ", &myfloat01)
	fmt.Println("myfloat02 addr: ", &myfloat02)
	fmt.Println(myfloat01 == myfloat02) // 单精度不足会导致截断，需要使用双精度浮点类型
	fmt.Println(&myfloat01 == &myfloat02)
}
