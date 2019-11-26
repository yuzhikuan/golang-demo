/*
 * goroutine/channel/select
 * Channel 类型属于引用类型，它的零值即为nil
 * 1、对channel的重复关闭会引发运行时恐慌
 * 2、向一个已被关闭的channel发送数据会引发运行时恐慌
 * 3、发送操作会在channel已满时被阻塞
 * 4、接收操作会在channel已空时被阻塞
 *
 * 接收通道: type Receiver <-chan int
 * 发送通道: type Sender chan<- int
 */

package base

import (
	"fmt"
	"sync"
)

/* TestConcurrency 测试函数 goroutine
 * channel 先存后取，无缓存的会阻塞，直到有取到值，而有缓存的不会阻塞
 * channel 先取后存，有无缓存都会阻塞，直到存入值
 */
func TestConcurrency() {
	// c1 := make(chan bool) // 创建无缓存的channel
	// go func() {
	// 	fmt.Println("GO GO GO!!!")
	// 	<-c1 // 取出值
	// }()
	// c1 <- true // 阻塞，向channel存值，等待被读出

	// c2 := make(chan bool, 1) // 创建有缓存的channel
	// go func() {
	// 	fmt.Println("BO BO BO!!!")
	// 	<-c2
	// }()
	// c2 <- true // 不阻塞，直接退出

	// 打印CPU的核心数量
	// fmt.Println("CPUs:", runtime.NumCPU())
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// 启动10个goroutine算加法1 -- 缓存channel
	// c3 := make(chan bool, 10)
	// for i := 0; i < 10; i++ {
	// 	go Go(c3, i)
	// }
	// for i := 0; i < 10; i++ {
	// 	<-c3
	// }

	// 启动10个goroutine算加法2 -- waitgroup
	// wg := sync.WaitGroup{}
	// wg.Add(10)
	// for i := 0; i < 10; i++ {
	// 	go Go2(&wg, i)
	// }
	// wg.Wait()

	// 测试select
	c4, c5 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c4:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c4", v)
			case v, ok := <-c5:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c5", v)
			}
		}
	}()

	c4 <- 1
	c5 <- "zkyucn"
	c4 <- 10
	c5 <- "hello"

	close(c4)

	<-o
}

// Go 算一亿的加法
func Go(c chan bool, index int) {
	sum := 0
	for i := 1; i <= 100000000; i++ {
		sum += i
	}
	fmt.Println(index, sum)

	c <- true
}

// Go2 算一亿的加法
func Go2(wg *sync.WaitGroup, index int) {
	sum := 0
	for i := 1; i <= 100000000; i++ {
		sum += i
	}
	fmt.Println(index, sum)

	wg.Done()
}
