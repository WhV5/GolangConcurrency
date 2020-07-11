/**
* @Author : henry
* @Data: 2020-07-04 21:47
* @Note:
**/

package main

import "time"

type Context interface {
	//当为该context工作的work被取消时，deadline将返回时间，在没有设定期限的情况下，
	//deadline返回 ok==false,连续的调用deadline返回相同的结果
	Deadline() (deadline time.Time, ok bool)

	//当为该context工作的work被取消时，返回一个关闭的channel，如果这个context
	//不能被取消，那么Done可能返回nil。连续调用完成返回相同的值
	Done() <-chan struct{}

	//Err在完成后返回一个non-nil值，如果context被取消，或者在context
	//的deadline结束时，如果context被取消，Err将被取消。没有定义Err的其它值。
	//连续调用结束后，用Err返回相同的值
	Err() error

	//值返回与此context关联的key或者nil，如果没有与键关联的值，则返回值为nil。
	//连续调用具有相同key的值将返回相同的结果
	Value(key interface{}) interface{}
}
