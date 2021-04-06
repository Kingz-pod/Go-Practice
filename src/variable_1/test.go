package variable_1

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

//第五种，声明指针变量
