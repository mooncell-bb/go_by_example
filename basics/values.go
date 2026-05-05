package basics

import "fmt"

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
