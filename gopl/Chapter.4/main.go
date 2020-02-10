// package main

// import "fmt"

// func main() {
// 	var a [3][3]int = [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
// 	for _, i := range a {
// 		for _, val := range i {
// 			fmt.Printf("%d ", val)
// 		}
// 		fmt.Println()
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	a := [2]int{1, 2}
// 	b := [...]int{1, 2}
// 	c := [2]int{1, 3}
// 	fmt.Println(a == b, a == c, b == c)
// }

// package main

// import (
// 	"crypto/sha256"
// 	"fmt"
// )

// func main() {
// 	c1 := sha256.Sum256([]byte("x"))
// 	c2 := sha256.Sum256([]byte("X"))
// 	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
// }

// package main

// import "fmt"

// func reverse(s []int) {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }

// func main() {
// 	a := [3]int{0, 1, 2}
// 	s := a[:]
// 	reverse(s) // 这里居然是引用？是的，因为slice使用a做底层数组，存在一个间接引用
// 	fmt.Println(a)
// }

// package main

// import "fmt"

// func main() {
// 	s := "this is just a test"
// 	ss := []rune(s)
// 	for _, v := range ss {
// 		fmt.Printf("%c", v)
// 	}
// }

// package main

// import "fmt"

// func nonempty(strings []string) []string {
// 	i := 0
// 	for _, s := range strings {
// 		if s != "" {
// 			strings[i] = s
// 			i++
// 		}
// 	}
// 	return strings[:i]
// }

// func main() {
// 	fmt.Println(nonempty([]string{"a", "b", "c"}))
// }

// package main

// import "fmt"

// func remove(slice []int, i int) []int {
// 	copy(slice[i:], slice[i+1:])
// 	return slice[:len(slice)-1]
// }

// func main() {
// 	s := []int{5, 6, 7, 8, 9}
// 	fmt.Println(remove(s, 2))
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	var name []string = []string{"charlie", "alice"}
// 	ages := map[string]int{
// 		"alice":   13,
// 		"charlie": 34,
// 	}
// 	fmt.Println(ages)
// 	delete(ages, "alice")
// 	fmt.Println(ages)
// 	ages["alice"] = 13
// 	sort.Strings(name)
// 	for _, val := range name {
// 		fmt.Print(ages[val], " ")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var ages map[string]int = make(map[string]int)
// 	ages["carol"] = 21
// 	fmt.Println(ages)
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	seen := make(map[string]bool)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		line := input.Text()
// 		if !seen[line] {
// 			seen[line] = true
// 			fmt.Println(line)
// 		}
// 	}
// 	if err := input.Err(); err != nil {
// 		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
// 		os.Exit(1)
// 	}
// }

// package main

// import "fmt"

// var m = make(map[string]int)

// func k(list []string) string {
// 	return fmt.Sprintf("%q", list)
// }

// // Add .
// func Add(list []string) {
// 	m[k(list)]++
// }

// // Count .
// func Count(list []string) int {
// 	return m[k(list)]
// }

// func main() {
// 	a := []string{"a", "b", "c"}
// 	aa := []string{"aa", "bb", "cc"}
// 	Add(a)
// 	fmt.Println(Count(a))
// 	fmt.Println(Count(aa))
// 	Add(aa)
// 	fmt.Println(Count(aa))
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"unicode"
// 	"unicode/utf8"
// )

// func main() {
// 	counts := make(map[rune]int)
// 	var utflen [utf8.UTFMax + 1]int
// 	invalid := 0

// 	in := bufio.NewReader(os.Stdin)
// 	for {
// 		r, n, err := in.ReadRune()
// 		if err != io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
// 			os.Exit(1)
// 		}
// 		if r == unicode.ReplacementChar && n == 1 {
// 			invalid++
// 			continue
// 		}
// 		counts[r]++
// 		utflen[n]++
// 	}
// 	fmt.Printf("rune\tcount\n")
// 	for c, n := range counts {
// 		fmt.Printf("%q\t%d\n", c, n)
// 	}
// 	fmt.Printf("\nlen\tcount\n")
// 	for i, n := range utflen {
// 		if i > 0 {
// 			fmt.Printf("%d\t%d\n", i, n)
// 		}
// 	}
// 	if invalid > 0 {
// 		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// // Employee .
// type Employee struct {
// 	ID            int
// 	Name, Address string
// 	DoB           time.Time
// 	Position      string
// 	Salary        int
// 	ManagerID     int
// }

