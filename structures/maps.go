package structures

import (
	"fmt"
	"maps"
)

func maps_example() {
	// 使用内置函数 'make(map[key-type]val-type)' 创建空 map。
	m := make(map[string]int)

	// 使用典型的 'name[key] = val' 语法来设置键值对。
	m["k1"] = 7
	m["k2"] = 13

	// 使用 fmt.Println() 打印一个 map，会输出它所有的键值对。
	fmt.Println("map:", m)

	// 使用 'name[key]' 来获取一个键的值。
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// 内置函数 len() 返回一个 map 的键值对数量。
	// 内置函数 delete() 可以从一个 map 中移除键值对。
	// 内置函数 clear() 可以删除所有键值对。
	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	clear(m)
	fmt.Println("map:", m)

	// 当从一个 map 中取值时，可以选择是否接收第二个返回值，该值表明该 map 中是否存在这个键。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 也可以在一行代码中声明并初始化一个新的 map。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// maps 包包含许多有用的 map 工具函数。
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
