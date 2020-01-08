package main

import "fmt"

// Person ... named type
// type Person struct {
// 	name string
// 	age  int
// }

// Map ...
// type Map map[string]string
// type iMap Map
// type slice []int

// // Print ... this is method, not a function
// func (m Map) Print() {
// 	for _, value := range m {
// 		fmt.Println(value)
// 	}
// }

// func (m iMap) Print() {
// 	for _, value := range m {
// 		fmt.Println(value)
// 	}
// }

// func (s slice) Print() {
// 	for _, v := range s {
// 		fmt.Println(v)
// 	}
// }

// func main() {
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
// type Test = int
// var a Test = 1
// fmt.Printf("%T\n", a)
// fmt.Printf("%v\n", a)

// mp := make(map[string]string, 10)
// mp["hi"] = "tata"
// var ma Map = mp
// // the following code is illegal
// // var im iMap = ma
// ma.Print()
// var i interface {
// 	Print()
// } = ma
// i.Print()
// }

// Map ...
// type Map map[string]string
// type iMap Map

// // Print ...
// func (m Map) Print() {
// 	for _, value := range m {
// 		fmt.Println(value)
// 	}
// }

// func (m iMap) Print() {
// 	for _, value := range m {
// 		fmt.Println(value)
// 	}
// }

// func main() {
// 	mp := make(map[string]string, 10)
// 	mp["hi"] = "tata"
// 	var ma Map = mp
// 	var im iMap = (iMap)(ma)
// 	ma.Print()
// 	im.Print()
// }

// func main() {
// 	s := "hello, world!"
// 	var a []byte // byte is alia of uint8
// 	a = []byte(s)
// 	var b string = string(a)
// 	var c []rune // rune is alia of int32
// 	c = []rune(s)
// 	fmt.Printf("%T\n", a)
// 	fmt.Printf("%T\n", b)
// 	fmt.Printf("%T\n", c)
// }

// Person ...
type Person struct {
	name string
	age  int
}

func main() {
	a := Person{"andes", 18}
	b := Person{
		"andes",
		18,
	}
	c := Person{
		"andes",
		18}
	fmt.Println(a, b, c)
}
