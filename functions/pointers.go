package functions

import "fmt"

// 函数会得到实参的拷贝。
func zeroval(ival int) {
	ival = 0
}

// 函数使用了 int 指针，*iptr 会解引用这个指针，读取对应的值。
func zeroptr(iptr *int) {
	*iptr = 0
}

func pointers_example() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	// 通过 '&i' 语法来取得 i 的内存地址，即指向 i 的指针。
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)
}
