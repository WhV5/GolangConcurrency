/**
* @Author : henry
* @Data: 2020-07-03 16:05
* @Note:
**/

package main

import (
	"fmt"
)

func main() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5) //包含了这个goroutine的写入，防止其它goroutine写入
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // channel只读副本
		for results := range results {
			fmt.Printf("Received: %d\n", results)
		}
		fmt.Println("Done receiving")
	}
	results := chanOwner() // 收到 channel的读处理，消费者只能从中读取信息
	consumer(results)
}
