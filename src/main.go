package main

import (
	"go/src/Object_oriented"
)

func main() {
	mydog := Object_oriented.Dog{Name: "小狗", Age: 3}
	mycat := Object_oriented.Cat{Name: "小猫", Age: 3}
	mydog.Wang()
	//定义一个Animal类型的切片，由于用的指针实现的接口方法，所以这里必须使用结构体指针
	myanimal := []Object_oriented.Animal{&mydog, &mycat}
	myanimal[0].Eat()
	myanimal[0].Eat()
	myanimal[1].Eat()
}
