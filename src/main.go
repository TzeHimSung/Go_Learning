// package main

// // Reader ...
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// // Writer ...
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// type readAndWrite interface {
// 	Reader
// 	Writer
// }

// type readAndWrite2 interface {
// 	Read(p []byte) (n int, err error)
// 	Writer
// }

// package main

// type printer interface {
// 	print()
// }

// type s struct{}

// func (ss s) print() {
// 	println("print successfully")
// }

// func main() {
// 	var i printer // this is an interface example
// 	// the following code will cause a panic
// 	// i.print()
// 	i = s{}   // an interface example must be inited, or it will cause panic
// 	i.print() // call function print via interface example
// 	var a s
// 	a.print() // call function print via struct example
// }

// package main

// import "fmt"

// // Inter ...
// type Inter interface {
// 	Ping()
// 	Pang()
// }

// // Anter ...
// type Anter interface {
// 	Inter
// 	String()
// }

// // St ...
// type St struct {
// 	Name string
// }

// // Ping ...
// func (St) Ping() {
// 	println("ping")
// }

// // Pang ...
// func (*St) Pang() {
// 	println("pang")
// }

// func main() {
// 	st := &St{"andes"}
// 	var i interface{} = st
// 	// 判断i绑定的实例是否实现了接口类型Inter
// 	o := i.(Inter)
// 	o.Ping()
// 	o.Pang()
// 	// 如下语句会引发panic，因为i没有实现接口Anter
// 	// p := i.(Anter)
// 	// p.String()
// 	// 判断i绑定的实例是否就是具体类型St
// 	s := i.(*St)
// 	fmt.Printf("%s", s.Name)
// }

package main

import "fmt"

// Inter ...
type Inter interface {
	Ping()
	Pang()
}

// Anter ...
type Anter interface {
	Inter
	String()
}

// St ...
type St struct {
	Name string
}

// Ping ...
func (St) Ping() {
	println("ping")
}

// Pang ...
func (*St) Pang() {
	println("pang")
}

func main() {
	st := &St{"andes"}
	var i interface{} = st
	// 判断i绑定的实例是否实现了接口类型Inter
	if o, ok := i.(Inter); ok {
		o.Ping()
		o.Pang()
	}

	if p, ok := i.(Anter); ok {
		// 由于i没有实现接口Anter，所以程序不会执行到这里
		p.String()
	}

	// 判断i绑定的实例是否就是具体类型St
	if s, ok := i.(*St); ok {
		fmt.Printf("%s", s.Name)
	}
}
