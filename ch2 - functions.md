# Functions

## Basic

函数是 Go 的核心。

```go
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
```

在调用函数时，所有的参数传递都是 “值传递”，即对传递的值进行拷贝。

## Multiple Return Values

```go
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
```

在 Go 中，多返回值主要用于处理错误。

遵循 (Result, Error) 的模式，当函数执行可能失败时，约定俗成地将 error 作为最后一个返回值。

```go
func findUser(id int) (User, error) {
    if id <= 0 {
        return User{}, errors.New("invalid id")
    }

    return user, nil
}

func main() {
    u, err := findUser(10)
    if err != nil {
        return
    }
    
    fmt.Println(u.Name)
}
```

当函数 A 的返回值与函数 B 的参数值完全匹配时，可以直接调用。

```go
func getCoords() (int, int) {
    return 10, 20
}

func move(x, y int) {
    fmt.Printf("Moving to %d, %d\n", x, y)
}

func main() {
    move(getCoords())
}
```

## Variadic Functions

可变参数函数在调用时可以传递任意数量的参数，常见的 fmt.Println() 就是一个变参函数。

```go
// 此函数接受任意数量的 int 作为参数。
// 此处 nums 等价于传入 []int，可以调用 len(nums) 等函数。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func variadic_functions() {
	// 常规调用方式。
	sum(1, 2)
	sum(1, 2, 3)

	// 如果有一个含多个值的 slice，将其作为参数给变参函数，需要使用 'func(slice...)'。
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
```

在一个函数定义中，只能有一个可变参数，且它必须放在参数列表的最后一位。

使用 sum(1, 2, 3) 方式时，Go 编译器会在底层自动创建一个新的切片，将 1, 2, 3 拷贝进去，然后把切片传给函数。

使用 sum(nums...) 方式时，在参数内部修改了 slice，其外部的原切片也会被修改。

## Closures

Go 支持匿名函数，并能用其构造闭包。

```go
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
```

闭包 = 函数 + 引用环境，闭包不仅仅是匿名函数，它还 “捕获” 了外部作用域中的变量。其对外部变量的捕获是引用传递，而不是值传递，这意味着闭包内部对变量的修改会影响外部。

Go 1.22 之前，在循环里创建闭包引用循环变量会导致逻辑错误。

```go
func loop_closure() {
    funcs := []func(){}
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() {
            fmt.Println(i)
        })
    }
    
    for _, f := range funcs {
        f()
    }
}
```

Go 1.22 之前，闭包都会引用同一个 i 的地址，循环结束时 i 变成了 3。之后，每次迭代都会创建一个新的 i 变量。

可以用闭包来创建一个 “带有上下文” 的函数：

```go
func logger(name string) func() {
    return func() {
        fmt.Printf("[%s] action executed at %v\n", name, time.Now())
    }
}

func main() {
    adminLog := logger("ADMIN")
    userLog := logger("USER")
    
    adminLog()
}
```

## Recursion

Go 支持递归。

```go
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
```

Go 不支持尾递归优化，每一次递归调用都会占用一定的栈空间，但 Go 的 Goroutine 栈是动态伸缩的。

递归通常用于文件系统遍历、组织架构树等场景。

## Pointers

Go 支持指针，允许在程序中通过引用传递来传递值和数据结构。

```go
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
```

一个指针的默认初始值是 nil，若对指针进行 “解引用”，程序会立即崩溃。

内置函数 new() 会分配一块内存存放 int，将其初始化为零值 0，并返回其地址 *int。
