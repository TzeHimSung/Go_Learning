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

// package main

// func main() {
// 	// 创建一个没有缓冲区的通道，通道存放元素类型为datatype
// 	make(chan datatype)
// 	// 创建一个缓冲区大小为10的通道，通道存放元素类型为datatype
// 	make(chan datatype2, 10)
// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	c := make(chan struct{})
// 	go func(i chan struct{}) {
// 		sum := 0
// 		for i := 0; i < 10000; i++ {
// 			sum += i
// 		}
// 		fmt.Println(sum)
// 		// write chan
// 		c <- struct{}{}
// 	}(c)
// 	fmt.Println("NumGoroutine =", runtime.NumGoroutine())
// 	// read chan c
// 	<-c
// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	c := make(chan struct{})
// 	ci := make(chan int, 100)
// 	go func(i chan struct{}, j chan int) {
// 		for i := 0; i < 10; i++ {
// 			ci <- i
// 		}
// 		close(ci)
// 		// write channel
// 		c <- struct{}{}
// 	}(c, ci)
// 	fmt.Println("NumGoroutine =", runtime.NumGoroutine())
// 	<-c
// 	fmt.Println("NumGoroutine =", runtime.NumGoroutine())
// 	for v := range ci {
// 		fmt.Println(v)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// var wg sync.WaitGroup

// var urls = []string{
// 	"http://www.baidu.com/",
// 	"http://cn.bing.com/",
// 	"http://www.qq.com/",
// }

// func main() {
// 	for _, url := range urls {
// 		// 每一个url启动一个goroutine，同时给wg加1
// 		wg.Add(1)
// 		go func(url string) {
// 			// 当前goroutine结束后给wg计数减1，wg.Done()等价于wg.Add(-1)
// 			defer wg.Done()
// 			// 发送http get请求并打印http返回码
// 			resp, err := http.Get(url)
// 			if err == nil {
// 				fmt.Println(url, resp.Status)
// 			}
// 		}(url)
// 	}
// 	// 等待所有请求结束
// 	wg.Wait()
// }

package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	go func(chan int) {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
