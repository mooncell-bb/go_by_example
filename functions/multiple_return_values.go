package functions

import "fmt"

// 使用 '(int, int)' 的形式使函数反回两个 int 值。
func vals() (int, int) {
	return 3, 7
}

func multiple_return_values_example() {
	// 通过多赋值操作来使用这两个不同的返回值。
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 如果仅仅需要返回值的一部分，可以使用空白标识符 '_'。
	_, c := vals()
	fmt.Println(c)
}
