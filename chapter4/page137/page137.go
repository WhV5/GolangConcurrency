/**
* @Author : henry
* @Data: 2020-07-03 23:55
* @Note:
**/

package main

import "fmt"

func main() {
	var myChan chan interface{}
	myChan = make(chan interface{})
	for val := range myChan {
		// 用val执行某些操作
		fmt.Println(val)
	}

	var done chan interface{}
	done = make(chan interface{})
loop:
	for {
		select {
		case <-done:
			break loop
		case maybeVal, ok := <-myChan:
			if ok == false {
				return //或许从for循环中退出
			}
			// 用val执行某些操作
			fmt.Println(maybeVal)
		}
	}

	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}
	for val := range orDone(done, myChan) {
		fmt.Println(val)
	}
}
