/**
* @Author : henry
* @Data: 2020-07-03 10:46
* @Note: read data from close channel
**/

package main

import "fmt"

func main() {
	intStream := make(chan int)
	go func() {
		intStream <- 1
	}()
	close(intStream)
	integer, ok := <-intStream
	fmt.Printf("(%v):%v", ok, integer)
}
