package structs

import (
	"fmt"
	"math"
)

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
