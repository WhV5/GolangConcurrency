/**
* @Author : henry
* @Data: 2020-07-03 18:24
* @Note: pipeline
**/

package main

import "fmt"

func main() {
	// 将数据输入，对其进行转换并将数据发回
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	// 批处理
	intS1 := []int{1, 2, 3, 4}
	for _, v := range multiply(add(multiply(intS1, 2), 1), 2) {
		fmt.Println(v)
	}

	// 流处理
	intS2 := []int{1, 2, 3, 4}
	for _, v := range intS2 {
		fmt.Println(2 * (v*2 + 1))
	}
}
