package main

import "fmt"

func main() {

	var i int
	var pi *int

	fmt.Println(i)
	fmt.Println(pi)
	fmt.Println()

	i = 34
	pi = &i
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(pi)
	fmt.Println(&pi)
	fmt.Println()

	fmt.Println(pi)
	fmt.Println(*pi)
	fmt.Println()

	*pi = 78

	fmt.Println(i)

}
