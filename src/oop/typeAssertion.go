package main

import "fmt"

/* 类型断言 */

func main() {
	//var i interface{} = 10
	// 断言接口对象(i)里不是nil，并且接口对象(i)存储类型是T
	//t1 := i.(int)
	//fmt.Println(t1)
	//
	//fmt.Println("======分割线======")

	//t2 := i.(bool)
	//fmt.Println(t2)

	// 断言接口对象为nil
	//var i1 interface{}
	//var _ = i1.(interface{})

	var i interface{} = 10
	t1, ok := i.(int)
	fmt.Printf("%d-%t\n", t1, ok)

	fmt.Println("==================分割线1====================")

	t2, ok := i.(string)
	fmt.Printf("%s-%t\n", t2, ok)

	fmt.Println("==================分割线2====================")

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok)

	fmt.Println("==================分割线3====================")

	k = 10
	t4, ok := k.(interface{})
	fmt.Printf("%d-%t\n", t4, ok)

	fmt.Println("==================分割线4====================")

	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok)

	fmt.Println("==================分割线4====================")

	// 执行findType()函数
	var a interface{}
	findType(a)
	findType(10)
	findType("王磊")
	findType(false)
	findType(131.4)
	findType(13.14)

}
func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case bool:
		fmt.Println(x, "is bool")
	case nil:
		fmt.Println(x, "is nil")
	case float32:
		fmt.Println(x, "is float32")
	case float64:
		fmt.Println(x, "is float64")
	default:
		fmt.Println("老子也不知道")
	}
}
