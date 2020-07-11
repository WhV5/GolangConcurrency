/**
* @Author : henry
* @Data: 2020-07-02 17:17
* @Note: sync.WaitGroup 闭包
**/

package main

import (
	"fmt"
	"sync"
)

func main() {
	//var wg sync.WaitGroup
	//salutation := "hello"
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	salutation = "welcome"  //修改变量salutation的值
	//}()
	//wg.Wait()
	//fmt.Println(salutation)

	//runtime.GOMAXPROCS(1)
	//var wg sync.WaitGroup
	//for _, salutation := range []string{"hello", "greetings", "good day"} {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		fmt.Println(salutation) // 引用了字符串类型的切片作为创建循环变量的salutation
	//	}()
	//}
	//wg.Wait()

	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greeting", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			//fmt.Println(salutation)
			fmt.Printf("%s %v\n", salutation, &salutation)
		}(salutation)
	}
	wg.Wait()
}
