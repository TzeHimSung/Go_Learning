package main

import "fmt"

func main() {
	// 相同类型的变量可以在末尾加上类型
	var a, b int = 1, 2
	// 如果不带上类型，编译器可以自动推断
	var c, d, e = "haha", 3.14, 2
	// 不同类型的变量声明和隐式初始化可以使用如下语法
	var (
		x int
		y string
	)
	// 注意，多值赋值语句中每个变量后面都不能带上类型
	fmt.Println(a, b, c, d, e, x, y)
}

package main

import "fmt"

func f() (a, b int) { return 1, 2 }

func main() {
	// 函数调用
	x, y := f()
	fmt.Println(x, y)
	// range表达式
	a := []int{5, 4, 3, 2, 1}
	for index, val := range a {
		fmt.Printf("%d:%d ", index, val)
	}
	fmt.Println()
	// type assertion
	// v, ok := i.(xxxx)
	var p, o, i = 1, 3, 4
	fmt.Println(p, o, i)
}

package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	i := 0
	i, x[i] = 1, 2
	fmt.Println(i, x)

	x = []int{1, 2, 3}
	i = 0
	x[i], i = 2, 1
	fmt.Println(i, x)

	x = []int{1, 2, 3}
	i = 0
	x[i], i = 2, x[i]
	fmt.Println(i, x)

	x[0], x[0] = 1, 2
	fmt.Println(x[0])
}

package main

import "fmt"

var n int

func foo() (int, error) {
	return 1, nil
}

// 访问全局变量n
func g() {
	fmt.Println(n)
}

func main() {
	// 此时main函数作用域里面没有n，所以创建新的局部变量n
	n, _ := foo()
	// 访问全局变量n
	g()
	// 访问的是局部变量n
	fmt.Println(n)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range si {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range si {
		wg.Add(1)
		// 这里有一个实参到形参的值拷贝
		go func(a int) {
			fmt.Println(a)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

package main

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f4() (r int) {
	defer func(r *int) {
		*r = *r + 5
	}(&r)
	return 1
}

func main() {
	fmt.Println("f1 =", f1()) // 1
	fmt.Println("f2 =", f2()) // 5
	fmt.Println("f3 =", f3()) // 1
	fmt.Println("f4 =", f4()) // 6
}

package main

import "fmt"

func f4() int {
	r := 0
	defer func() {
		r++
	}()
	return r
}

func f5() int {
	r := 0
	defer func(i int) {
		i++
	}(r)
	return 0
}

func main() {
	fmt.Println(f4(), f5()) // 0, 0
}

package main

import "fmt"

func main() {
	// 指定大小的显式初始化
	a := [3]int{1, 2, 3}
	// 通过"..."由后面的元素个数推断数组大小
	b := [...]int{1, 2, 3}
	// 指定大小，并通过索引值初始化，未显式初始化的元素被置为零值
	c := [3]int{1: 1, 2: 3}
	// 指定大小但不显式初始化，数组元素全被置为零值
	var d [3]int
	fmt.Println(a, b, c, d)
}

package main

import "fmt"

func f(a [3]int) {
	a[2] = 10
	fmt.Printf("%p, %v\n", &a, a)
}

func main() {
	a := [3]int{1, 2, 3}
	// 直接赋值是值拷贝
	b := a
	// 修改a元素并不影响b
	a[2] = 4
	fmt.Printf("%p, %v\n", &a, a)
	fmt.Printf("%p, %v\n", &b, b)
	// 数组作为函数参数仍然是值拷贝
	f(a)
	c := struct {
		s [3]int
	}{
		s: a,
	}
	// 结构是值拷贝，内部的数组也是值拷贝
	d := c
	// 修改c中的数组元素值并不影响a
	c.s[2] = 30
	// 修改d中的数组元素值并不影响c
	d.s[2] = 20
	fmt.Printf("%p, %v\n", &a, a)
	fmt.Printf("%p, %v\n", &c, c)
	fmt.Printf("%p, %v\n", &d, d)
}

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a []int
	b := make([]int, 0)
	if a == nil {
		fmt.Println("a is nil") // is nil
	} else {
		fmt.Println("a is not nil")
	}
	// 虽然b底层数组大小为0，但切片并不是nil
	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil") // not nil
	}
	// 使用反射中的SliceHeader来获取切片运行时的数据结构
	as := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("len = %d, cap = %d, type = %d\n", len(a), cap(a), as.Data) // len=0, cap=0, type=0
	fmt.Printf("len = %d, cap = %d, type = %d\n", len(b), cap(b), bs.Data) // len=0, cap=0, type=824634236352
}

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	b := a[0:4]
	as := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	// a、b共享底层数组
	fmt.Printf("a = %v, len = %d, cap = %d, type = %d\n", a, len(a), cap(a), as.Data)
	fmt.Printf("b = %v, len = %d, cap = %d, type = %d\n", b, len(b), cap(b), bs.Data)
	b = append(b, 10, 11, 12)
	// a、b继续共享底层数组，修改b会影响共享的底层数组，间接影响a
	fmt.Printf("a = %v, len = %d, cap = %d, type = %d\n", a, len(a), cap(a), as.Data)
	fmt.Printf("b = %v, len = %d, cap = %d, type = %d\n", b, len(b), cap(b), bs.Data)
	// len(b) == 7, 底层数组容量是7，此时需要重新分配数组，并将原来数组值复制到新数组
	b = append(b, 13, 14)
	as = (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs = (*reflect.SliceHeader)(unsafe.Pointer(&b))
	// 可以看到a和b指向底层数组的指针已经不同了
	fmt.Printf("a = %v, len = %d, cap = %d, type = %d\n", a, len(a), cap(a), as.Data)
	fmt.Printf("b = %v, len = %d, cap = %d, type = %d\n", b, len(b), cap(b), bs.Data)
}

package main

import "fmt"

func fa(a int) func(i int) int {
	return func(i int) int {
		fmt.Println(&a, a)
		a = a + i
		return a
	}
}

func main() {
	// f是一个闭包，包括对函数fa形式参数a的“同名引用”
	f := fa(1)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

/*
0xc0000b2010 1
2
0xc0000b2010 2
4
*/

package main

import "fmt"

func main() {
	m := make(map[string]string)
	v, ok := m["some"]
	// 通过ok进行判断
	if !ok {
		fmt.Println("m[some] is nil")
	} else {
		fmt.Println("m[some] is not nil, it is", v)
	}
}

package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		close(c)
	}()
	// for {
	// 	// 使用comma,ok判断通道是否关闭
	// 	v, ok := <-c
	// 	if ok {
	// 		fmt.Println(v)
	// 	} else {
	// 		break
	// 	}
	// }

	// 使用range更加简洁
	for v := range c {
		fmt.Println(v)
	}
}
