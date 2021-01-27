package main

import "fmt"

/* 数据类型：字典、布尔 */
func main() {
	// 声明初始化字典

	// 方法1
	var user1 map[string]string = map[string]string{"username": "王磊", "password": "123456"}
	fmt.Printf("字典user1为: %s\n", user1)

	// 方法2
	user2 := map[string]string{"username": "kagura", "password": "11111"}
	fmt.Printf("字典user2为: %s\n", user2)

	// 方法3
	user3 := make(map[string]string)
	user3["username"] = "神乐"
	user3["address"] = "重庆"
	fmt.Printf("字典user3为: %s\n", user3)

	// 使用方法1声明、初始化、赋值
	// 声明一个名为dict1的字典
	var dict1 map[string]string
	// 直接对字典赋值会出错
	//dict1["dict_code"] = "001"
	//dict1["dict_name"] = "性别"
	//fmt.Printf("字典dict1为: %s\n", dict1)
	// 未初始化的dict1的零值为nil，无法直接进行赋值
	if dict1 == nil {
		// 需要使用make()函数先对其进行初始化
		dict1 = make(map[string]string)
	}
	// 经过初始化后，便可以进行赋值
	dict1["dict_code"] = "001"
	dict1["dict_name"] = "性别"
	fmt.Printf("字典dict1为: %s\n", dict1)

	// 字典相关操作
	// 添加元素
	dict1["flag"] = "启用"
	// 更新元素
	dict1["flag"] = "作废"
	// 读取元素
	fmt.Println(dict1["flag"])
	fmt.Println(dict1["hello"])
	// 删除元素
	delete(dict1, "flag")
	fmt.Println(dict1["flag"])

	// 判断key是否存在
	county := map[string]string{"name1": "china", "name2": "america"}
	name3, ok := county["name3"]
	if ok {
		fmt.Printf("name3的值是: %s\n", name3)
	} else {
		fmt.Println("name3不存在")
	}

	// 对字典进行循环
	// 循环出key, value
	order := map[string]string{"id": "1", "code": "DZ001223", "name": "MacBook Pro", "price": "￥12999"}
	for key, value := range order {
		fmt.Printf("Key: %s, value: %s\n", key, value)
	}

	// 循环出key, 不使用占位符
	for key := range order {
		fmt.Printf("key: %s\n", key)
	}

	// 循环出value, 需要使用占位符
	for _, value := range order {
		fmt.Printf("value: %s\n", value)
	}

	// 布尔类型
	var male bool = true
	fmt.Println(!male == false)
	fmt.Println(male != false)
}
