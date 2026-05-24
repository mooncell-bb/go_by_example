package structs

import "fmt"

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
	fmt.Println(d.id)       // 访问 derived 定义的 id。
	fmt.Println(d.base2.id) // 必须通过完整嵌入类型名来访问 base2 的 id。
}


