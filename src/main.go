package main

import "fmt"

// Person ... named type
type Person struct {
	name string
	age  int
}

func main() {
	// unnamed type
	// a := struct {
	// 	name string
	// 	age  int
	// }{"andes", 18}
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%v\n", a)
	// b := Person{"tom", 21}
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%v\n", b)

	// named type, like typedef
	// type Test int
	// var a Test = 1
	// fmt.Println(a)
	// fmt.Printf("%T\n", a)

	// alias
	type Test = int
	var a Test = 1
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)
}
