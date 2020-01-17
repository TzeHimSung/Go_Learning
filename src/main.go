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
// 	if o, ok := i.(Inter); ok {
// 		o.Ping()
// 		o.Pang()
// 	}

// 	if p, ok := i.(Anter); ok {
// 		// 由于i没有实现接口Anter，所以程序不会执行到这里
// 		p.String()
// 	}

// 	// 判断i绑定的实例是否就是具体类型St
// 	if s, ok := i.(*St); ok {
// 		fmt.Printf("%s", s.Name)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"io"
// )

// func main() {
// 	var i io.Reader
// 	switch v := i.(type) {
// 	case nil:
// 		fmt.Printf("%T\n", v)
// 	default:
// 		fmt.Printf("default")
// 	}
// }

// package main

// import (
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {
// 	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	var i io.Reader = f
// 	switch v := i.(type) {
// 	case io.ReadWriter:
// 		v.Write([]byte("io.ReadWriter\n"))
// 	case *os.File:
// 		v.Write([]byte("*os.File\n"))
// 		v.Sync()
// 	default:
// 		return
// 	}
// }

// package main

// import (
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {
// 	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	var i io.Reader = f
// 	switch v := i.(type) {
// 	case *os.File:
// 		v.Write([]byte("*os.File\n"))
// 		v.Sync()
// 	case io.ReadWriter:
// 		v.Write([]byte("io.ReadWriter\n"))
// 	default:
// 		return
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {
// 	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	var i io.Reader = f
// 	switch v := i.(type) {
// 	case *os.File, io.ReadWriter:
// 		if v == i {
// 			fmt.Println(true)
// 		}
// 	default:
// 		return
// 	}
// }

// package main

// import "fmt"

// // Inter ...
// type Inter interface {
// 	Ping()
// 	Pang()
// }

// // St ...
// type St struct{}

// // Ping ...
// func (St) Ping() {
// 	println("ping")
// }

// // Pang ...
// func (*St) Pang() {
// 	println("pang")
// }

// func main() {
// 	var st *St = nil
// 	var it Inter = st
// 	fmt.Printf("%p\n", st)
// 	fmt.Printf("%p\n", it)
// 	if it != nil {
// 		it.Pang()
// 		// the next centence will lead a panic
// 		// it.Ping()
// 	}
// }

// type itab struct {
// 	inter *interfacetype
// 	_type *_type
// 	hash  uint32 // copy of _type.hash. Used for type switches.
// 	_     [4]byte
// 	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
// }

