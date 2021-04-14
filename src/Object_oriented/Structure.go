package Object_oriented

import "fmt"

//结构体，存储多个变量类型(不存方法)。Go中没有class的概念，只有struct的概念。所以也没有继承
/*
`````基本结构如下：```````
type 结构体名 struct {
    属性名   属性类型
    属性名   属性类型
    ...
}
*/

//举例定义一个人的结构体
//注意！！！：首字母大写的才可以被其他包调用，小写只能本包调用.相当于private和public
type Person struct {
	Name   string
	Age    int
	Gender string
}

//那么需要给结构体定义方法怎么办，可以用`组合函数`的办法
//p表示Person的实例化,Eat为方法名，此时Eat和Person结构体就绑定了
//另外：这里使用指针是因为Go默认是进行数据拷贝，若想在方法内修改值，则需要使用指针
func (p *Person) Eat() {
	//建议通用指针，无论修不修改
	fmt.Printf("%s今年%d岁，现在正在吃饭。", p.Name, p.Age)
}

//前面说了，struct只有变量，所以没有继承。原因就是推荐使用组合代替继承
//组合就是把一个结构体嵌入到另外一个结构体
type School struct {
	//先定义一个学校结构体
	School_name    string
	School_address string
}
type Class struct {
	//再定义一个班级结构体,需要继承学校
	Class_name string
	School     //匿名字段，实现组合
}

/*
main函数中测试如下:
func main() {
	myschool := Object_oriented.School{School_name: "超级学院",School_address: "天堂国1号"}
	myclass := Object_oriented.Class{Class_name: "七班",School:myschool}
	fmt.Println(myclass.School_address)
	//输出天堂国1号
}
*/
