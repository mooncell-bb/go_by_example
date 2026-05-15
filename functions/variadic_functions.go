package functions

import "fmt"

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

func variadic_functions_example() {
	// 常规调用方式。
	sum(1, 2)
	sum(1, 2, 3)

	// 如果有一个含多个值的 slice，将其作为参数给变参函数，需要使用 'func(slice...)'。
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
