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
// package main

// // Needs to be in sync with ../cmd/link/internal/ld/decodesym.go:/^func.commonsize,
// // ../cmd/compile/internal/gc/reflect.go:/^func.dcommontype and
// // ../reflect/type.go:/^type.rtype.
// type _type struct {
// 	size       uintptr  // 大小
// 	ptrdata    uintptr  // size of memory prefix holding all pointers
// 	hash       uint32   // 类型hash
// 	tflag      tflag    // 类型的特征标记
// 	align      uint8    // _type作为整体变量存放时的对齐字节数
// 	fieldalign uint8    // 当前结构字段的对齐字节数
// 	kind       uint8    // 基础类型枚举值和反射中的kind一致，kind决定了如何解析该类型
// 	alg        *typeAlg // 指向一个函数指针表，该表有两个函数，一个是计算类型hash函数，另一个是比较两个类型是否相同的equal函数
// 	// gcdata stores the GC type data for the garbage collector.
// 	// If the KindGCProg bit is set in kind, gcdata is a GC program.
// 	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
// 	gcdata    *byte   // GC相关信息
// 	str       nameOff // str用来表示类型名称字符串在编译后二进制文件中某个section的偏移量
// 	ptrToThis typeOff // ptrToThis用来表示类型元信息的指针在编译后二进制文件中某个section的偏移量，由链接器负责填充
// }

// type imethod struct {
// 	name nameOff  // 方法名在编译后的section里面的偏移量
// 	ityp typeOff  /// 方法类型在编译后的section里面的偏移量
// }

// type interfacetype struct {
// 	typ     _type // 类型通用部分
// 	pkgpath name // 接口所属包的名字信息，name内存放的不仅有名称，还有描述信息
// 	mhdr    []imethod // 接口的方法
// }

package main

import "fmt"

// Caler ...
type Caler interface {
	Add(a, b int) int
	Sub(a, b int) int
}

// Adder ...
type Adder struct {
	id int
}

// Add ...
func (adder Adder) Add(a, b int) int {
	return a + b
}

// Sub ...
func (adder Adder) Sub(a, b int) int {
	return a - b
}

func main() {
	var m Caler = Adder{id: 1234}
	fmt.Println(m.Add(1, 2))
}
