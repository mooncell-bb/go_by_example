package structs

import "fmt"

type rect struct {
	width, height int
}

// 拥有指针类型的方法。
func (r *rect) area() int {
	return r.width * r.height
}

// 拥有值类型的方法。
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
