package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("这是主包")

	fmt.Println("===============================")

	// 调用函数
	fmt.Println("调用函数的结果是: ", sum(1, 2))

	fmt.Println("===============================")

	// 函数实现可变参数
	fmt.Println("可变参数函数的结果是: ", sumAny(1, 2, 3, 4, 5))

	fmt.Println("===============================")

	// 多个类型不一致参数函数调用
	var v1 int = 10
	var v2 string = "测试"
	var v3 bool = false
	var v4 float64 = 99.99
	myPrintf(v1, v2, v3, v4)

	fmt.Println("===============================")

	// 多个可变参数函数传递参数, 解序列后
	fmt.Println("解序列后结果为: ", Sum(1, 2, 3))

	fmt.Println("===============================")

	// go协程：goroutine
	// 协程的初步使用
	go myTest()
	fmt.Println("你好, mebius")
	time.Sleep(time.Second)

	fmt.Println("===============================")

	// 多个协程执行效果
	go myGo("协程1号")
	go myGo("协程2号")
	go myGo("协程3号")
	go myGo("协程4号")
	time.Sleep(time.Second)

	fmt.Println("===============================")

	// 信道
	// 信道定义及发送数据
	pipline := make(chan int, 10)
	fmt.Printf("信道中可缓冲 %d 个数据\n", cap(pipline))
	// 往信道中发送数据
	pipline <- 3
	pipline <- 1
	fmt.Printf("信道中当前有 %d 个数据\n", len(pipline))

	fmt.Println("===============================")

	// 双向信道
	// 定义双向信道
	pipline2 := make(chan string, 10)
	go func() {
		fmt.Println("准备发送数据")
		pipline2 <- "云韵"
		fmt.Printf("信道的容量为 %d, 长度为 %d\n", cap(pipline2), len(pipline2))
	}()

	// 接收发送出去的数据
	go func() {
		name := <-pipline2
		fmt.Printf("接收到的数据是 %s\n", name)
	}()
	// 主函数Sleep， 使得上面两个goroutine有机会同时运行
	time.Sleep(time.Second)

	fmt.Println("===============================")

	// 单向信道
	// 定义只写信道类型
	type Sender = chan<- int
	// 定义只读信道类型
	type Receiver = <-chan int

	var pipline3 = make(chan int)
	go func() {
		var sender Sender = pipline3
		fmt.Println("准备发送数据")
		fmt.Printf("信道的容量是 %d\n", cap(sender))
		sender <- 66
	}()
	go func() {
		var reverser Receiver = pipline3
		num := <-reverser
		fmt.Printf("接收到的数据是 %d\n", num)
	}()
	// 主函数Sleep，使得上面两个goroutine有机会同时运行
	time.Sleep(time.Second)

	fmt.Println("===============================")

	// 信道遍历
	pipline4 := make(chan int, 10)
	go fibonacci(pipline4)

	for k := range pipline4 {
		fmt.Print(k, " ")
	}
	fmt.Println()

	fmt.Println("===============================")

	// 利用信道来做锁
	// 设置容量为1的缓冲信道
	pipline5 := make(chan bool, 1)

	var x int
	for i := 0; i < 1000; i++ {
		go increment(pipline5, &x)
	}
	// 确保所有协程都已经完成
	// 主函数Sleep
	time.Sleep(time.Second)
	fmt.Println("x 的值: ", x)

	fmt.Println("===============================")

	// WaitGroup
	// 使用信道来标记完成，定义信道，在任务完成后往信道写入true，主协程获取到true就认为子协程已经执行完毕
	done := make(chan bool)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
		done <- true
	}()
	<-done

	fmt.Println("===============================")

	// 使用sync包提供的WaitGroup类型
	var wg sync.WaitGroup
	// 传入子协程的数量
	wg.Add(3)
	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)

	// 阻塞当前协程，直到实例里面的计数器归零
	wg.Wait()

	fmt.Println("===============================")

	// 互斥锁与读写锁
	// 互斥锁不加锁情况
	var wg1 sync.WaitGroup
	count := 0
	wg1.Add(3)
	go add1(&count, &wg1)
	go add1(&count, &wg1)
	go add1(&count, &wg1)

	wg1.Wait()
	fmt.Printf("在不加锁的情况下最终计算的count的值为: %d\n", count)

	// 加互斥锁情况
	var wg2 sync.WaitGroup
	lock := &sync.Mutex{}
	count2 := 0
	wg2.Add(3)
	// 传入锁
	go add2(&count2, &wg2, lock)
	go add2(&count2, &wg2, lock)
	go add2(&count2, &wg2, lock)

	wg2.Wait()
	fmt.Printf("在加锁的情况下最终计算的count的值为: %d\n", count2)

	fmt.Println("===============================")

	// 读写锁
	lock2 := &sync.RWMutex{}
	// 加写锁
	lock2.Lock()

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("第 %d 个协程准备开始...\n", i)
			// 加读锁
			lock2.RLock()
			fmt.Printf("第 %d 个协程获得读锁，sleep 1s 后释放锁...\n", i)
			time.Sleep(time.Second)
			lock2.RUnlock()
		}(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("准备释放写锁，读锁不再阻塞")
	// 写锁一旦释放，读锁就全部自由了
	lock2.Unlock()

	// 读锁全部释放完成后，获得死锁
	// 最后执行
	lock2.Lock()
	fmt.Println("程序退出...")
	lock2.Unlock()
}