// var dilbert Employee = Employee{
// 	ID:        10,
// 	Name:      "xiaohong",
// 	Address:   "guangzhou",
// 	Position:  "guangzhou",
// 	Salary:    45000,
// 	ManagerID: 2,
// }

// func (emp Employee) getAddress() *Employee {
// 	return &emp
// }

// func (emp Employee) getIDAddress() *int {
// 	return &emp.ID
// }

// func main() {
// 	ptr := &dilbert
// 	fmt.Println(dilbert)
// 	fmt.Println(&ptr)
// 	fmt.Println(ptr.getIDAddress())
// 	fmt.Println(dilbert.getIDAddress())
// }

// /*
// 利用二叉树中序遍历实现排序
// */
// package main

// import "fmt"

// type node struct {
// 	value       int
// 	left, right *node
// }

// func sort(values []int) {
// 	var root *node
// 	for _, v := range values {
// 		root = add(root, v)
// 		appendValues(values[:0], root)
// 	}
// }

// func appendValues(values []int, t *node) []int {
// 	if t != nil {
// 		values = appendValues(values, t.left)
// 		values = append(values, t.value)
// 		values = appendValues(values, t.right)
// 	}
// 	fmt.Println(values)
// 	return values
// }

// func add(t *node, value int) *node {
// 	// if meet leaf node
// 	if t == nil {
// 		t = new(node)
// 		t.value = value
// 		return t
// 	}
// 	// if not left, compare value
// 	if value < t.value {
// 		t.left = add(t.left, value)
// 	} else {
// 		t.right = add(t.right, value)
// 	}
// 	return t
// }

// func main() {
// 	val := []int{10, 9, 2, 4, 1, 5, 6}
// 	sort(val)
// 	fmt.Println(val)
// }

// package main

// import "fmt"

// // Point .
// type Point struct{ X, Y int }

// func main() {
// 	p := Point{1, 2}
// 	fmt.Println(p)
// }

// package main

// import "fmt"

// // Point .
// type Point struct {
// 	X, Y int
// }

// // Circle .
// type Circle struct {
// 	Point
// 	Radius int
// }

// // Wheel .
// type Wheel struct {
// 	Circle
// 	Spokes int
// }

// func main() {
// 	// compile error
// 	// var w Wheel = Wheel{8, 8, 5, 20}
// 	// var w Wheel = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20}
// 	// following is right
// 	var w Wheel = Wheel{Circle{Point{8, 8}, 5}, 20}
// 	w = Wheel{
// 		Circle: Circle{
// 			Point:  Point{X: 8, Y: 8},
// 			Radius: 5,
// 		},
// 		Spokes: 20,
// 	}
// 	fmt.Println(w)
// 	// main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
// 	fmt.Printf("%#v\n", w)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// // Movie .
// type Movie struct {
// 	Title  string
// 	Year   int  `json:"released"`
// 	Color  bool `json:"color,omitempty"`
// 	Actors []string
// }

// var movies = []Movie{
// 	{Title: "Casablanca", Year: 1942, Color: false,
// 		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
// 	{Title: "Cool Hand Luke", Year: 1967, Color: true,
// 		Actors: []string{"Paul Newman"}},
// 	{Title: "Bullitt", Year: 1968, Color: true,
// 		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
// }

// func main() {
// 	// data, err := json.Marshal(movies)
// 	data, err := json.MarshalIndent(movies, "", " ")
// 	if err != nil {
// 		log.Fatalf("JSON marshaling failed: %s", err)
// 	}
// 	fmt.Printf("%s\n", data)
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// IssuesURL .
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult .
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue .
type Issue struct {
	Number       int
	HTMLURL      string `json:"html_url"`
	Title, State string
	User         *User
	CreatedAt    time.Time `json:"created_at"`
	Body         string
}

// User .
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
