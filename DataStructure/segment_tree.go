/*
segment tree
implementation with golang
including signal point fix / query of maxinum element in specific interval
*/
package main

import "fmt"

// Node node of segment tree
type Node struct {
	val int
}

const maxn int = 1e5 + 10

var segt [maxn]Node
var a [maxn]int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// maintain node of segment tree
func maintain(curpos int) {
	segt[curpos].val = max(segt[curpos*2].val, segt[curpos*2+1].val)
}

// build segment tree
func build(curpos, curl, curr int) {
	if curl == curr {
		segt[curpos].val = a[curl]
		return
	}
	mid := (curl + curr) / 2
	build(curpos*2, curl, mid)
	build(curpos*2+1, mid+1, curr)
	maintain(curpos)
}

// update value of specific position
func update(curpos, pos, curl, curr, val int) {
	if curl == curr {
		segt[curpos].val = val
		return
	}
	mid := (curl + curr) / 2
	if pos <= mid {
		update(curpos*2, pos, curl, mid, val)
	} else {
		update(curpos, pos, mid+1, curr, val)
	}
	maintain(curpos)
}

// query maxinum of specific interal
func queryMax(curpos, curl, curr, ql, qr int) (ret int) {
	if ql <= curl && curr <= qr {
		return segt[curpos].val
	}
	mid := (curl + curr) / 2
	if ql <= mid {
		ret = max(ret, queryMax(curpos*2, curl, mid, ql, qr))
	}
	if qr > mid {
		ret = max(ret, queryMax(curpos*2+1, mid+1, curr, ql, qr))
	}
	return
}

func main() {
	var n int = 0
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	build(1, 1, n)

	var q int = 0
	fmt.Scanf("%d", &q)
	for cnt := 1; cnt <= q; cnt++ {
		var op int = 0
		var arg1 int = 0
		var arg2 int = 0
		fmt.Scanf("%d %d %d", &op, &arg1, &arg2)
		if op == 1 {
			update(1, arg1, 1, n, arg2)
		} else {
			fmt.Println(queryMax(1, 1, n, arg1, arg2))
		}
	}
}
