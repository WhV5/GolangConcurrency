/**
* @Author : henry
* @Data: 2020-07-02 23:13
* @Note: channel deadlock
  //  可以有1个或2个返回值
**/

package main

import "fmt"

func main() {
	//stringStream := make(chan string)
	//go func() {
	//	if 0 != 1 {
	//		return
	//	}
	//	stringStream <- "hello channels!"
	//}()
	//fmt.Println(<-stringStream)

	stringStream := make(chan string)
	go func() {
		stringStream <- "hello channels!"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v):%v", ok, salutation)
}
