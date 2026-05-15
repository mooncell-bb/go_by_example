package functions

import "fmt"

// 定义一个函数，接受两个 int 类型的参数，并以 int 类型返回结果。
func plus(a int, b int) int {
	return a + b
}

// 当多个连续的参数类型相同时，可以仅声明最后一个参数的类型。
func plusPlus(a, b, c int) int {
	return a + b + c
}

func basics_example() {
	// 使用 "name(args)" 来调用一个函数。
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res2 := plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res2)
}
