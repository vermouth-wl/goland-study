package main

import "fmt"

/* 流程控制：if-else */
func main() {
	// 单分支判断
	var age int = 11
	gender := "男性"
	if age >= 18 {
		fmt.Println("已经成年了")
	} else {
		fmt.Println("未成年")
	}

	// 多条件语句
	if age >= 18 && gender == "男性" {
		fmt.Println("成年男性")
	} else {
		fmt.Println("不知道")
	}

	// 多分支判断
	if age >= 12 {
		fmt.Println("已经是青少年了")
	} else if age >= 18 {
		fmt.Println("已经成年了")
	} else {
		fmt.Println("还不是青少年")
	}
}
