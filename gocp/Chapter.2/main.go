// package main

// import "fmt"

// // A ... this function can be viewed by other package
// func A() {
// 	// do something
// }

// func add(a, b int) int {
// 	return a + b
// }

// func add2(a, b int) (sum int) {
// 	sum = a + b
// 	return
// }

// func swap(a, b int) (int, int) {
// 	return b, a
// }

// func main() {
// 	fmt.Println(add(3, 4))
// 	fmt.Println(add2(3, 4))
// 	fmt.Println(swap(3, 4))
// }

// func chvalue(a int) int {
// 	a = a + 1
// 	return a
// }

// func chpointer(a *int) {
// 	*a = *a + 1
// 	return
// }

// func main() {
// 	a := 10
// 	chvalue(a) // value copy
// 	fmt.Println(a)
// 	chpointer(&a) // copy address value
// 	fmt.Println(a)
// }

// func add(a, b int) (sum int) {
// 	sum = a + b
// 	return
// }

// func sum(arr ...int) (sum int) { // this argument is a slice, "..." is after slice's name, not front at type name
// 	for _, v := range arr { // enum each index and value
// 		sum = sum + v
// 	}
// 	return
// }

// func sumb(arr []int) (sum int) {
// 	for _, v := range arr {
// 		sum = sum + v
// 	}
// 	return
// }

// func main() {
// 	sli := []int{1, 2, 3, 4} // this is slice
// 	// arr := [...]int{1, 2, 3, 4} // this is array
// 	fmt.Println(sum(sli...)) // array can not do that
// 	fmt.Println(sumb(sli))
// 	// this two function is not same type
// 	fmt.Printf("%T\n", add)
// 	fmt.Printf("%T\n", sum)
// 	fmt.Printf("%T\n", sumb)
// }

// func add(a, b int) int {
// 	return a + b
// }

// func sub(a, b int) int {
// 	return a - b
// }

// //Op ...
// type Op func(int, int) int

// func do(f Op, a, b int) int {
// 	return f(a, b)
// }

// func main() {
// 	a := do(add, 1, 2)
// 	s := do(sub, 1, 2)
// 	f := add
// 	fmt.Println(a, s, f(1, 2))
// }

// var sum = func(a, b int) int {
// 	return a + b
// }

// func doinput(f func(int, int) int, a, b int) int {
// 	return f(a, b)
// }

// func wrap(op string) func(int, int) int {
// 	switch op {
// 	case "add":
// 		return func(a, b int) int {
// 			return a + b
// 		}
// 	case "sub":
// 		return func(a, b int) int {
// 			return a - b
// 		}
// 	default:
// 		return nil
// 	}
// }

// func main() {
// 	defer func() { // no need to care what it is
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()

// 	sum(1, 2)

// 	doinput(func(x, y int) int {
// 		return x + y
// 	}, 1, 2)

// 	opFunc := wrap("add")
// 	re := opFunc(2, 3)
// 	fmt.Println(re)
// }

// func f() int {
// 	a := 0
// 	defer func(i int) {
// 		println("defer i =", i)
// 	}(a)

// 	a++
// 	return a
// }

// func main() {
// 	// defer func() {
// 	// 	println("first")
// 	// }()

// 	// defer func() {
// 	// 	println("second")
// 	// }()

// 	// println("function body")

// 	// f()
// }

// func main() {
// 	defer func() {
// 		println("defer")
// 	}()
// 	println("func body")
// 	os.Exit(1)
// }

// func copyFile(dst, src string) (w int64, err error) {
// 	src, err := os.Open(src)
// 	if err != nil {
// 		return
// 	}
// 	dst, err := os.Create(dst)
// 	if err != nil {
// 		src.Close()
// 		return
// 	}
// 	w, err = io.Copy(dst, src)
// 	dst.Close()
// 	sec.Close()
// 	return
// }

// func copyFile(dst, src string) (w int64, err error) {
// 	src, err := os.Open(src)
// 	if err != nil {
// 		return
// 	}
// 	defer src.Close()
// 	dst, err := os.Create(dst)
// 	if err != nil {
// 		return
// 	}
// 	defer dst.Close()
// 	w, err = io.Copy(dst, src)
// 	return
// }

// closure
// func fa(a int) func(i int) int {
// 	return func(i int) int {
// 		println(&a, a) // print the address and current value of a
// 		a = a + i
// 		return a
// 	}
// }

// func main() {
// 	f := fa(1)    // create a var a
// 	g := fa(1)    // create a var a
// 	println(f(1)) // these two call of f called the same copy of a
// 	println(f(1))
// 	println(g(1)) // same of above, but this a is different from the above one (from address we can know)
// 	println(g(1))
// }

// var (
// 	a = 0
// )

// func fa() func(i int) int {
// 	return func(i int) int {
// 		println(&a, a)
// 		a = a + i
// 		return a
// 	}
// }

// func main() {
// 	f := fa()
// 	g := fa()
// 	println(f(1))
// 	println(f(1))
// 	println(g(1))
// 	println(g(1))
// }

// func fa(base int) (func(int) int, func(int) int) {
// 	println(&base, base)
// 	add := func(i int) int {
// 		base += i
// 		println(&base, base)
// 		return base
// 	}
// 	sub := func(i int) int {
// 		base -= i
// 		println(&base, base)
// 		return base
// 	}
// 	return add, sub
// }

// func main() {
// 	f, g := fa(0)
// 	s, k := fa(0)
// 	println(f(1), g(2))
// 	println(s(1), k(2))
// 	// result
// 	// 0xc000098000 0
// 	// 0xc000098008 0
// 	// 0xc000098000 1
// 	// 0xc000098000 -1
// 	// 1 -1
// 	// 0xc000098008 1
// 	// 0xc000098008 -1
// 	// 1 -1
// }

// func except() {
// 	recover()
// }

// func test() {
// 	defer except()
// 	panic("test panic")
// }

// func main() {
// 	test()

// }

// func main() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()

// 	defer func() {
// 		panic("first defer panic")
// 	}()

// 	defer func() {
// 		panic("second defer panic")
// 	}()

// 	panic("main body panic")
// }

// error
