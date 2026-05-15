# Data structures

## Arrays

数组是一个具有特定长度的编号元素序列，但在典型的 Go 代码中，切片更为常见。

```go
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
```

在 Go 中，数组承载了值。当把一个数组赋值给另一个数组，或者作为参数传递给函数时，Go 会完整地复制整个数组的内容。由于长度固定且传递时会产生拷贝，它不适合处理大规模或动态长度的数据流。

```go
a := [3]int{1, 2, 3}
b := a
b[0] = 100
fmt.Println(a) // [1, 2, 3]
fmt.Println(b) // [100, 2, 3]
```

如果数组的元素类型是可比较的，那么数组本身也是可比较的。

```go
a := [2]int{1, 2}
b := [2]int{1, 2}
c := [2]int{2, 1}
fmt.Println(a == b) // true
fmt.Println(a == c) // false
```

只有当明确知道数据长度固定且不会改变时，才使用数组。

## Slices

切片是 Go 中重要的数据类型，它提供了比数组更强大的序列交互方式。

```go
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
```

### 切片头

切片本身并不是数组，它是一个运行时结构体，包括：

- Pointer：指向底层数组的指针。
- Len（长度）：当前切片内的元素个数。
- Cap（容量）：底层数组中的元素数量。

把切片传递给函数时，Go 拷贝的是切片头，而不是底层的数组。函数内部修改切片元素时，外部也会变。

make() 内置函数创建切片时，第一个参数表示长度，第二个参数表示容量（可选）：

```go
s := make([]int, 3, 6)
```

访问长度以外的元素时，会引发 panic 错误。可以使用内置的 append() 函数向切片 s 添加一个新元素，此时长度会 +1。

### 扩容机制

Len 是当前能访问的范围，Cap 是不需要重新分配内存就能增长到的极限。

当 append() 超过 Cap 时，Go 会申请一块更大的内存，将旧数据拷贝过去并指向新的数组。

如果 oldCap < 256，则 newCap 直接翻倍，如果 oldCap > 256，newCap 每次增长约 25% 并加上一个常数项。

将切片的长度设置为 0，依据切片扩容两倍的机制，会在初期添加元素时多次触发扩容机制。因此使用给定的长度或容量进行初始化，可以提升运行效率。

```go
func convert(foos []Foo) []Bar {
    n := len(foos)
    bars := make([]Bar, 0, n)

    for _, foo := range foos {
        bars = append(bars, fooToBar(foo))
    }
    
    return bars
}
```

```go
func convert(foos []Foo) []Bar {
    n := len(foos)
    bars := make([]Bar, n)

    for _, foo := range foos {
        bars[i] = fooToBar(foo)
    }
    
    return bars
}
```

### 切片操作

通过 's[low:high]' 创建的新切片与原切片共享同一个底层数组。

```go
s1 := make([]int, 3, 6)
s2 := s1[1:3]
```

s2 从索引 1 开始而不再是索引 0。如果更新 s1[1] 或 s2[0]，更改会作用于同一个数组，两个切片都能看到这个变化。

如果在 s2 上追加一个元素，则只有 s2 的 Len 属性会变化，s1 的各个属性依然不会改变，添加的数据仅对 s2 可见。若继续向 s2 追加元素，直到底层数组没有足够的容量时，会导致创建另一个底层数组并使 s2 指向该数组。

```go
s1 := make([]int, 3, 6)
s2 := s1[1:3]
fmt.Println(s1) // [0 0 0]
fmt.Println(s2) // [0 0]

s2 = append(s2, 2)
fmt.Println(s1) // [0 0 0]
fmt.Println(s2) // [0 0 2]

s1 = append(s1, 1)
fmt.Println(s1) // [0 0 0 1]
fmt.Println(s2) // [0 0 1]

s2 = append(s2, 2, 2, 2)
s1 = append(s1, 1, 1)
fmt.Println(s1) // [0 0 0 1 1 1]
fmt.Println(s2) // [0 0 1 2 2 2]
```

append() 触发扩容，将返回指向一个全新的底层数组。若不接收返回值，旧变量依然指向旧数组，将丢失新加入的元素。

若要在子切片中使用 append() 函数，且又不希望改变原切片，可以传递切片的一个副本。

另一种选项是使用完整切片表达式 s[low:high:max]，该方式会将生成切片的容量设置为 max - low。append() 函数发现切片已满，则会自行复制一个新的底层数组来存储新切片值。

### 空切片

空切片有两种类型，可表示为：Nil Slice 和 Empty Slice。

- Nil Slice：使用 var s []int 初始化，此时不分配内存，且 len(s) = 0、s == nil。
- Empty Slice：使用 s := []string{} 初始化，此时分配了内存，且 len(s) = 0、s != nil。

