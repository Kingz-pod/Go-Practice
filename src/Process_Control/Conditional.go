package Process_Control

import "fmt"

//最经典的if else
func If_Else() {
	score := 50
	if score >= 90 {
		fmt.Println("成绩优秀")
	} else if score >= 80 { //这里 { 必须和else if同一行，不能写到下一行
		fmt.Println("成绩良好")
	} else {
		fmt.Println("成绩一般")
	}
}

//switch case语句，不需要break，进入case会直接退出
func Switch_Case() {
	subject := "英语"
	switch subject {
	case "语文":
		fmt.Println("这是语文课")
	case "数学":
		fmt.Println("这是数学课")
	case "英语":
		fmt.Println("这是英语课")
		fallthrough //switch的穿透，该语句会无条件再次执行下一个的case，且只会执行下一个case（无论之后条件是否符合）
	case "政治":
		fmt.Println("这是政治课")
	default:
		fmt.Println("反正是课，但是没找到你这个课")
	}

}
