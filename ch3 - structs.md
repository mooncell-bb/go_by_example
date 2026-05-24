# Structs

## Structs

Go 的结构体 (struct) 是带类型字段 (fields) 的集合。

```go
// person 结构体包含了 name 和 age 两个字段。
type person struct {
	name string
	age  int
}

// newPerson() 函数使用给定的名字构造一个新的 person 结构体。
// Go 有 GC 垃圾回收器，当没有活跃指针指向结构体时会被自动回收。
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

func structs_example() {
	// 可以简略或指定字段名字来创建结构体，省略的字段将被初始化为零值。
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})

	// & 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	// 使用 . 来访问结构体字段，指针会被自动解引用。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	// 结构体是可变的。
	sp.age = 51
	fmt.Println(sp.age)
}
```

在字段声明后，可以用反引号 ` 定义标签，此方式称为结构体标签，可以告诉 JSON 编码器或数据库引擎如何映射字段名。

```go
type User struct {
    ID       int    `json:"id"`
    UserName string `json:"user_name"`
    Password string `json:"-"`
}
```

只需要在某个特定的地方临时组织一下数据，可以使用匿名结构体 (anonymous structs)。

```go
config := struct {
    APIKey string
    Debug  bool
}{
    APIKey: "secret-123",
    Debug:  true,
}
fmt.Println(config.APIKey)
```

struct{} 是空结构体，不占用任何内存空间，通常用 map[string]struct{} 实现内置的 Set 集合，或使用 chan struct{} 来通知某个事件发生，但不传递任何实际数据。

字段定义的顺序会影响它占用的内存大小，在定义包含大量字段且会被创建数百万次的结构体时，可将相同类型的字段排在一起。

## Methods

Go 支持为结构体类型定义方法 (methods)。

```go
type rect struct {
	width, height int
}

// 拥有指针类型的方法，每次调用时只拷贝指针，内部修改直接影响外部原变量。
func (r *rect) area() int {
	return r.width * r.height
}

// 拥有值类型的方法，每次调用都会拷贝整个结构体，内部修改不影响外部原变量	
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func methods_example() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// 调用方法时，Go 会自动处理值和指针之间的转换。
	// 可以使用指针来避免在调用方法时产生一个拷贝。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
```

不限于 struct，可以为当前包内定义的任何类型增加方法：

```go
type MyInt int

func (m MyInt) IsPositive() bool {
    return m > 0
}
```

首字母大写，方法可以被其它包访问；首字母小写，方法只能在当前包内使用。

## Interfaces

方法签名的集合叫做接口 (interfaces)。

```go
// 几何体的基本接口。
type geometry interface {
	area() float64
	perim() float64
}

// rect_new、circle 实现接口方法。
type rect_new struct {
	width, height float64
}

func (r rect_new) area() float64 {
	return r.width * r.height
}
func (r rect_new) perim() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 一个变量实现了某个接口，就可以调用指定接口中的方法。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect_new{width: 3, height: 4}
	c := circle{radius: 5}

	// circle 和 rect_new 都实现了 geometry 接口，可以将其实例作为 measure 的参数。
	measure(r)
	measure(c)
}
```

Go 接口是隐式实现，只要一个类型拥有了接口要求的所有方法，它就自动实现了该接口。

Go 1.18 引入了 any 关键字，它是 interface{} 的别名，任何类型都实现了空接口。

使用 "断言"，可以使接口转化为具体类型：

```go
if r, ok := g.(rect_new); ok {
    ...
}

switch v := g.(type) {
case rect_new:
    ...
case circle:
    ...
default:
    ...
}
```

如果接口方法是用值接收者实现的，那么 "值变量" 和 "指针变量" 都实现了该接口。如果接口方法是用指针接收者实现的，那么只有 "指针变量" 实现了该接口。

```go
type notifier interface { notify() }

type user struct{}
func (u *user) notify() { fmt.Println("Sending notify") }

var n notifier
u := user{}
n = &u // n = u 错误，编译不通过。
```

接口在底层由两个部分组成：动态类型和动态值，只有当类型和值都为 nil 时，接口才等于 nil。将具体类型赋值给接口变量时，Go 会在运行时进行装箱，创建一个 iface 结构。

```go
var r *rect_new = nil
var g geometry = r

if g != nil {
    fmt.Println("g") // 会被打印。
}
```

单方法接口通常以方法名 + er 后缀命名，这是 Go 社区的惯用风格：

```go
type Reader interface { Read(p []byte) (n int, err error) }
type Writer interface { Write(p []byte) (n int, err error) }
type Stringer interface { String() string }
```

error 是 Go 中最核心的接口，定义在 builtin 包：

```go
type error interface {
    Error() string
}
```

任何实现了 Error() string 方法的类型都是一个 error：

```go
type MyError struct {
    Msg  string
    Code int
}