```go
var s []string // 1: empty = true	nil = true
log(1, s)

s = []string(nil) // 2: empty = true	nil = true
log(2, s)

s = []string{} // 3: empty = true	nil = false
log(3, s)

s = make([]string, 0) // 4: empty = true	nil = false
log(4, s)

func log(i int, s []string) {
	fmt.Printf("%d: empty = %t\tnil = %t\n", i, len(s) == 0, s == nil)
}
```

- []string(nil) 与 var s []string 效果一致，可以作为语法糖提供便利。
- s := []string{} 推荐用于创建带有初始元素的切片，否则不推荐使用该方式。

由于 Nil 切片不需要任何分配，处于节省内存的角度，应该倾向于返回 Nil 切片而不是 Empty 切片。若在需要生成已知长度的切片的情况下，我们应该使用选项 s := make([]string, length)。

一些库会区分 Nil 切片和 Empty 切片，例如 encoding/json 包对这两个结构体进行序列化会得到不同结果：

```go
type customer struct {
    ID         string
    Operations []float32
}

// 1. 使用 Nil 切片。
var s1 []float32
customer1 := customer{
    ID:         "foo",
    Operations: s1,
}
b, _ := json.Marshal(customer1)
fmt.Println(string(b))

// 2. 使用 Empty 切片。
s2 := make([]float32, 0)
customer2 := customer{
    ID:         "bar",
    Operations: s2,
}
b, _ = json.Marshal(customer2)
fmt.Println(string(b))
```

```
{"ID":"foo","Operations":null}
{"ID":"bar","Operations":[]}
```

通常在逻辑判断中使用 len(s) == 0 来判断切片是否为空，这样无论是 Nil 还是 Empty 都能正确处理。

### 切片复制

copy() 内置函数支持将源切片的元素复制到目标切片，目标参数（dst）在前，而源参数（src）在后。

但是 copy() 函数拷贝的元素个数是 min(len(dest), len(src))，若 src 长度为 0，则不会发生复制。

```go
src := []int{0, 1, 2}
var dst []int
copy(dst, src)
fmt.Println(dst) // []
```

还可以使用 append() 的方式来复制切片：

```go
src := []int{0, 1, 2}
dst := append([]int(nil), src...)
```

### 内存泄漏

存在一个 1GB 大小的切片，但只需要少数几个元素，于是使用了 small := large[:2] 来获取子切片。但即使 large 变量不再使用，只要 small 还存在，Go 的 GC 就不会释放元素。

为了解决此问题，可以创建切片副本而非直接使用子切片：

```go
small := make([]byte, 2)
copy(small, large[:2])
```

## Maps

Map 是 Go 内建的关联数据类型，其他语言中也被称为哈希或字典。

```go
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
```

在 Go 的 Map 中，如果一个 Key 不存在，它会返回该 Value 类型的零值。如果存放了一个键值对是 "a": 0，当取 m["a"] 时得到 0，则无法分辨是 “键不存在” 还是 “值本来就是 0”。

因此使用 '_, prs := m["k2"] 方式'，来判断一个 Key 是否存在。

### 类型限制

不是所有类型都能作为 Map 的键，其类型必须是可比较的（可以使用 == 运算符）。

允许的类型：布尔、数值、字符串、指针、通道、接口、以及只包含可比较元素的结构体。

禁止的类型：切片、Map、函数。

### 扩容机制

在 Go 语言中，映射基于哈希表数据结构实现，哈希表是由多个桶（buckets）组成的数组，每个桶指向一个键值对数组。

初始的键值对数组默认有 8 个，在插入到一个已满的桶的情况下，会创建另一个包含 8 个元素的桶，并将前一个桶链接到它。

对于读取、更新和删除操作，Go 必须计算对应的数组索引，然后顺序遍历所有键，直到找到提供的键。因此，这三种操作的最坏时间复杂度为 O(p)，其中 p 是桶中元素的总数（默认情况下为一个桶，溢出时则为多个桶）。

当映射增长时，其桶的数量会翻倍，Map 增长的条件为：

- 桶中项目的平均数量（称为负载因子）大于一个常数值，这个常数值默认为 6.5。
- 过多的桶已溢出（包含超过 8 个元素）。

当映射扩容时，所有键会被重新分配到各个桶中。最坏情况下，插入一个键可能成为 O(n) 操作，n 代表映射中元素的总数。

可以使用内置函数 make() 在创建映射时指定初始大小。对于映射，它只接受单个初始化参数，而不像切片还可以指定容量。

```go
m := make(map[string]int, 1000000)
```

### 内存泄漏

从 Map 中移除元素并不会影响现有桶的数量，它只是将桶中的槽位清零。Map 只能增长并拥有更多桶，而永远不会收缩。因此，删除元素后映射仍保持着相同数量的桶。

如果需要 Map 自动缩容，一种方案是定期重新创建当前映射表的副本。例如，每小时构建一个新 Map 表，复制所有元素，然后释放旧表。这种方法的主要缺点是，在复制完成后、下一次垃圾回收之前，短时间内可能会消耗双倍于当前内存的资源。

