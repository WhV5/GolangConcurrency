/**
* @Author : henry
* @Data: 2020-07-03 17:01
* @Note:
**/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandStream := func(flag <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-flag:
					return
				}
			}
		}()
		return randStream
	}

	flag := make(chan interface{})
	randStream := newRandStream(flag)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d : %d\n", i, <-randStream)
	}
	close(flag)

	// 模拟正在进行的工作
	time.Sleep(1 * time.Second)
}
