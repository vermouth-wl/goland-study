package main

import "fmt"

/* go接口中的三个潜规则 */

// 1. 对方法的调用限制

// 创建一个Phone1接口，接口中有一个call()方法
type Phone1 interface {
	call()
}

// 创建一个iPhone结构体
type iPhone struct {
	name string
}

// 实现接口
func (phone iPhone) call() {
	fmt.Println("hello, phone.")
}

func (phone iPhone) send_wechat() {
	fmt.Println("hello, wechat.")
}

func main() {

	// 如果将phone对象显示声明为Phone1接口类型， 会报错，解决需要隐式声明Phone1接口类型
	// var phone Phone1
	phone := iPhone{
		name: "zero",
	}
	phone.call()
	phone.send_wechat()
}
