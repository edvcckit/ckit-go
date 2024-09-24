package ckit

// AnyValue
// @Description: 结构定义, 用于类型转换
// @Author: Edv Chen <edvcc72@gmail.com>
type AnyValue struct {
	Val any
	Err error
}

var test = AnyValue{}
var test2 = AnyValue{}
