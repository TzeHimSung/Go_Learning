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

