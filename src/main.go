package main

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
// type Person struct {
// 	name string
// 	age  int
// }

// Element ...
// type Element struct {
// 	next, prev *Element
// 	Value      interface{}
// }

// func main() {
// a := Person{"andes", 18}
// b := Person{
// 	"andes",
// 	18,
// }
// c := Person{
// 	"andes",
// 	18}
// fmt.Println(a, b, c)
// p := new(Person) // p is a pointer here
// p2 := Person{}
// p2.age = 11
// p2.name = "andes"
// fmt.Println(p, p2)
// fmt.Println("111")
// }

// method
// func (t TypeName) MethodName(ParamList) Returnlist {

// }

// func (t *TypeName) MethodName(ParamList) Returnlist {

// }

// SliceInt ...
// type SliceInt []int

// // Sum ...
// func (s SliceInt) Sum() (ret int) {
// 	ret = 0
// 	for _, i := range s {
// 		ret += i
// 	}
// 	return
// }

// func main() {
// 	var s SliceInt = []int{1, 2, 3, 4}
// 	fmt.Println(s.Sum())
// }

// Map ...
// type Map map[string]string
// type myInt int

// // Print ...
// func (m Map) Print() {
// 	for _, key := range m {
// 		fmt.Print(key, " ")
// 	}
// }

// func main() {
// 	var a myInt = 10
// 	var b myInt = 10
// 	c := a + b
// 	d := a * b
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// type t struct {
// 	a int
// }

// func (tt t) get() int {
// 	return tt.a
// }

// func (tt *t) set(val int) {
// 	tt.a = val
// }

// func main() {
// 	tt := &t{}
// 	tt.set(2)
// 	fmt.Println(tt.get())
// }

// type t struct {
// 	a int
// }

// func (tt t) get() int {
// 	return tt.a
// }

// func (tt *t) set(val int) {
// 	tt.a = val
// }

// func (tt *t) print() {
// 	// %p is the address of var
// 	fmt.Printf("%p %v %d\n", tt, tt, tt.a)
// }

// func main() {
// 	tt := &t{}
// 	f := tt.set
// 	f(2)
// 	tt.print()
// 	f(3)
// 	tt.print()
// }

// type t struct {
// 	a int
// }

// func (tt *t) set(val int) {
// 	tt.a = val
// }

// func (tt t) get() int {
// 	return tt.a
// }

// func (tt *t) print() {
// 	fmt.Printf("%p %v %d\n", tt, tt, tt.a)
// }

// func main() {
// 	tt := t{a: 1}
// 	fmt.Println(tt.get())
// 	(*t).set(&tt, 2)
// 	fmt.Println((t).get(tt))
// }

// Int ...
// type Int int

// // Max ...
// func (a Int) Max(b Int) Int {
// 	if a >= b {
// 		return a
// 	}
// 	return b
// }

// func (a *Int) set(b Int) {
// 	*a = b
// }

// func (a Int) print() {
// 	fmt.Printf("value = %d\n", a)
// }

// func main() {
// 	var a Int = 10
// 	var b Int = 20
// 	c := a.Max(b)
// 	c.print()
// 	(&c).print()
// 	(&a).set(30)
// 	a.print()
// 	a.set(10)
// 	a.print()
// }
