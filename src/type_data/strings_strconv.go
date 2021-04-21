package type_data

import (
	"fmt"
	"strconv"
	"strings"
)

//strings内置包，完成字符串的一些操作

//判断前缀/后缀:HasPrefix/HasSuffix
func Pre_Suf() {
	fmt.Println(strings.HasPrefix("Hello,word!", "Hello")) //判断是否以hello开头
	fmt.Println(strings.HasSuffix("okgogoggo!", "!"))      //判断是否以!结尾
}

//判断字符串是否包含 Contains
func Contains() {
	fmt.Println(strings.Contains("Hello,Word!", "Hello")) //判断是否包含Hello
}

//获取指定字符串出现的第一个位置的索引,Index
func Index() {
	fmt.Println(strings.Index("号,Word!", "号")) //获取号出现的第一个位置索引，-1表示不包含
}

//获取指定字符串出现的最后一个位置的索引,LastIndex
func LastIndex() {
	fmt.Println(strings.LastIndex("Hello,Word!!!", "!")) //获取最后一个出现的!的索引,-1表示不包含
}

/*
字符串替换 Replace(str,old,new,n) 返回string类型
具体是把str中的前n个old替换成new。如果n为-1则替换所有
*/
func Replace() {
	fmt.Println(strings.Replace("你好，世界，这里是Go的世界", "世界", "朋友", -1)) //替换全部的世界为朋友
	fmt.Println(strings.Replace("你好，世界，这里是Go的世界", "世界", "朋友", 1))  //只把前1个世界替换成朋友
}

//统计出现字符串的个数Count
func Count() {
	fmt.Println(strings.Count("Hello,World!!!!!", "!")) //统计!出现的次数
}

//重复字符串Repeat，相当于python里字符串乘法，"t1"*3="t1t1t1"。（Go很严格，不同类型不能做运算）
func Repeat() {
	fmt.Println(strings.Repeat("ok!", 10)) //输出10次ok!
}

//去除开头和结尾的空白符号TrimSpace
func TrimSpace() {
	fmt.Println(strings.TrimSpace("  确实!    "))
}

//去除开头结尾的指定字符Trim
func Trim() {
	fmt.Println(strings.Trim("http://www.baidu.com", "http://"))
}

//字符串分割 Split,返回一个slice
func Split() {
	s := strings.Split("GO,来吧,冲!", ",") //指定使用,来分割字符串，返回一个slice
	for _, value := range s {           //循环输出slice值
		fmt.Println(value)
	}
}

//拼接字符型slice为字符串,返回一个字符串:Join
func Join() {
	s := []string{"first", "second", "third"} //定义一个slice
	str := strings.Join(s, ",")               //用逗号拼接每一个slice元素
	fmt.Println(str)
}

//strconv包，用来字符串和其他类型的转换

//转换测试
func Strconv() {
	//int 转 string
	str := strconv.Itoa(123) //把int转换成对应的字符串
	fmt.Println("int的123成功转为字符串:" + str)

	//string 转成 int,返回2个参数，所以需要接收错误，或者用_占位也行
	i, err := strconv.Atoi("123456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

}
