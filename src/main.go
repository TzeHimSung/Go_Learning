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

package main

type printer interface {
	print()
}

type s struct{}

func (ss s) print() {
	println("print successfully")
}

func main() {
	var i printer // this is an interface example
	// the following code will cause a panic
	// i.print()
	i = s{}   // an interface example must be inited, or it will cause panic
	i.print() // call function print via interface example
	var a s
	a.print() // call function print via struct example
}
