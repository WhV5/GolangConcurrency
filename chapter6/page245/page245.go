/**
* @Author : henry
* @Data: 2020-07-07 16:23
* @Note:
**/

package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	log.SetFlags(log.Ltime | log.LUTC)
	log.SetOutput(os.Stdout)

	// 每1s log都会记录有多少个goroutine在并发执行
	go func() {
		goroutines := pprof.Lookup("goroutine")
		for range time.Tick(1 * time.Second) {
			log.Printf("goroutine count: %d\n", goroutines.Count())
		}
	}()

	// 创建一些永远不会退出的 goroutine
	var blockForever chan struct{}
	for i := 0; i < 10; i++ {
		go func() { <-blockForever }()
		time.Sleep(500 * time.Millisecond)
	}
}
