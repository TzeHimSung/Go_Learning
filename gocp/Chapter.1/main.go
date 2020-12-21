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
	s := "hello world"
	c := []rune(s)
	for i := 0; i < len(c); i++ {
		fmt.Print(s[i])
		fmt.Print(" ")
	}

	// pointer
	var a = 11
	p := &a
	fmt.Println(*p)

	// struct pointer
	type User struct {
		name string
		age  int
	}
	andes := User{
		name: "adnes",
		age:  18,
	}
	pp := &andes
	fmt.Println(pp.name)

	// array
	const maxn int = 100
	arr := [maxn]int{0, 1, 2, 3}
	for _, i := range arr { // Attention _ !
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	brr := [...]int{1, 2, 3}
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", brr[i])
	}
	a := [3]int{1: 1, 2: 3}
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", a[i])
	}
	a := [...]int{1, 2, 3}
	for i := range a {
		fmt.Printf("%d ", a[i])
	}

	// slice
	var array = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := array[0:4]
	s2 := array[:4]
	s3 := array[2:]
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s3)
	a := make([]int, 10)
	b := make([]int, 10, 15)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	b := make([]int, 2, 4) // [0, 0], it's empty
	c := a[0:3]            // [0, 1, 2]
	// fmt.Println(len(b))    // 2
	// fmt.Println(cap(b))    // 4
	b = append(b, 1)       // append a 1 to end of b
	fmt.Println(b)         // [0, 0, 1]
	fmt.Println(len(b))    // 3
	fmt.Println(cap(b))    // 4, not changed
	fmt.Println(cap(c))    // 7, it's length of array a
	b = append(b, c...)    // append c to end of b
	fmt.Println(b)         // [0, 0, 1, 0 ,1, 2]
	fmt.Println(len(b))    // 6
	fmt.Println(cap(b))    // why it is 8, not 7?
	d := make([]int, 2, 2) // [0, 0], it's empty
	copy(d, c)
	fmt.Println(d)      // [0, 1]
	fmt.Println(len(d)) // 2
	fmt.Println(cap(d)) // 2
	str := "hello, 世界"
	aa := []byte(str)
	ba := []rune(str)
	for _, i := range aa {
		fmt.Print(i)
		fmt.Print(" ")
	}
	for _, i := range ba {
		fmt.Print(i)
		fmt.Print(" ")
	}

	// map
	ma := map[string]int{"a": 1, "b": 2}
	for a, b := range ma {
		fmt.Print(a)
		fmt.Print(" ")
		fmt.Println(b)
	}
	mp1 := make(map[int]string)
	mp2 := make(map[int]string, 10) // fixed length
	mp1[1] = "tom"
	mp2[1] = "jerry"
	mp1[2] = "haha"
	for _, i := range mp1 { // the order is not fixed
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println()
	delete(mp1, 1)
	for _, i := range mp1 {
		fmt.Print(i)
		fmt.Print(" ")
	}
	type User struct {
		name string
		age  int
	}
	ma := make(map[int]User)
	andes := User{
		name: "andes",
		age:  18,
	}
	ma[1] = andes
	fmt.Println(andes.name, andes.age)
	andes.age = 19
	fmt.Println(andes.name, andes.age)

	// struct
	type Person struct {
		Name string
		Age  int
	}
	type Student struct {
		*Person
		Number int
	}
	a := Person{"Tom", 21}
	p := &Person{
		Name: "tata",
		Age:  12,
	}
	s := Student{
		Person: p,
		Number: 110,
	}
	fmt.Println(a.Age, a.Name)
	fmt.Println(s.Age, s.Name, s.Number)

	// if
	x := 5
	y := 10
	if x <= y {
		fmt.Println(y)
	} else {
		fmt.Println(x)
	}
	x := 5
	y := 10
	if a := 7; a < x {
		fmt.Println(x)
	} else if a < y {
		fmt.Println(y)
	} else {
		fmt.Println(a)
	}

	// switch
	switch i := "y"; i {
	case "y", "Y":
		fmt.Println("YES")
		fallthrough
	case "n", "N":
		fmt.Println("NO")
	}
	score := 85
	grade := ' '
	switch {
	case score >= 90:
		grade = 'A'
	case score >= 80:
		grade = 'B'
	case score >= 70:
		grade = 'C'
	case score >= 60:
		grade = 'D'
	default:
		grade = 'F'
	}
	fmt.Printf("grade=%c\n", grade)

	// for and continue
L1:
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if i >= 5 {
				continue L1
			}
			if j > 10 {
				continue
			}
		}
	}

}
