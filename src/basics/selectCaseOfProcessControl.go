package main

import (
	"fmt"
	"time"
)

/* 流程控制：select-case */

// select 超时处理函数
func makeTimeout(ch chan bool, i int) {
	time.Sleep(time.Second * time.Duration(i))
	ch <- true
}
func main() {

	// 简单实例
	c1 := make(chan string, 1)
	c2 := make(chan string, 2)

	c2 <- "hello"

	select {
	case msg1 := <-c1:
		fmt.Println("c1接收: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2接收: ", msg2)
	default:
		fmt.Println("无数据接收")
	}

	// 模拟无接收数据，无default情况
	c3 := make(chan string, 1)
	c4 := make(chan string, 1)

	select {
	case msg3 := <-c3:
		fmt.Println("接收数据: ", msg3)
	case msg4 := <-c4:
		fmt.Println("接收数据: ", msg4)
	// 解决死锁方法1：任何情况下，尽量写default
	default:

	}

	// 解决死锁方法2：让其中一个信道可以接收到数据
	c5 := make(chan string, 1)
	c6 := make(chan string, 1)

	// 开启一个协程，让数据可以发送到信道
	go func() {
		// time.Sleep(time.Second * 1)
		c5 <- "hello"
	}()
	go func() {
		c6 <- "world"
	}()

	select {
	case msg5 := <-c5:
		fmt.Println(msg5)
	case msg6 := <-c6:
		fmt.Println(msg6)
	}

	// 执行顺序随机性
	c7 := make(chan string, 1)
	c8 := make(chan string, 1)

	c7 <- "你好"
	c8 <- "世界"

	select {
	case msg7 := <-c7:
		fmt.Println("接收数据: ", msg7)
	case msg8 := <-c8:
		fmt.Println("接收数据: ", msg8)
	default:

	}

	// select 超时
	c9 := make(chan string, 1)
	c10 := make(chan string, 1)
	timeout := make(chan bool, 1)

	// c9 <- "c9接收到了数据"

	go makeTimeout(timeout, 2)

	select {
	case msg9 := <-c9:
		fmt.Println("c9接收数据: ", msg9)
	case msg10 := <-c10:
		fmt.Println("c10接收数据: ", msg10)
	case <-timeout:
		fmt.Println("接收超时, 退出")
	}

}
