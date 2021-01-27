package main

import "fmt"

/* 数据类型：数组、切片 */
func main() {
	// 声明数组：方法一
	var arry1 [3]string
	arry1[0] = "神乐"
	arry1[1] = "神无"
	arry1[2] = "桔梗"
	fmt.Printf("arry1: %s\n", arry1)

	// 声明数组：方法二
	var arry2 [3]string = [3]string{"桔梗", "神乐", "桔梗"}
	fmt.Printf("arry2: %s\n", arry2)

	// 声明数组：方法三
	arry3 := [3]string{"神无", "桔梗", "神乐"}
	fmt.Printf("arry3: %s\n", arry3)

	// 查看不同类型长度数组的类型是否一致
	arry4 := [...]int{1, 2, 3, 4}
	arry5 := [...]int{1, 2, 3}
	fmt.Printf("%d 的类型是: %T\n", arry4, arry4)
	fmt.Printf("%d 的类型是: %T\n", arry5, arry5)

	// 切片
	// 定义一个数组
	arry6 := [...]string{"A", "B", "C", "D", "E", "F"}
	// 切片一
	sli1 := arry6[0:3]
	fmt.Printf("sli1: %s\n", sli1)
	fmt.Printf("sli1 长度为: %d, 容量为: %d\n", len(sli1), cap(sli1))
	// 切片二
	sli2 := arry6[0:3:4]
	fmt.Printf("sli2: %s\n", sli2)
	fmt.Printf("sli2 长度为: %d, 容量为: %d\n", len(sli2), cap(sli2))

	// 使用make()函数构建切片
	sli3 := make([]string, 2)
	sli4 := make([]string, 2, 10)
	fmt.Println(sli3, sli4)
	fmt.Println(len(sli3), len(sli4))
	fmt.Println(cap(sli3), cap(sli4))

	// 切片默认值
	var sli5 []int
	fmt.Println(sli5 == nil)

	// 添加元素
	myarr := []int{1}
	// 追加一个元素
	myarr = append(myarr, 2)
	// 追加多个元素
	myarr = append(myarr, 3, 4)
	// 追加一个切片， ...表示解包，不能省略
	myarr = append(myarr, []int{7, 8}...)
	// 在第一个位置插入元素
	myarr = append([]int{0}, myarr...)
	// 在中间插入一个切片
	myarr = append(myarr[:5], append([]int{5, 6}, myarr[5:]...)...)
	fmt.Println(myarr)

	// 思考
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Printf("myslice 为: %d, 长度为: %d, 容量为: %d\n", myslice, len(myslice), cap(myslice))
	// myslice的实际下标为0，实际对应numbers下标4，以此类推
	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice的第四个元素为: %d\n", myslice[3])

}
