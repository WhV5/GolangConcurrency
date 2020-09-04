/**
* @Author : henry
* @Data: 2020-07-06 20:26
* @Note:
**/

package main

import "testing"

//func TestDoWork_GeneratesAllNumbers(t *testing.T) {
//	done := make(chan interface{})
//	defer close(done)
//
//	intSlice := []int{0, 1, 2, 3, 5}
//	_, results := DoWork(done, intSlice...)
//
//	for i, expected := range intSlice {
//		select {
//		case r := <-results:
//			if r != expected {
//				t.Errorf(
//					"index %v: expected %v,but received %v,",
//					i,
//					expected,
//					r,
//				)
//			}
//		case <-time.After(1 * time.Second):
//			t.Fatal("test timed out")
//		}
//	}
//}

func TestDoWork_GeneratesAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	heartbeat, results := DoWork(done, intSlice...)

	<-heartbeat

	i := 0
	for r := range results {
		if expected := intSlice[i]; r != expected {
			t.Errorf("index %v:expected %v,but received %v,", i, expected, r)
		}
		i++
	}
}
