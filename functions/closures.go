package functions

import "fmt"

// intSeq() 函数返回一个在其函数体内定义的匿名函数。
// 返回的函数使用闭包的方式隐藏变量 i。
func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func closures_example() {
	// nextInt 函数的值包含了自己的值 i，每次调用 nextInt() 时，都会更新 i 的值。
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 此状态对于特定的函数是唯一的。
	newInts := intSeq()
	fmt.Println(newInts())
}
