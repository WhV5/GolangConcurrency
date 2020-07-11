/**
* @Author : henry
* @Data: 2020-07-03 10:51
* @Note: range channel
**/

package main

import "fmt"

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v", integer)
	}
}
