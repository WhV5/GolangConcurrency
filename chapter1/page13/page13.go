/**
* @Author : henry
* @Data: 2020-07-07 16:44
* @Note:
**/

package main

import "fmt"

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
