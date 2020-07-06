/**
* @Author : henry
* @Data: 2020-07-03 16:39
* @Note: goroutine泄漏
**/

package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// 做些有趣的操作
				fmt.Println(s)
			}
		}()
		return completed
	}
	doWork(nil)
	fmt.Println("Done.")
}
