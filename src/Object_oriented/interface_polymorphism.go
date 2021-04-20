package Object_oriented

import "fmt"

//接口，定义一个类型应该有的方法，具体的实现由该类型自行实现
//举例,定义一个接口（实际上也是定义了Animal这个类型）
type Animal interface {
	Eat() //必须会吃饭
}

//定义一个结构体：小猫
type Cat struct {
	Name string
	Age  int
}

//给结构体定义方法实现接口,完成eat()方法
func (c *Cat) Eat() {
	fmt.Printf("我是%s，在吃东西\n", c.Name)
}

/*
main中测试如下:
mycat := Object_oriented.Cat{Name:"小猫",Age:3}
mycat.Eat()  //输出：我是小猫，在吃东西
*/

//多态的实现，是利用接口实现。本质就是执行相同的函数，但结果表现不同
//接着上面的Animal接口，我们实现小狗和小猫
//小猫上面已经实现了，小猫有一个吃的方法
//接下来定义小狗
type Dog struct {
	Name string
	Age  int
}

//小狗实现接口的方法
func (d *Dog) Eat() {
	fmt.Printf("我是%s，在吃东西\n", d.Name)
}

//小狗还额外实现了一个狗叫方法
func (d *Dog) Wang() {
	fmt.Printf("我是%s，在狗叫\n", d.Name)
}

/*
func main() {
	mydog := Object_oriented.Dog{Name: "小狗",Age: 3}
	mycat := Object_oriented.Cat{Name: "小猫",Age: 3}
	mydog.Wang()
	//定义一个Animal类型的切片，由于用的指针实现的接口方法，所以这里必须使用结构体指针
	myanimal := []Object_oriented.Animal{&mydog,&mycat}  //显式赋值Animal类型
	//还可以隐性赋值比如: myanimal := &mydog  这样的话myanimal可以调用Wang()方法
	for _,animal := range myanimal{
		animal.Eat()  //由于这里显式定义了Animal类型，所以不能调用Wang()方法
	}
}
*/
