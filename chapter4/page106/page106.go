/**
* @Author : henry
* @Data: 2020-07-03 16:49
* @Note:  消除 goroutine泄漏
**/

package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{}, // 将完成的channel传递给doWork函数
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					//做一些有意思的操作
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		// 在1s之后取消本操作
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done")
}
