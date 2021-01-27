package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/* 面向对象：结构体里的tag用法 */

// 创建含Tag标签的结构体, 属性名需要大写，否则无法识别，json后接键名，omitempty在无该属性时不显示该键
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Addr  string `json:"addr,omitempty"`
	Email string `json:"email,omitempty"`
}

func main() {
	p1 := Person{
		Name:  "王磊",
		Age:   16,
		Email: "152",
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// 输出data1
	fmt.Printf("%s\n", data1)

	p2 := Person{
		Name: "王磊",
		Age:  16,
		Addr: "重庆",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data2)

	// 实例：美化输出
	user := User1{
		Username: "王磊",
		Password: "123456",
	}
	Print(user)
}

type User1 struct {
	Username string `label:"用户名是: "`
	Password string `label:"密 码是: "`
	Email    string `label:"邮箱地址是: " default:"未知"`
}

// Print()方法
func Print(obj interface{}) error {
	// 取到value
	v := reflect.ValueOf(obj)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取tag
		filed := v.Type().Field(i)
		tag := filed.Tag

		// 解析label和default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则用default替代
			value = defaultValue
		}
		fmt.Println(label + value)
	}
	return nil
}
