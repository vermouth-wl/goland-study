package main

import "fmt"

/* 流程控制：defer延时 */
func myfunc(a string, b string) {
	fmt.Println(a, b)
}
func manyDefer() {
	name := "go"
	defer fmt.Println(name)

	name = "python"
	defer fmt.Println(name)

	name = "java"
	fmt.Println(name)
}
func main() {

	// 简单调用
	defer myfunc("王", "磊")
	fmt.Println("B")

	// 即使求值的变量快照
	name := "go"
	defer fmt.Println(name)

	name = "python"
	fmt.Println(name)

	// 多个defer执行顺序
	manyDefer()
}