func (e *MyError) Error() string {
    return fmt.Sprintf("code %d: %s", e.Code, e.Msg)
}
```

fmt.Stringer 接口的作用等同于 Java 的 toString()：

```go
type Stringer interface {
    String() string
}
```

实现 String() 方法后，fmt.Print 系列函数会自动调用它来打印自定义格式。

## Struct Embedding

Go 支持结构体和接口的嵌入（Embedding），以实现更无缝的类型组合。

嵌入的字段是一个没有名字的匿名字段（anonymous field），外层结构体可以直接访问被嵌入类型的字段和方法。

```go
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

// container 嵌入了 base，其看起来像一个没有名字的字段。
type container struct {
	base
	str string
}

func struct_embedding_example() {
	// 当创建含有嵌入的结构体，必须对嵌入进行显式的初始化。
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// 可以直接在 co 上访问基类的字段，例如 co.num。
	// 也可以使用嵌入的类型名称来写出完整的路径，例如 co.base.num。
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)

	// 由于 container 嵌入了 base，因此 base 的方法也成为了 container 的方法。
	fmt.Println("describe: ", co.describe())

	// 可以使用带有方法的嵌入结构来赋予接口实现到其他结构上。
	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println("describer:", d.describe())
}
```

也可嵌入结构体指针，指针嵌入允许 nil 值，且多个外层结构体可以共享同一个嵌入实例。

```go
type sharedPool struct {
	*bufPool
	name  string
}
```

外层结构体定义了与嵌入类型同名的字段或方法时，外层会遮蔽嵌入层的成员，此时必须通过完整嵌入类型名来访问被遮蔽成员：

```go
type base2 struct {
	id int
}

type derived struct {
	base2
	id int // 此处遮蔽了 base2.id。
}

func shadow_example() {
	d := derived{
		base2: base2{id: 1},
		id:    2,
	}
	fmt.Println(d.id) // 访问 derived 定义的 id。
	fmt.Println(d.base2.id) // 必须通过完整嵌入类型名来访问 base2 的 id。
}
```

一个结构体可以同时嵌入多个类型，实现更灵活的组合。Go 鼓励用嵌入组合（Composition）代替传统的继承。

```go
type writer struct{ buf []byte }

func (w *writer) Write(p []byte) (n int, err error) {
	w.buf = append(w.buf, p...)
	return len(p), nil
}

type closer struct{}

func (c *closer) Close() error {
	fmt.Println("closed")
	return nil
}

// 多重嵌入：组合多个行为。
type writeCloser struct {
	*writer
	*closer
}
```

接口也可以嵌入到结构体中，用于依赖注入或延迟赋值，当接口字段为 nil 时调用其方法会 panic。

```go
type Logger interface {
	Log(string)
}

type Server struct {
	Logger // 嵌入接口。
	Addr   string
}

func (s *Server) Start() {
	s.Log("server started on " + s.Addr)
}
```

接口同样可以嵌入到另一个接口中，组合出新的接口：

```go
type ReadWriter interface {
	Reader
	Writer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
```

嵌入不限于结构体和接口，任何命名类型都可以作为匿名字段嵌入：

```go
type counter int

type stats struct {
	counter // 嵌入命名类型。
	label   string
}

func stats_example() {
	s := stats{counter: 0, label: "requests"}
	s.counter++
	fmt.Println(s.counter)
}
```

## Generics

从 1.18 版本开始，Go 添加了对泛型的支持，也即类型参数。

```go
// SlicesIndex() 函数接受任意 comparable 类型的切片，和该类型的一个元素。
// comparable 的约束意味着可以使用 == 等比较运算符。
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}

	return -1
}

// List 是一个包含任意类型值的单向链表。
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// 泛型类型的方法，传入的类型是 List[T]，而不是 List。
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}

	return elems
}

func generics_example() {
	// 可以依赖类型推断。
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo: ", SlicesIndex(s, "zoo"))

	// 也可以明确指定类型。
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list: ", lst.AllElements())
}
```

约束中的 `~` 表示近似约束，它允许匹配底层类型相同的所有类型，而不仅仅是字面类型本身。例如 `S ~[]E` 不仅匹配 `[]E`，还匹配任何以 `[]E` 为底层类型的命名类型。

```go
type MySlice []int

func PrintElements[S ~[]E, E any](s S) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func approx_example() {
	ms := MySlice{1, 2, 3}
	PrintElements(ms)
}
```

Go 1.18 引入了 `any` 和 `comparable` 两个预声明约束。`any` 等价于 `interface{}`，不限制类型；`comparable` 限制类型必须支持 `==` 和 `!=` 运算符。

自定义约束通过接口语法定义，不仅可以列出允许的具体类型，还可以要求实现特定方法：

```go
// 只允许 int 和 float64。
type Number interface {
	~int | ~float64
}

// 必须同时满足类型集和方法签名。
type Stringer interface {
	~int | ~string
	String() string
}

// 使用自定义约束。
func Sum[T Number](vals []T) T {
	var total T
	for _, v := range vals {
		total += v
	}
	return total
}
```

泛型函数支持多类型参数，且每个参数可以有独立的约束：

```go
func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
    
	return result
}
```

