/**
* @Author : henry
* @Data: 2020-07-04 20:32
* @Note:
**/

package main

import (
	"fmt"
	"time"
)

func main() {
	repeat := func(
		done <-chan interface{},
		values ...int,
	) <-chan int {
		valueStream := make(chan int)
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:

					}
				}
			}
		}()
		return valueStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	sleep := func(
		done <-chan interface{},
		timeSleep time.Duration,
		value <-chan interface{}) <-chan interface{} {
		sleepStream := make(chan interface{})
		go func() {
			defer close(sleepStream)
			for {
				select {
				case <-done:
					return
				case sleepStream <- value:
					time.Sleep(timeSleep)
				}
			}
		}()
		return value
	}

	done := make(chan interface{})
	defer close(done)
	intStream := make(chan interface{})
	intStream <- 3
	zeros := take(done, intStream, <-repeat(done, 0))
	short := sleep(done, 1*time.Second, zeros)
	long := sleep(done, 4*time.Second, short)
	pipeline := long

	fmt.Println(pipeline)
}
