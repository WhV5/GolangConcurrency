/**
* @Author : henry
* @Data: 2020-07-02 17:33
* @Note: goroutine size    my pc(windows) 8kb aliyun(linux) okb  demo 2.817kb
**/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } //创建一个永不退出的goroutine
	const numGoroutines = 1e4
	wg.Add(numGoroutines)   // 定义了goroutine的数量
	before := memConsumed() //测算在gouroutine之前消耗的内存总量
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed() //测算在goroutine之后消耗的内存总量
	fmt.Printf("%.3fkb", float64((after-before)/numGoroutines/1000))
}
