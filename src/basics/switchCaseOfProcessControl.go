package main

import "fmt"

/* 流程控制：switch-case */

// 判断是否有挂科记录
func getResult(args ...int) bool {
	for _, i := range args {
		if i < 60 {
			return false
		}
	}
	return true
}

func main() {

	// case 单条件
	educate := "高中"
	switch educate {
	case "博士后":
		fmt.Println("我是博士后")
	case "博士":
		fmt.Println("我是博士")
	case "研究生":
		fmt.Println("我是研究生")
	case "本科":
		fmt.Println("我是本科生")
	case "专科":
		fmt.Println("我是专科生")
	case "高中":
		fmt.Println("我是高中生")
	default:
		fmt.Println("学历未达标")
	}

	// case 多条件
	month := 2
	switch month {
	case 3, 4, 5:
		fmt.Println("春天")
	case 6, 7, 8:
		fmt.Println("夏天")
	case 9, 10, 11:
		fmt.Println("秋天")
	case 12, 1, 2:
		fmt.Println("冬天")
	default:
		fmt.Println("输入有误")
	}

	// switch 后接收函数
	chinese := 80
	english := 59
	math := 66

	switch getResult(chinese, english, math) {
	// case 后也必须接bool类型
	case true:
		fmt.Println("全部科目合格")
	case false:
		fmt.Println("有挂科记录")
	}

	// switch 穿透
	s := "hello"
	switch s {
	case "hello":
		fmt.Println("hello")
		fallthrough
	case "world":
		fmt.Println("world")
	}

}
