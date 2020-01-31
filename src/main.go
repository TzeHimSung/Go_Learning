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
