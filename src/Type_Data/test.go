package Type_Data

import "fmt"

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
func Slice() {
	//切片类似于python中的list，可
}
