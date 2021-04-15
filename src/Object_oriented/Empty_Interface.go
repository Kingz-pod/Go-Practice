package Object_oriented

import "fmt"

//空接口算是常用的东西之一了
//没有定义任何方法的接口就是空接口
//由于类型和方法都没有定义，所以都是nil

//常用于方法接受任意类型的参数
func Empty_Interface(obj interface{}) {
	fmt.Println(obj)
}

//还可以定义一种可以接收任意类型的数组、切片、map、struct
