package type_data

import (
	"fmt"
)

//数字相关，整型和浮点
//uint表示无符号数，则上限会高点
func Int() (int, int, int) {
	//以下为整型10的不同进制表示法
	a := 10
	b := 0o12
	c := 0xA
	return a, b, c
}
func Float() (float64, float64) {
	//浮点数在Go里只有10进制
	//以下为科学计数法
	a := 2.2e+3
	b := 2.2e-3
	return a, b
}

//字符和字符串
func Char() {
	//总体来说，单个字符必须用单引号，双引号特指字符串，和python不同
	//首先是byte，单字节0-256的ASCII字符
	var a byte = 65                          //等同于\101 \x41  分别为八进制和十六进制
	var b uint = 'A'                         //uint实际上和byte性质相同
	fmt.Printf("a的值为: %c \nb的值为 %c\n", a, b) //用%c则输出该字节对应的ASCII可见字符

	//还有unicode字符，是4字节的，用go中独有的rune类型，占4个字节
	var c rune = 66 //unicode前面256和ascii是一样的
	var d rune = 'B'
	fmt.Printf("c的值为:%c \n b的值为:%c", c, d)
}
func String() {
	//双引号，本质是[]byte ，即byte数组
	str1 := "hello,世界"
	str2 := [12]byte{104, 101, 108, 108, 111, 44, 228, 184, 150, 231, 149, 140}
	fmt.Printf("str1为:%s\nstr2为:%s\n----------------\n", str1, str2)
	//输出结果一样，可以看出字符串本质，这里世界二字用的utf-8编码，每个字占3个字节，世：{228,184,150} 界：{231,149,140}

	//另外，反引号表示不转义字符串
	str3 := "准备换行:\n"
	str4 := `我想看到换行符:\n` // 这里\n原样输出了,相当于python里的 r''
	fmt.Println(str3, str4, "\n---------------")

	//反引号的好处,简单理解就是内容原样输出，包括回车
	str5 := `你好
这是第二行了
这是第三行了，我的一切输入都原样输出了`
	fmt.Println(str5)
}

//数组和切片
func Arr() {
	//数组的长度就是容量，定义的时候必须指定类型和长度
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4} //注意这里并不是说数组长度可变，而是让编译器自己判断数组是多长，这里是4
	fmt.Println(arr1, arr2)
	//总之，数组不能在后面加长，不能append
}

//获取当前切片的长度和容量
func PrintSlice(str string, s []int) {
	//len()计算长度，cap(s)计算容量
	fmt.Println(s)
	fmt.Printf("%s长度len=%d,容量cap=%d\n------------------------\n", str, len(s), cap(s))
}
func Slice() {
	//切片类似于python中的list，可以append加长,切片拥有 长度 和 容量。
	//切片的长度就是它所包含的元素个数。容量指的是从第一个元素开始数，到底层数组元素末尾的数量
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	PrintSlice("创建一个切片：", s)

	s = s[:0]
	PrintSlice("右指针指向0：", s) //此时虽然没有元素长度为0，但是容量还是8

	s = s[:4]
	PrintSlice("右指针指向4：", s) //容量依然是8，没有拓展

	s = s[2:4]
	PrintSlice("左指针指向2，右指针指向4：", s) //容量变小了，因为左边没了2个
	//总结以上，切片容量减少一般是左边变少，切片容量增加一般是右边拓展

	//接下来尝试append函数,实际原理是创建一个新的切片进行拷贝替换
	//容量是每次翻倍提升，，若超过1024则每次增加百分之25
	var s1 []int //创建一个空切片
	s1 = append(s1, 1)
	PrintSlice("给空切片末尾加上一个元素:", s1) //拓展了末尾，长度和容量都增加了

	s1 = append(s1, 2, 3, 4, 5)
	PrintSlice("继续连续增加4个元素", s1)

	s1 = append(s1, 6, 7, 8)
	PrintSlice("再在增加三个元素：", s1)

	//另外切片的第三个参数是规定容量的
	s2 := s1[:4:6]
	PrintSlice("切前四个元素，容量索引到第六个:", s2)

	//除了基本定义方法，还可以使用make函数定义，类型、长度、容量
	s3 := make([]int, 4, 8)
	PrintSlice("用make定义个一个切片：", s3)
}

//字典map与布尔类型
//字典是hash表的实现，key必须是唯一的，不可变的
func Map() {
	//三种声明方法,均是map[key_type]value_type
	//第一种var
	var map_1 map[string]string = map[string]string{"name": "钢铁侠"}
	//第二种 短赋值（建议）
	map_2 := map[string]string{"name": "猛男侠", "password": "123456"}
	//三种make
	map_3 := make(map[string]string) //自动初始化为空
	map_3["name"] = "青峰侠"
	fmt.Println(map_1, map_2, map_3)
	fmt.Println("---------------------")

	//删除一个key
	delete(map_2, "name")
	fmt.Println(map_2)
	fmt.Println("---------------------")

	//由于访问一个不存在的key不会报错，会返回对应value类型的0值，所以有时候需要判断key是否存在
	//我们可以通过第二个临时参数来确认，这个参数会返回true和false
	name, flag := map_2["name"]
	if flag {
		fmt.Println("name这个key存在:" + name)
	} else {
		fmt.Println("name不存在")
	}
}

//接下来是字典的循环，由于没有python里的keys()和values()函数，只能自己循环取值
func Map_cycle() {
	map_1 := map[string]string{"name": "钢铁侠", "password": "admin123"}
	//第一种，获取key和value
	for name, value := range map_1 {
		fmt.Printf("key:%s,value:%s\n", name, value)
	}
	fmt.Println("---------------------")
	//第二种，只获取key,不需要占位符
	for name := range map_1 {
		fmt.Printf("key:%s\n", name)
	}
	fmt.Println("---------------------")
	//第三种，只获取value，此时需要占位符
	for _, value := range map_1 {
		fmt.Printf("value:%s\n", value)
	}
}

//关于Bool类型，Go中是以小写的true和false为标准，和python不同，true和1不同。GO中类型不同的数值不能比较
func Bool_Type() {
	age := 15
	gender := "male"
	fmt.Println(age > 18 && gender == "male")
	fmt.Println(age < 18 || gender == "male") //左边已经确认，则这里右边的表达式直接不执行
}

//关于指针类型
func Point() {
	//先看三种定义
	//第一种，先创建一个指针来分配内存，然后填入值。使用new()函数
	ptr := new(string)
	*ptr = "test"

	//第二种,先创建一个变量，然后获取该变量的地址
	a := "test2"
	ptr1 := &a
	println(ptr1)
	//第三种 先声明一个指针变量和普通变量，然后赋值
	var ptr2 *string
	b := "test3"
	ptr2 = &b
	println(ptr2)
}

//关于指针和切片，指针和切片都是引用类型。go中传参除了引用类型均是拷贝。
//比如想通过一个函数改变数组的值时，Go中常见做法是传入切片（或者数组指针），而不是传入数组
func Point_Slice() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	arr2 := Modify_arr_2(arr)
	fmt.Println(arr2)
}

//直接传入数组的切片，由于切片是引用类型(地址),所以会直接修改原数组的值
func Modify_arr_1(sli []int) {
	sli[0] = 666
}

//传入数组的话，仅仅是拷贝，所以不会改变原数组的值，会产生一个新的数组,需要return
func Modify_arr_2(arr [6]int) [6]int {
	arr[0] = 666
	return arr
}