另一种解决方案是将存储类型改为存储指针，每个桶条目将仅为值预留指针大小的空间。

### 空 Map

和切片的情况类似，空 Map 可分为 Nil Map 和 Empty Map。

- Nil Map：var m map[string]int
  - 使用 v := m["any"] 会得到零值 0，不会报错。
  - 使用 m["any"] = 1 会直接 Panic。
- Empty Map：make(map[string]int)
  - 读写都是安全的。

应该始终使用 make() 初始化 Map，或者在写入前检查是否为 nil。

## Range

range 用于迭代各种各样的数据结构，其也可以遍历 Arrays、Slices 和 Maps。

```go
func range_example() {
	// 遍历数组求和。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range 在数组和 slice 中提供对每项的索引和值的访问。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range 在 map 中迭代键值对。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range 也可以只遍历 map 的键。
    // Go 为了防止开发者依赖 Map 的顺序，在每次 for range 迭代时，都会随机化遍历顺序。
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range 在字符串中迭代 unicode 码点。
	// 第一个返回值是字符的起始字节位置，然后第二个是字符本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
```

### 值拷贝

使用 'for i, v := range slice' 时，变量 v 是元素的一个副本，而不是对原始元素的引用。

```go
type account struct {
    balance float32
}

accounts := []account{
    {balance: 100.},
    {balance: 200.},
    {balance: 300.},
}

for _, a := range accounts {
    a.balance += 1000
}

fmt.Println(accounts) // [{100} {200} {300}]
```

在 Go 中，分配的一切都是副本：

- 赋值一个返回结构体的函数的结果，它会执行该结构体的复制。
- 赋值一个返回指针的函数结果，它会复制内存地址。

若想要更新切片元素，主要有两种选择。第一种选择是使用切片索引访问元素：

```go
for i := range accounts {
    accounts[i].balance += 1000
}
```

另一个选择是继续使用范围循环并访问值，但将切片类型修改为账户指针切片：

```go
accounts := []*account{
    {balance: 100.},
    {balance: 200.},
    {balance: 300.},
}

for _, a := range accounts {
    a.balance += 1000
}
```

> **指针陷阱**
>
> Go 1.22 之前，使用 'i, a := range accounts' 循环时，只会创建一个临时空间 a，每一次迭代后的 a 都会被赋值为其它变量。因此，若使用 b[i] = &a 来存放临时 a 元素的地址时，b[] 中的所有元素都会指向同一个 a 地址，且该地址最终指向迭代后的最后一个地址。
>
> Go 1.22 之后，每一轮循环都会创建一个新的变量 a，在循环内取 &a 并存入列表，可以得到预期的不同地址。

### 表达式求值

使用 'for i, v := range slice' 时，range 在开始迭代前会先对切片长度做一个 “快照”，只会迭代到快照时的长度。

当表达式 slice 被求值时，结果是一个切片副本，其 Len 和 Cap 属性值都和原切片一样，也会指向同一个底层数组。

```go
s := []int{0, 1, 2}
for range s {
    s = append(s, 10)
}
```

此循环与经典 for 循环的行为不同，其不会无限迭代下去，因为改变外部 s 的属性并不会影响里面 s 切片副本的属性。

对于数组，则是完全的值拷贝，它会创建一个与原数组一致的副本，且迭代过程中不会影响数组副本：

```go
a := [3]int{0, 1, 2}
for i, v := range a {
    a[2] = 10
    if i == 2 {
        fmt.Println(v) // 2
    }
}
```

如果想打印最后一个元素的实际值，可以通过两种方式实现。可通过从其索引访问元素，或使用数组指针实现：

```go
a := [3]int{0, 1, 2}
for i := range a {
    a[2] = 10
    if i == 2 {
        fmt.Println(a[2])
    }
}
```

```go
a := [3]int{0, 1, 2}
for i, v := range &a {
    a[2] = 10
    if i == 2 {
        fmt.Println(v)
    }
}
```

若迭代一个元素巨大的切片，使用 'for i, v := range slice' 会产生大量的值拷贝开销。

若不需要值，只声明 'for i := range slice'，此时 Go 编译器就不会去执行拷贝操作。

### Map 迭代

Map 的迭代顺序是未指定的，不能保证一次迭代到下一次迭代的顺序相同。

在 Go 中，允许在迭代期间更新 Map（插入或删除元素），且不会导致编译错误或运行时错误。

```go
m := map[int]bool{
    0: true,
    1: false,
    2: true,
}
 
for k, v := range m {
    if v {
        m[10+k] = true
    }
}
 
fmt.Println(m)
```

这段代码的结果是不可预测的：如果在迭代过程中创建了一个映射条目，它可能在迭代过程中被产生，也可能被跳过。

如果想在迭代映射时更新它，并确保添加的条目不属于迭代的一部分，解决方案是循环迭代 Map 的副本。

