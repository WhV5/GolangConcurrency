/**
* @Author : henry
* @Data: 2020-07-02 21:17
* @Note: Cond Signal
**/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Remove from queue . len=", len(queue))
		c.L.Unlock()
		c.Signal() //让一个正在等待的goroutine知道发生了什么事
	}

	for i := 0; i < 10; i++ {
		c.L.Lock() //通过在条件的锁存器上调用锁来进入临界区
		for len(queue) == 2 {
			c.Wait() //调用Wait，这将暂停main goroutine直到要给信号的条件已经发送
		}
		fmt.Println("Adding to queue . len = ", len(queue))
		queue = append(queue, struct {
		}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
