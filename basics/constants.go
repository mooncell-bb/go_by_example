package basics

import (
	"fmt"
	"math"
)

// const 用于声明一个常量，Go 支持字符、字符串、布尔和数值常量。
const s string = "constant"

// 使用 iota 进行常量枚举。
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func constants_example() {
	fmt.Println(s)
	fmt.Println("weekday: ", Friday)

	// const 语句可以出现在任何 var 语句可以出现的地方。
	const n = 500000000

	// 常数表达式可以执行任意精度的运算。
	const d = 3e20 / n
	fmt.Println(d)

	// 数值型常量没有确定的类型，直到被给定某个类型。
	fmt.Println(int64(d))

	// 一个数字可以根据上下文的需要自动确定类型。
	fmt.Println(math.Sin(n))
}
