package main

import "fmt"

/* 空接口 */

func main() {

	// 空接口类型和值都是nil
	var a interface{}
	fmt.Printf("type: %T, value: %v\n", a, a)

	// 空接口的使用
	// 声明一个空接口
	var i interface{}

	// 存int值
	i = 1
	fmt.Println(i)

	// 存字符串值
	i = "你好, 世界"
	fmt.Println(i)

	// 存布尔值
	i = false
	fmt.Println(i)

	fmt.Println("=================分割线====================")

	// 调用reverser()函数
	name := "王磊"
	age := 16
	address := "重庆市南川区"
	isGirl := false
	reverse(name)
	reverse(age)
	reverse(address)
	reverse(isGirl)

	fmt.Println("=================分割线====================")

	// 接收任意个任意类型的函数
	reverseAll(name, age, address, isGirl)

	fmt.Println("=================分割线====================")

}

// 让函数接收任意类型的值, 使用空接口
func reverse(i interface{}) {
	fmt.Println(i)
}

// 接收任意个任意类型的值
func reverseAll(ifaces ...interface{}) {
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}
