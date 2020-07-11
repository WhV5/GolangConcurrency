/**
* @Author : henry
* @Data: 2020-07-03 15:44
* @Note:
**/

package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:

		}
		//模拟工作行为
		workCounter++
		time.Sleep(time.Second)
	}
	fmt.Printf("Achivevd %v cycles of work before signalled to stop.\n", workCounter)
}
