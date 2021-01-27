package main

import "fmt"

/* 异常机制 */

// 捕获异常方法
func setData(x int) {
	defer func() {
		// recover()函数, 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 故意制造数组越界，产生panic
	var arr [10]int
	arr[x] = 100
}
func main() {

	// 手动抛出异常
	// panic("出现异常")

	// 捕获panic
	setData(20)

	// 如果执行以下语句，则证明panic被成功捕获
	fmt.Println("正在正常运行")
}
