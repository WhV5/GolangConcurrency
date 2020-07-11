/**
* @Author : henry
* @Data: 2020-07-02 16:57
* @Note: go 触发goroutine
**/

package main

import "fmt"

func main() {
	//go sayHello()

	//go func() {
	//	fmt.Println("hello")
	//}()

	sayHello := func() {
		fmt.Println("hello")
	}
	go sayHello()
}

func sayHello() {
	fmt.Println("hello")
}
