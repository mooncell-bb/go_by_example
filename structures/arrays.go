package structures

import "fmt"

func arrays_example() {
	// 数组 a 包含 5 个 int。
	// 定义数组时，需指定元素的类型和长度，且数组默认值是零值。
	var a [5]int
	fmt.Println("emp: ", a)

	// 使用 array[index] = value 在索引处设置值。
	// 使用 array[index] 获取值。
	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	// 内置函数 len() 返回数组的长度。
	fmt.Println("len: ", len(a))

	// 声明并初始化数组。
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// 使用 '...' 让编译器自动统计元素数量。
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// 使用 'index: value' 指定索引，中间的元素将被自动初始化为零值。
	// 只有类型和长度都相同时，才能进行赋值。
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx: ", b)

	// 数组类型是一维的，但可以通过组合类型来构建多维数据结构。
	// 使用 fmt.Println() 函数打印数组时，数组会以 [v1 v2 v3 ...] 的形式显示。
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
