package main

import "fmt"

// const
const (
	c0 = iota
	c1
	c2
)

func main() {
	// string and rune
	// s := "hello world"
	// c := []rune(s)
	// for i := 0; i < len(c); i++ {
	// 	fmt.Print(s[i])
	// 	fmt.Print(" ")
	// }

	// pointer
	// var a = 11
	// p := &a
	// fmt.Println(*p)

	// struct pointer
	// type User struct {
	// 	name string
	// 	age  int
	// }
	// andes := User{
	// 	name: "adnes",
	// 	age:  18,
	// }
	// pp := &andes
	// fmt.Println(pp.name)

	// array
	// const maxn int = 100
	// var arr [maxn]int
	// for i := 0; i < maxn; i++ {
	// 	arr[i] = i
	// 	fmt.Printf("%d ", arr[i])
	// }
	// fmt.Println()
	// brr := [...]int{1, 2, 3}
	// for i := 0; i < 3; i++ {
	// 	fmt.Printf("%d ", brr[i])
	// }
	// a := [3]int{1: 1, 2: 3}
	// for i := 0; i < 3; i++ {
	// 	fmt.Printf("%d ", a[i])
	// }
	// a := [...]int{1, 2, 3}
	// for i := range a {
	// 	fmt.Printf("%d ", a[i])
	// }

	// slice
	// var array = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// s1 := array[0:4]
	// s2 := array[:4]
	// s3 := array[2:]
	// fmt.Printf("%v\n", s1)
	// fmt.Printf("%v\n", s2)
	// fmt.Printf("%v\n", s3)
	// a := make([]int, 10)
	// b := make([]int, 10, 15)
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	b := make([]int, 2, 4) // [0, 0], it's empty
	c := a[0:3]            // [0, 1, 2]
	fmt.Println(len(b))    // 2
	fmt.Println(cap(b))    // 4
	b = append(b, 1)       // append a 1 to end of b
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	b = append(b, c...)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	d := make([]int, 2, 2)
	copy(d, c)
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))
}
