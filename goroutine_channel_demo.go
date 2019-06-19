package main

import (
	"fmt"
	"time"
)

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}

func all()  {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	// 如果不等待，上面两个还没有执行完就结束了
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

/**
channel
 */
//var ch1 chan int
//ch1 = make(chan int)

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

func queueModel()  {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

/**
 TODO select
 */
var ch1 = make(chan int32, 1000)
var ch2 = make(chan int32, 1000)

func ch1Produce()  {
	for i :=int32(0); i < 100; i++ {
		ch1 <- i
	}
}

func ch2Produce()  {
	for i := int32(1000); i < 2000; i++ {
		ch2 <- i
	}
}

// 这个地方需要一个循环
func consumerChAll()  {
	for {
		select {
		case v := <-ch1:
			fmt.Println(v)

		case v := <-ch2:
			fmt.Println(v)
		}
	}

}


/**
 TODO 分发任务这种避免共有参数

 */

func work()  {
	pendingChan := make(chan int, 1000)
	doneChan := make(chan int, 1000)
	go submitTask(pendingChan)
	go execTask(pendingChan, doneChan)
	go finishTask(doneChan)
}

func submitTask(pendingChan chan int)  {
	for i := range [...]int{1,2,3} {
		pendingChan <- i
	}
}

func execTask(pendingChan, doneChan chan int)  {
	for {
		value := <- pendingChan
		fmt.Println("execTask ", value)
		doneChan <- value
	}
}

func finishTask(doneChan chan int)  {
	for{
		value := <- doneChan
		fmt.Println("finishTask ", value)
	}
}

/**
TODO 懒性生成器，这个感觉是利用了类似消费队列阻塞的特性来实现的，for 不断给里面放数据，但每次取只取一个
 */

// TODO 这个变量如果不像放在这里
var incrCount int

func generator() chan int{

	vchan := make(chan int)


	go func() {

		for {
			vchan <- incrCount
			incrCount++
		}
	}()

	return vchan
}

func runGenerator()  {
	fmt.Println(<- generator())
	fmt.Println(<- generator())
	fmt.Println(<- generator())
}


func main() {
	//go ch1Produce()
	//go ch2Produce()
	//go consumerChAll()


	//work()
	runGenerator()
	time.Sleep(1e9)
}
