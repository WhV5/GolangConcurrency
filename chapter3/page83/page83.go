/**
* @Author : henry
* @Data: 2020-07-03 11:10
* @Note: close channel to send signal to goroutines
**/

package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unbolcking goroutines...")
	close(begin)
	wg.Wait()
}
