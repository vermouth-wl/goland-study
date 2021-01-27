package main

import "fmt"

/* 流程控制：for循环 */

func main() {

	// 接一个条件表达式
	a := 1
	for a <= 5 {
		fmt.Print(a, " ")
		a++
	}
	fmt.Println()

	// 接3个表达式
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 不接受表达式，无限往下执行，可以使用break中断
	//for {
	//	fmt.Println("王磊")
	//}
	// 等价于
	//for ;; {
	//	fmt.Println("高二(9)班-王磊")
	//}

	// 接for-range
	myArr := [...]string{"贝尔摩德", "苦艾酒", "神乐"}
	for _, item := range myArr {
		fmt.Printf("你好, %s\n", item)
	}

	// 打印出带索引号的内容
	for i, item := range myArr {
		fmt.Printf("%d-你好：%s\n", i, item)
	}
}
