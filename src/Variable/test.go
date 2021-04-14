package Variable

/*
说在前面：方法名大写表示Public，其他包可以调用，小写则只能本包调用。结构体中的变量也是如此。
*/

//第一种var 关键字定义遍历
func Var_variable() (string, string, string) {
	var a = "test"        //编译器自动识别成string
	var b string = "test" //指定string
	var c string          //自动初始化，不需要设置为NULL，但必须指定类型
	return a, b, c
}

//第二种，var的复合定义
func Var_variable_complex() (string, string) {
	//复合定义
	var (
		a        = "test"
		b string = "test2"
	)
	return a, b
}

//第三种,常用的短声明法
func Varable_3() int {
	//该方法只能用于函数内
	a := 1123 //必须赋值，自动识别类型
	return a
}

//第四种，多变量同时赋值
func Variable_4() (int, int) {
	a, b := 124, 444
	return a, b
}

//第五种，声明指针
func Variable_5() *int {
	p := new(int) //声明一个指针p
	*p = 123      //给地址赋值

	//当然你也可以先对a赋值，然后直接p=&a 这样也可以定义一个指针变量
	return p
}

//匿名变量，可以不用到的变量，占位
func Variable_6() int {
	a, _ := 123, 123 //下划线占位
	return a
}
