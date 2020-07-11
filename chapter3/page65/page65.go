/**
* @Author : henry
* @Data: 2020-07-02 21:04
* @Note: cond 好像被取消了，待查证
**/

package main

import "sync"

func main() {
	c := sync.NewCond(&sync.Mutex{})
	c.L.Lock()
	// Wait atomically unlocks c.L and suspends execution
	// of the calling goroutine. After later resuming execution,
	// Wait locks c.L before returning. Unlike in other systems,
	// Wait cannot return unless awoken by Broadcast or Signal.
	//
	// Because c.L is not locked when Wait first resumes, the caller
	// typically cannot assume that the condition is true when
	// Wait returns. Instead, the caller should Wait in a loop:
	//
	//    c.L.Lock()
	//    for !condition() {
	//        c.Wait()
	//    }
	//    ... make use of condition ...
	//    c.L.Unlock()
	for condition() == false {
		c.Wait()
	}
	c.L.Unlock()
}
