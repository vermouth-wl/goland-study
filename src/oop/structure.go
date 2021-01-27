package main

import "fmt"

/* 结构体 */

// 定义结构体
type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}

// 定义与Profile结构体绑定的函数
func (person Profile) FmtProfile() {
	fmt.Printf("姓名: %s\n", person.name)
	fmt.Printf("年龄: %d\n", person.age)
	fmt.Printf("性别: %s\n", person.gender)
}

// 定义一个User结构体
type User struct {
	username string
	password string
	age      int
	sex      string
	address  string
	email    string
	status   int
}

// 定义与User结构体绑定的函数（如果需要在方法体内改变结构体内容时，需要以指针作为方法接受者）
func (user *User) addAge() {
	user.age += 1
}

// 结构体中实现继承
// 创建公司结构体
type company struct {
	companyName string // 公司名称
	companyAddr string // 公司地址
}

// 创建职员结构体
type staff struct {
	name     string // 姓名
	age      int    // 年龄
	gender   string // 性别
	position string // 位置
	company         // 匿名字段
}

// main()方法
func main() {

	// 实例化Profile对象
	myself := Profile{name: "王磊", age: 16, gender: "男性"}
	// 调用函数
	myself.FmtProfile()

	// 实例化User对象，并调用addAge()函数
	user := User{username: "王磊", password: "123456", age: 16, sex: "男", address: "重庆市"}
	fmt.Printf("%s 的年龄是: %d\n", user.username, user.age)
	user.addAge()
	fmt.Printf("%s 的年龄是: %d\n", user.username, user.age)

	// 实例化公司实例和人员实例
	myCom := company{
		companyName: "重庆农村商业银行",
		companyAddr: "重庆市南川区南大街47号",
	}
	staffInfo := staff{
		name:     "王磊",
		age:      16,
		gender:   "男",
		position: "重庆市南川区",
		company:  myCom,
	}
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyAddr)
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyAddr)
}
