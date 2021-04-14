package main

import (
	"fmt"
	"go/src/Object_oriented"
)

func main() {
	myschool := Object_oriented.School{School_name: "超级学院", School_address: "天堂国1号"}
	myclass := Object_oriented.Class{Class_name: "七班", School: myschool}
	fmt.Println(myclass.School_address)
}
