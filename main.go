// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	fmt.Println(strings.Join(os.Args[1:], ""))
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	// note: ignoring potential error from input.Err()
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		countLines(os.Stdin, counts)
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 			}
// 			countLines(f, counts)
// 			f.Close()
// 		}
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strings"
// )

// func main() {
// 	counts := make(map[string]int)
// 	for _, filename := range os.Args[1:] {
// 		data, err := ioutil.ReadFile(filename)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
// 			continue
// 		}
// 		for _, line := range strings.Split(string(data), "\n") {
// 			counts[line]++
// 		}
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// package main

// import (
// 	"image"
// 	"image/color"
// 	"image/gif"
// 	"io"
// 	"math"
// 	"math/rand"
// 	"os"
// 	"time"
// )

// var palette = []color.Color{color.White, color.Black}

// const (
// 	whiteIndex = 0
// 	blackIndex = 1
// )

// func main() {
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	lissajous(os.Stdout)
// }

// func lissajous(out io.Writer) {
// 	const (
// 		cycles  = 5
// 		res     = 0.001
// 		size    = 100
// 		nframes = 64
// 		delay   = 8
// 	)
// 	freq := rand.Float64() * 3.0
// 	anim := gif.GIF{LoopCount: nframes}
// 	phase := 0.0
// 	for i := 0; i < nframes; i++ {
// 		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
// 		img := image.NewPaletted(rect, palette)
// 		for t := 0.0; t < cycles*2*math.Pi; t += res {
// 			x := math.Sin(t)
// 			y := math.Sin(t*freq + phase)
// 			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
// 		}
// 		phase += 0.1
// 		anim.Delay = append(anim.Delay, delay)
// 		anim.Image = append(anim.Image, img)
// 	}
// 	gif.EncodeAll(out, &anim)
// }

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	for _, url := range os.Args[1:] {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}
// 		b, err := ioutil.ReadAll(resp.Body)
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("%s", b)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"time"
// )

// func main() {
// 	start := time.Now()
// 	ch := make(chan string)
// 	for _, url := range os.Args[1:] {
// 		go fetch(url, ch)
// 	}
// 	for range os.Args[1:] {
// 		fmt.Println(<-ch)
// 	}
// 	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
// }

// func fetch(url string, ch chan<- string) {
// 	start := time.Now()
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		ch <- fmt.Sprint(err)
// 		return
// 	}
// 	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
// 	resp.Body.Close()
// 	if err != nil {
// 		ch <- fmt.Sprintf("while reading %s: %v", url, err)
// 		return
// 	}
// 	secs := time.Since(start).Seconds()
// 	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe("localhost:8000", nil))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"sync"
// )

// var (
// 	mu    sync.Mutex
// 	count int
// )

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.HandleFunc("/count", counter)
// 	log.Fatal(http.ListenAndServe("localhost:8000", nil))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	count++
// 	mu.Unlock()
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

// func counter(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	fmt.Fprintf(w, "Count %d\n", count)
// 	mu.Unlock()
// }

// package main

// import "fmt"

// const boilingF = 212.0

// func main() {
// 	var f = boilingF
// 	var c = (f - 35) * 5 / 9
// 	fmt.Printf("boiling point = %gF or %gC\n", f, c)
// }

package main

import "fmt"

func main() {
	var p1, p2 = getV(), getV()
	fmt.Println(p1, p2, p1 == p2)
}

func getV() *int {
	value := 100
	return &value
}
