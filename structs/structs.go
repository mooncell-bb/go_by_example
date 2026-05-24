package structs

import "fmt"

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
