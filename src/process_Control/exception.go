package process_Control

import "fmt"

//异常机制

//1、手动触发宕机:panic("xxx")  类似于python中的exit()
func Panic() {
	panic("在此退出") //运行就会报错，提示在此退出
}

//2、捕获异常，确保主程序继续执行，或者做收尾工作
//通过recover()函数捕获异常，必须在defer中才能生效

func Recover(count int) {
	defer func() {
		//如果有异常就打印异常
		if err := recover(); err != nil {
			//这里可以写Log之类的，做收尾工作
			fmt.Println(err)
			//至此该携程会退出，但main仍然会继续执行
		}
	}()
	arr := [10]int{}
	arr[count] = 111        //创造一个数组越界，此时会进入defer，随后携程因报错退出
	fmt.Println("我上面好像越界了") //若数组越界，则这里显然不会执行了
}
