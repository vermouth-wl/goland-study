package main

import "fmt"

/* 数据类型：指针 */
func main() {

	// 方法1：定义变量，通过变量取得内存地址，从而创建指针
	name := "王磊"     // 定义变量
	namePtr := &name //使用&取得内存地址创建指针
	fmt.Printf("变量name的值为: %s, 内存地址为: %p\n", name, namePtr)

	// 方法2: 先创建指针，分配好内存后，再给指针指向的内存地址写入对应的值
	astr := new(string) // 创建指针
	*astr = "王磊"        // 给指针赋值
	fmt.Printf("变量astr的值为: %s, 内存地址为: %p\n", *astr, astr)

	// 方法3: 先声明一个指针变量，再从其他变量取得内存地址给他赋值
	age := 16       // 定义变量
	var newAge *int // 声明一个指针
	newAge = &age   // 初始化
	fmt.Printf("变量newAge的值为: %d, 内存地址为: %p\n", *newAge, newAge)

	// 区分*与&
	address := "重庆市沙坪坝区"   // 定义普通变量
	addressPtr := &address // 定义指针变量
	fmt.Println("普通变量存储的是: ", address)
	fmt.Println("普通变量存储的是: ", *addressPtr)
	fmt.Println("指针变量存储的是: ", addressPtr)
	fmt.Println("指针变量存储的是: ", &address)

	// 指针的类型
	astr1 := "hello"
	aint := 123
	abool := false
	arune := 'a'
	afloat := 1.2
	fmt.Printf("astr1 指针类型是: %T\n", &astr1)
	fmt.Printf("aint 指针类型是: %T\n", &aint)
	fmt.Printf("abool 指针类型是: %T\n", &abool)
	fmt.Printf("arune 指针类型是: %T\n", &arune)
	fmt.Printf("afloat 指针类型是: %T\n", &afloat)

	// 指针的零值
	a := 25    // 声明一个普通变量
	var b *int // 声明一个指针变量
	if b == nil {
		fmt.Println("指针b的零值是: ", b)
		b = &a // 将普通变量a的内存地址赋值给b
		fmt.Printf("a的内存地址为: %p\n", &a)
		fmt.Printf("指针b的类型为: %T, 值为: %b, 内存地址为: %p\n", b, *b, b)
	}

}
