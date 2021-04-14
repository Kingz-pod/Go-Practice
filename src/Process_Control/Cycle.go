package Process_Control

import (
	"fmt"
)

//循环语句

//for 循环，可以有3种用法。
//特殊用法 : 直接for不接任何表达式，表示无限循环

//第一种，只接一个条件表达式
func For_Cycle_1() {
	a := 1
	for a < 5 {
		fmt.Printf("a的值是%d\n", a)
		a++
	}
}

//第二种接三个表达式的，类似C
func For_Cycle_2() {
	for i := 0; i < 5; i++ {
		fmt.Printf("a的值是%d\n", i)
	}
}

//第三种，接range，range用于遍历一个可以迭代的对象，比如数组、切片、字符串
func For_Cycle_3() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	for index, value := range arr { //如果不需要索引，则用_代替占位
		fmt.Printf("第%d个值为%d\n", index, value)
	}
	fmt.Printf("----------------------\n")
	str := "password!!!"
	for _, value := range str {
		fmt.Printf("值为%c\n", value)
	}
}
