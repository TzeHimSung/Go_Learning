// package main

// import (
// 	"fmt"
// 	"math"
// )

// func hypot(x, y float64) (ret float64) {
// 	ret = math.Sqrt(x*x + y*y)
// 	return
// }

// func main() {
// 	fmt.Println(hypot(3, 4))
// }

// package main

// import (
// 	"fmt"
// 	"os"

// 	"golang.org/x/net/html"
// )

// func main() {
// 	doc, err := html.Parse(os.Stdin)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
// 		os.Exit(1)
// 	}
// 	for _, link := range visit(nil, doc) {
// 		fmt.Println(link)
// 	}
// }

// func visit(links []string, n *html.Node) []string {
// 	if n.Type == html.ElementNode && n.Data == "a" {
// 		for _, a := range n.Attr {
// 			if a.Key == "href" {
// 				links = append(links, a.Val)
// 			}
// 		}
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		links = visit(links, c)
// 	}
// 	return links
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"golang.org/x/net/html"
// )

// func findLinks(url string) ([]string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		resp.Body.Close()
// 		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
// 	}
// 	doc, err := html.Parse(resp.Body)
// 	resp.Body.Close()
// 	if err != nil {
// 		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
// 	}
// 	return visit(nil, doc), nil
// }

// func visit(links []string, n *html.Node) []string {
// 	if n.Type == html.ElementNode && n.Data == "a" {
// 		for _, a := range n.Attr {
// 			if a.Key == "href" {
// 				links = append(links, a.Val)
// 			}
// 		}
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		links = visit(links, c)
// 	}
// 	return links
// }

// func main() {
// 	for _, url := range os.Args[1:] {
// 		links, err := findLinks(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
// 			continue
// 		}
// 		for _, link := range links {
// 			fmt.Println(link)
// 		}
// 	}
// }

package main

import "fmt"

func get() *int {
	a := 10
	return &a
}

func main() {
	b := get()
	fmt.Println(b, *b)
}
