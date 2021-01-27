package main

import "fmt"

/* 变量 */

// 多变量声明
var (
	username string
	password string
	address  string
	age      int
)

// 匿名变量
func getInt() (int, int) {
	return 100, 200
}

func main() {
	// 单变量声明
	var text string = "你好，世界"
	fmt.Println(text)

	// 多变量赋值
	username, password, address, age := "王磊", "123456", "北斗星域", 24
	email := "1591934284@qq.com"
	fmt.Println(username, password, address, age, email)

	// 指针变量
	year := 2020
	var ptr = &year // &后面接变量名，表示取出该变量的内存地址
	fmt.Println("year: ", year)
	fmt.Println("ptr: ", ptr)

	ptr2 := new(int)
	fmt.Println("ptr2 address: ", ptr2)
	fmt.Println("ptr2 value: ", *ptr2) // *后面接指针变量，表示从地址中取出值

	// 匿名变量
	a, _ := getInt()
	_, b := getInt()
	fmt.Println(a, b)

}
