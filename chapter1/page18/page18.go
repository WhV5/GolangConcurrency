/**
* @Author : henry
* @Data: 2020-07-08 18:07
* @Note:
**/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	//time.Sleep(time.Second)
	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock()
}
