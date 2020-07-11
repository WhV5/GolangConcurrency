/**
* @Author : henry
* @Data: 2020-07-03 16:24
* @Note: for-select 循环
**/

package main

func main() {

	//for { // 要不就无限循环，要不就使用range语句循环
	//	select {
	//	// 使用channel进行作业
	//	}
	//}

	//var done <-chan int
	//done = make(chan int)
	//var stringStream chan<- string
	//stringStream = make(chan string)

	//for _, s := range []string{"a", "b", "c"} {
	//	select {
	//	case <-done:
	//		return
	//	case stringStream <- s:
	//	}
	//}

	//// 创建循环，无限循环直到停止的goroutine很常见
	//// 方式一
	//for {
	//	select {
	//	case <-done:
	//		return
	//	default:
	//	}
	//	// 进行非抢占式任务
	//}
	//
	//// 方式二
	//for {
	//	select {
	//	case <-done:
	//		return
	//	default:
	//		// 进行非抢占式任务
	//	}
	//}

}
