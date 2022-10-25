package goroutine

import (
	"sync"
	"time"
)

func TestGoroutine() {
	//testSleep()
	//TestWaitGroup()

	//testChannel()
	//testChannelClose()
	//testChannelSelect()

	testGetAndSet()

}
func testGetAndSet() {
	c := make(chan int)
	var read <-chan int = c
	var write chan<- int = c

	go setChan(write)

	getChan(read)
}

func getChan(read <-chan int) {
	for i := 0; i < 20; i++ {
		println("get", <-read)
	}
}

func setChan(write chan<- int) {
	for i := 0; i < 25; i++ {
		println("set", i)
		write <- i
	}
}

func testChannelSelect() {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c3 := make(chan int, 1)
	c4 := make(chan int, 1)

	c1 <- 1
	c2 <- 2
	c3 <- 3
	select {
	case <-c1:
		println("我是1", c1)
	case <-c2:
		println("我是2", c2)
	case <-c3:
		println("我是3", c3)
	case <-c4:
		println("我是4", c4)
	default:
		println("我是default")
	}
}
func testChannelClose() {
	c := make(chan int, 5)
	c <- 1
	c <- 2
	c <- 3
	//没有关闭不可以使用range函数，close后不可以write但是可以read
	close(c)
	for v := range c {
		println(v)
	}
}
func testChannel() {
	c := make(chan int, 5)
	go func() {
		for i := 0; i < 20; i++ {
			println("存储", i)
			c <- i
		}
	}()

	for i := 0; i < 20; i++ {
		println("取值", <-c)
	}
}

func Run(wg *sync.WaitGroup) {
	println("我开始跑了")
	wg.Done()
}

func TestWaitGroup() {
	//testSleep()
	wg := sync.WaitGroup{}
	wg.Add(1)

	go Run(&wg)
	wg.Wait()

	println("TestGoroutine 结束了")
}

func Run1() {
	println("我开始跑了")
}

func testSleep() {

	go Run1()

	time.Sleep(1 * time.Second)
	for i := 0; i < 30; i++ {
		println(i)
	}

	println("TestGoroutine 结束了")
}
