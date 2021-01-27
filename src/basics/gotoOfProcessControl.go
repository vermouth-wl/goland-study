package main

import "fmt"

/* 流程控制：goto */
func main() {
	// 简单跳转
	goto flag
	fmt.Println("打印A")
flag:
	fmt.Println("打印B")

	// 使用goto语句实现循环打印1-5
	i := 1
flag1:
	if i <= 5 {
		fmt.Print(i, " ")
		i++
		goto flag1
	}
	fmt.Println()

	// 使用goto 语句实现类似循环中break效果
	i2 := 1
	for {
		if i2 > 5 {
			goto flag2
		} else {
			fmt.Print(i2, " ")
			i2++
		}
	}
flag2:
	fmt.Println()

	// 使用goto 实现类似continue效果
	i3 := 1
flag3:
	for i3 <= 10 {
		if i3%2 == 0 {
			fmt.Print(i3, " ")
			i3++
		}
		i3++
		goto flag3
	}
	fmt.Println()
}
