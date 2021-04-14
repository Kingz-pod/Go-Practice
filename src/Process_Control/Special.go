package Process_Control

import "fmt"

//goto 语句省略，就是直接跳转到指定的标签

//defer语句，当前函数执行完毕后再执行，甚至在return之后，常见场景如下
func Defer_Test() int {
	a := 1
	defer fmt.Println("测试结束") //不再需要在每个return都写上打印数据，常用于资源统一释放
	if a == 1 {
		//fmt.Println("测试结束")
		return a + 4
	}
	if a == 2 {
		//fmt.Println("测试结束")
		return a + 3
	}
	if a == 3 {
		//fmt.Println("测试结束")
		return a + 2
	}
	return -1
}
