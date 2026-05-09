package structures

import (
	"fmt"
	"slices"
)

func slices_example() {
	// 切片仅由其包含的元素类型定义，而无需指定长度。
	// 未初始化的切片等于 nil，且长度为 0。
	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)

	// 使用内置的 make() 函数来创建非零长度的切片，初始值为零。
	// 如果预先知道切片会增长，可以显式地将容量作为额外参数传递给 make() 函数。
	s = make([]string, 3)
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	// 可以像数组一样设置和获取值。
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// 内置的 append() 函数会返回一个包含一个或多个新值的切片。
	// 由于 append() 返回一个新的切片，因此需要接收其返回值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// 切片还可以复制，使用内置的 copy() 函数。
	// 创建一个与 s 长度相同的空切片 c，并将内容复制到 c 中。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	// 切片支持通过 'slice[low:high]' 语法进行切片操作（包含 low，不包含 high）。
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 'slice[:high]' —— 包含从 s[0] 到 s[high]（不包含 high）的元素。
	// 'slice[low:]' —— 包含从 s[low]（包含 low）之后的元素。
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)

	// 在一行内声明并初始化一个切片变量。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	// slices 包中包含许多实用的切片工具函数。
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// 切片可以组合成多维数据结构。
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
