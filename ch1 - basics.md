# Basics

## Hello World

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

Go 程序由包组成，main 是一个特殊的包（一个程序有且仅有一个），定义了一个可独立运行的程序。

func main() 是入口函数，程序从这里开始执行。main() 函数不接受参数，也不返回值。

命令 `go run hello-world.go` 编译并执行程序，不产生二进制文件，适合开发阶段快速验证。

命令 `go build hello-world.go` 生成可执行文件，后序可直接运行此二进制文件。

命令 `go install` 会将编译后的二进制文件安装到 `$GOPATH/bin` 目录，方便随处调用。

## Values

```go
func values_example() {
	// strings
	fmt.Println("go" + "lang")

	// integers and floats
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	// booleans
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
```

- bool - true / false
- string - UTF-8 字符串，不可变。
- int / uint - 长度与平台相关：32 位系统下 32-bit，64 位下 64-bit。
  - int8 / int16 / int32 / int64 - 指定 int 位宽。
  - uint8 / uint16 / uint32 uint64 - 指定 uint 位宽。
- byte - 等价于 uint8，表示字节。
- rune - 等价于 int32，表示 Unicode 码点。
- float32 / float64 - 浮点数，符合 IEEE-754 标准。
- complex64 / complex128 - 虚数，实部 + 虚部。

## Variables

```go
func variables_example() {
	// var 声明 1 个或者多个变量。
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go 会自动推断已经有初始值的变量的类型。
	var d = true
	fmt.Println(d)

	// 声明后却没有给出对应的初始值时，变量将会初始化为零值。
	var e int
	fmt.Println(e)

	// := 语法是声明并初始化变量的简写。
	f := "short"
	fmt.Println(f)
}
```

Go 中所有变量声明后都有默认值：

- bool - false
- string - ""
- numbers (int, float, byte, rune) - 0
- pointer, interface, slice, channel, map, function - nil

Go 没有隐式类型转换，不同类型之间必须显式转换：

```go
var x int = 42
var y float64 = float64(x)   // int -> float64

var a float64 = 3.14
var b int = int(a)           // float64 -> int

var c int32 = 100
var d int64 = int64(c)       // int32 -> int64

var e int = 65
var f byte = byte(e)         // int -> byte (f = 'A')
```

## Constants

```go
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
```
常量可以是布尔、字符串、字符或数值。

数值常量在未被赋予具体类型前是 “无类型” 的，具有任意精度，并且在用于需要具体类型的上下文时会被赋予相应的类型。

在 Go 中，变量是不允许 float32 和 float64 直接混用的，但常量可以赋值给任何兼容的类型。

```go
const Pi = 3.14159
var f32 float32 = Pi
var f64 float64 = Pi
```

`iota` 是常量生成器，每个常量声明组内从 0 开始递增，可用于实现枚举和按位掩码。

```go
const (
	FlagA = 1 << iota // 1, 2, 4, ...
	FlagB
	FlagC
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)
```

## For

```go
func for_example() {
	// 最基础的方式，单个循环条件。
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的 “初始/条件/后续” for 循环。
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// 使用 range 关键字实现 “循环 n 次” 效果。
	for i := range 3 {
		fmt.Println("range: ", i)
	}

	// 不带条件的 for 循环将一直重复执行。
	// 直到在循环体内使用了 break 或者 return 跳出循环。
	for {
		fmt.Println("loop")
		break
	}

	// 也可以使用 continue 直接进入下一次循环。
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
```

for 是 Go 中唯一的循环结构。

## If/Else

```go
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
```

Go 没有三目运算符，即使是基本的条件判断，依然需要使用完整的 if 语句。

`if num := 9; num < 0` 中定义的 num，其作用域仅限于该 if-else 块，若在外部也定义了一个同名的 num，内部的 num 会 “遮蔽” 外部的值。

## Switch

```go
func switch_example() {
	// 一个基本的 switch。
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 使用逗号来分隔多个表达式，default 语句用于默认匹配。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// 不带表达式的 switch 是实现 if/else 逻辑的另一种方式。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// 使用 .(type) 会比较类型而非值。
	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
```

switch 是多分支情况时快捷的条件语句。
