/**
* @Author : henry
* @Data: 2020-07-02 23:11
* @Note: goroutine
**/

package main

import "fmt"

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()

	fmt.Println(<-stringStream)
}
