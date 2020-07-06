/**
* @Author : henry
* @Data: 2020-07-02 20:28
* @Note: waitGroup defer
**/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) //调用Add,参数为1，表示一个goroutine开始了
	go func() {
		defer wg.Done() // 使用defer关键字来确保goroutine退出之前执行done操作
		// 我们向WaitGroup表明我们已经退出了
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		//time.Sleep(2)
	}()

	//fmt.Println("All goroutines complete.")
	wg.Wait() //执行wait操作，这将阻塞 main goroutine，直到所有goroutine
	// 表明它们已经退出
	fmt.Println("All goroutines complete.")
}
