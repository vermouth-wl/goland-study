package main

import (
	"fmt"
	"sync"
	"time"
)

// 函数的声明
func sum(a int, b int) int {
	return a + b
}

// 可变参数函数声明
func sumAny(args ...int) int {
	var sum int
	for _, v := range args {
		sum = sum + v
	}
	return sum
}

// 多个参数类型不一致的函数声明
func myPrintf(args ...interface{}) {
	for _, v := range args {
		switch v.(type) {
		case int:
			fmt.Println(v, "是int类型")
		case string:
			fmt.Println(v, "是string类型")
		case bool:
			fmt.Println(v, "是bool类型")
		default:
			fmt.Println(v, "是未知类型")
		}
	}
}

// 多个可变参数函数传递参数
func Sum(args ...int) int {
	// 利用...来解序列
	result := sumAny(args...)
	return result
}

// go协程：goroutine
// 协程的初步使用
func myTest() {
	fmt.Println("你好, jessica")
}

// 多个协程效果
func myGo(name string) {

	for i := 0; i < 10; i++ {
		fmt.Printf("In goroutine: %s\n", name)
		// 为了避免协程执行过快，添加睡眠
		time.Sleep(10 * time.Millisecond)
	}
}

// 遍历信道
func fibonacci(myChan chan int) {
	n := cap(myChan)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		myChan <- x
		x, y = y, x+y
	}
	// close信道，否则程序会阻塞
	close(myChan)
}

// 利用信道做锁
// 由于x = x + 1不是原子操作，所以应避免多个协程对x进行操作，使用容量为1的信道可以达到锁的效果
func increment(ch chan bool, x *int) {
	ch <- true
	*x = *x + 1
	<-ch
}

// 使用WaitGroup替代time.Sleep()
func worker(x int, wg *sync.WaitGroup) {
	// 子协程完成后，计数器减1
	defer wg.Done()
	for i := 0; i < 6; i++ {
		fmt.Printf("woker %d: %d\n", x, i)
	}
}

// 互斥锁与读写锁
// 不加互斥锁情况
func add1(count *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		*count = *count + 1
	}
}

// 加互斥锁情况
func add2(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		// 加锁
		lock.Lock()
		*count = *count + 1
		// 解锁
		lock.Unlock()
	}
}
