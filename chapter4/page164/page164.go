/**
* @Author : henry
* @Data: 2020-07-05 11:32
* @Note:
**/

package main

import "fmt"

func main() {
	type foo int
	type bar int
	m := make(map[interface{}]int)
	m[foo(1)] = 1
	m[bar(1)] = 2
	fmt.Printf("%v", m)

}
