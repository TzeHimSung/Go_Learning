// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func main() {
// 	go func() {
// 		sum := 0
// 		for i := 0; i < 1000; i++ {
// 			sum += i
// 		}
// 		fmt.Println(sum)
// 		time.Sleep(1 * time.Second)
// 	}()
// 	// NumGoroutine可以返回当前程序的goroutine数目
// 	fmt.Println("NumGoroutine =", runtime.NumGoroutine())
// 	// main goroutine故意sleep几秒防止提前退出
// 	time.Sleep(3 * time.Second)
// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func sum() {
// 	sum := 0
// 	for i := 0; i < 1000; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// 	time.Sleep(1 * time.Second)
// }

// func main() {
// 	go sum()
// 	fmt.Println("NumGoroutine =", runtime.NumGoroutine())
// 	time.Sleep(3 * time.Second)
// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	// 获取当前的GOMAXPROCS值
// 	fmt.Println("GOMAXPROCS =", runtime.GOMAXPROCS(0))
// 	// 设置GOMAXPROCS的值为2
// 	runtime.GOMAXPROCS(2)
// 	// 获取当前的GOMAXPROCS值
// 	fmt.Println("GOMAXPROCS =", runtime.GOMAXPROCS(0))
// }

package main

func main() {
	// 创建一个没有缓冲区的通道，通道存放元素类型为datatype
	make(chan datatype)
	// 创建一个缓冲区大小为10的通道，通道存放元素类型为datatype
	make(chan datatype2, 10)
}
