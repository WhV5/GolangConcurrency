/**
* @Author : henry
* @Data: 2020-07-02 22:09
* @Note: sync.Once deadlock
**/

package main

import "sync"

func main() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }
	onceA.Do(initA)
}
