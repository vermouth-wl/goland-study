package main

import (
	"fmt"
	"reflect"
)

/* 反射 */

func main() {

	fmt.Println("----------------------------------------------")

	// 为了实现从接口变量到反射对象的转换，需要使用reflect包里面的TypeOf()和ValueOf()方法
	// 查看reflect.TypeOf()方法和reflect.valueOf()方法的使用
	var age interface{} = 25
	fmt.Printf("原始接口变量的类型为: %T, 值为: %v\n", age, age)

	// 使用TypeOf()和ValueOf()对接口变量进行转换为反射对象
	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象: Type对象的类型为: %T\n", t)
	fmt.Printf("从接口变量到反射对象: Value对象的类型为: %T\n", v)

	fmt.Println("----------------------------------------------")

	// 为了实现从反射对象到接口变量的转换，需要使用value中的Interface()方法
	// 从反射对象到接口变量，注意，只能将Value进行逆向转换为接口变量，Type不行
	i := v.Interface()
	fmt.Printf("从反射对象到接口变量：新对象的类型为: %T, 值为: %v\n", i, i)

	fmt.Println("----------------------------------------------")

	// 可写性
	// 使用CanSet()方法判断反射对象是否具有可写性
	var name string = "坍缩的宇宙"

	// 将接口变量转换为反射对象
	v1 := reflect.ValueOf(name)
	fmt.Println("v1可写性为: ", v1.CanSet())

	// 要使反射对象具有可写性，必须满足创建反射对象时接收变量必须是指针类型且使用Elem()函数返回指针指向的数据
	v2 := reflect.ValueOf(&name)
	v3 := v2.Elem()
	fmt.Println("v3可写性为: ", v3.CanSet())

	// 对反射对象进行更新，查看真实值的变化
	v3.SetString("量子力学")
	fmt.Println("通过反射对象v3进行更新后，真实世界的name变为: ", name)
}
