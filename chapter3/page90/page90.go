/**
* @Author : henry
* @Data: 2020-07-03 11:50
* @Note: Producer Consumer
**/

package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Receive: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
