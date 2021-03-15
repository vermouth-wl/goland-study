package main

import (
	"fmt"
	"reflect"
)

/* 反射中的函数 */
func main() {

	// 1、获取类型函数Kind()类型，注意：获取类型为Go原生类型，而非自定义类型
	// 实例化User结构体
	user := PersonClass{
		name:    "贝尔",
		age:     17,
		address: "重庆市",
	}

	// 将user转换为反射对象
	userReflectType := reflect.TypeOf(&user)
	fmt.Println("&userReflectType Type: ", userReflectType)              // 获取userReflectType的类型
	fmt.Println("&userReflectType Kind: ", userReflectType.Kind())       //获取userReflectType的基本类型
	fmt.Println("userReflectType Type: ", userReflectType.Elem())        // 获取&userReflectType指向的数据类型
	fmt.Println("userReflectType Kind: ", userReflectType.Elem().Kind()) // 获取&userReflectType指向的基本类型

	// 2、进行类型转换
	// 转换成int
	var age int = 17
	v1 := reflect.ValueOf(age)
	fmt.Printf("转换前, type: %T, value: %v\n", v1, v1)
	v2 := v1.Int()
	fmt.Printf("转换后, Type: %T, value: %v\n", v2, v2)

	// 转换成float
	var score float64 = 99.5
	v3 := reflect.ValueOf(score)
	fmt.Printf("转换前, Type: %T, value: %v\n", v3, v3)
	v4 := v3.Float()
	fmt.Printf("转换后, Type: %T, value: %v\n", v4, v4)

	// 转换成string
	var name string = "苦艾酒"
	v5 := reflect.ValueOf(name)
	fmt.Printf("转换前, Type: %T, value: %v\n", v5, v5)
	v6 := v5.String()
	fmt.Printf("转换后, Type: %T, value: %v\n", v6, v6)

	// 转换成布尔值
	var isMale bool = true
	v7 := reflect.ValueOf(isMale)
	fmt.Printf("转换前, Type: %T, value: %v\n", v7, v7)
	v8 := v7.Bool()
	fmt.Printf("转换后, Type: %T, value: %v\n", v8, v8)

	// 转换成指针
	var num int = 50000
	v9 := reflect.ValueOf(&num)
	fmt.Printf("转换前, Type: %T, value: %v\n", v9, v9)
	v10 := v9.Pointer()
	fmt.Printf("转换后, Type: %T, value: %v\n", v10, v10)

	// 转换成接口
	var sum int = 5201314
	v11 := reflect.ValueOf(sum)
	fmt.Printf("转换前, Type: %T, value: %v\n", v11, v11)
	v12 := v11.Interface()
	fmt.Printf("转换后, Type: %T, value: %v\n", v12, v12)

	// 3、对切片的操作
	// 对切片再切片（两下标）
	var numList []int = []int{1, 2}
	v13 := reflect.ValueOf(numList)
	fmt.Printf("转换前, Type: %T, value: %v\n", v13, v13)
	// Slice函数接收两个参数
	v14 := v13.Slice(1, 2)
	fmt.Printf("转换后: Type: %T, value: %v\n", v14, v14)

	// Set()和Append()更新切片
	arr := []int{1, 2, 3}
	appendToSlice(&arr)
	fmt.Println(arr)

	// 4、对属性的操作
	// 实例化PersonClass对象
	p := PersonClass{
		name:    "王磊",
		age:     17,
		address: "重庆市",
	}
	v15 := reflect.ValueOf(p)
	fmt.Println("字段数: ", v15.NumField())
	fmt.Println("第1个字段: ", v15.Field(0))
	fmt.Println("第2个字段: ", v15.Field(1))
	fmt.Println("第3个字段: ", v15.Field(2))

	fmt.Println("=========================")

	for i := 0; i < v15.NumField(); i++ {
		fmt.Printf("第%d个字段是: %v\n", i+1, v15.Field(i))
	}

	fmt.Println("=========================")

	// 5、对方法的操作
	p2 := &PersonClass{
		name:    "嘤嘤嘤",
		age:     17,
		address: "北京市",
	}
	v16 := reflect.TypeOf(p2)
	/*fmt.Println("方法数（可导出的）: ", v16.NumMethod())
	fmt.Println("第1个方法: ", v16.Method(0).Name)
	fmt.Println("第2个方法: ", v16.Method(1).Name)*/

	// 动态调用函数（使用索引且无参数）
	v17 := reflect.ValueOf(p2)

	for i := 0; i < v17.NumField(); i++ {
		fmt.Printf("调用 %d 个方法: %v, 调用结果是: %v\n",
			i+1, v16.Method(i).Name, v17.Elem().Method(i).Call(nil))
	}
}

// 创建结构体
type PersonClass struct {
	name    string
	age     int
	address string
}

// 创建添加切片操作函数
func appendToSlice(arrPtr interface{}) {
	valuePtr := reflect.ValueOf(arrPtr)
	value := valuePtr.Elem() // 设置可写性
	value.Set(reflect.Append(value, reflect.ValueOf(4)))

	fmt.Println(value)
	fmt.Println(value.Len())
}

func (p PersonClass) sayBye() {
	fmt.Println("Bye")
}
func (p PersonClass) sayHello() {
	fmt.Println("Hello")
}
