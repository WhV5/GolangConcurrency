/**
* @Author : henry
* @Data: 2020-07-03 17:56
* @Note:
**/

package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{
					Error:    err,
					Response: resp,
				}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}
	done := make(chan interface{})
	defer close(done)
	errCount := 0

	urls := []string{"a", "https://www.baidu.com", "b", "c", "d"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many errors,breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response:%v\n", result.Response.Status)
	}
}
