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

// 2. 调用函数时的隐形转换
func printType(i interface{}) {
	switch (i).(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	case float64:
		fmt.Println("参数的类型是 f" +
			"loat64")
	}
}

func main() {

	// 如果将phone对象显示声明为Phone1接口类型， 会报错，解决需要隐式声明Phone1接口类型
	// var phone Phone1
	phone := iPhone{
		name: "zero",
	}
	phone.call()
	phone.send_wechat()

	// 调用printType()函数
	name := "王磊"
	printType(name) // 调用函数不会报错

	// 直接在函数外部使用逻辑，会报错
	/*switch name.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	case float64:
		fmt.Println("参数的类型是 float64")
	}*/
}
