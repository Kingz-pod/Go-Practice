package type_data

import (
	"fmt"
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
