package basics

import "fmt"

func variables_example() {
	// var 声明 1 个或者多个变量。
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go 会自动推断已经有初始值的变量的类型。
	var d = true
	fmt.Println(d)

	// 声明后却没有给出对应的初始值时，变量将会初始化为零值。
	var e int
	fmt.Println(e)

	// := 语法是声明并初始化变量的简写。
	f := "short"
	fmt.Println(f)
}
