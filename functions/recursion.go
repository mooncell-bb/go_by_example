package functions

import (
	"fmt"
)

// fact() 函数在到达 fact(0) 前一直调用自身。
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func recursion_example() {
	fmt.Println(fact(7))

	// 闭包也可以递归，但要求在定义闭包之前用类型化的 'var' 显式声明闭包。
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
