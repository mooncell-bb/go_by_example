package basics

import "fmt"

func if_else_example() {
	// 最基本的例子。
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 可以不使用 else 而只用 if 语句。
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 通常在条件中使用 && 和 || 等逻辑运算符。
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// 在条件语句之前可以有一个声明语句，此声明变量可在条件分支中使用。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
