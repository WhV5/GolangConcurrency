/**
* @Author : henry
* @Data: 2020-07-02 22:07
* @Note: sync.Once
**/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)
	fmt.Printf("Count:%d\n", count)
}
