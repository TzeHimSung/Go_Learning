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

// package main

// import "fmt"

// func main() {
// 	ch := make(chan int, 1)
// 	go func(chan int) {
// 		for {
// 			select {
// 			case ch <- 0:
// 			case ch <- 1:
// 			}
// 		}
// 	}(ch)
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(<-ch)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"runtime"
// )

// func generateIntA(done chan struct{}) chan int {
// 	ch := make(chan int)
// 	go func(chan int) {
// 	Lable:
// 		for {
// 			select {
// 			case ch <- rand.Int():
// 			case <-done:
// 				break Lable
// 			}
// 		}
// 		close(ch)
// 	}(ch)
// 	return ch
// }

// func main() {
// 	done := make(chan struct{})
// 	ch := generateIntA(done)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	close(done)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println("NumGoroutine =", runtime.NumGoroutine())
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func generateIntA() chan int {
// 	ch := make(chan int, 10)
// 	// 启动一个goroutine用于生成随机数，函数返回一个通道用于获取随机数
// 	go func() {
// 		for {
// 			ch <- rand.Int()
// 		}
// 	}()
// 	return ch
// }

// func generateIntB() chan int {
// 	ch := make(chan int, 10)
// 	go func() {
// 		for {
// 			ch <- rand.Int()
// 		}
// 	}()
// 	return ch
// }

// func generateInt() chan int {
// 	ch := make(chan int, 10)
// 	go func() {
// 		for {
// 			// 使用select的扇入技术来增加生成的随机源
// 			select {
// 			case ch <- <-generateIntA():
// 			case ch <- <-generateIntB():
// 			}
// 		}
// 	}()
// 	return ch
// }

// func main() {
// 	ch := generateInt()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(<-ch)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func generateIntA(done chan struct{}) chan int {
// 	ch := make(chan int)
// 	go func() {
// 	Lable:
// 		for {
// 			// 通过select监听一个信号chan来确定是否停止生成
// 			select {
// 			case ch <- rand.Int():
// 			case <-done:
// 				break Lable
// 			}
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }

// func main() {
// 	done := make(chan struct{})
// 	ch := generateIntA(done)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	close(done)
// 	for v := range ch {
// 		fmt.Print(v)
// 	}
// }

package main

import "math/rand"

import "fmt"

func generateIntA(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func generateIntB(done chan struct{}) chan int {
	ch := make(chan int, 10)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func generateInt(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})
	go func() {
	Lable:
		for {
			select {
			case ch <- <-generateIntA(send):
			case ch <- <-generateIntB(send):
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := generateInt(done)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	done <- struct{}{}
	fmt.Println("stop generate")
}
