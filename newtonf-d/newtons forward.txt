package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	set(n)
}

func set(n int) {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for y := 0; y < n; y++ {
		fmt.Println(a[y])
	}

	b := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scan(&b[j])
	}
	for x := 0; x < n; x++ {
		fmt.Println(b[x])
	}
	c := make([]int, n-1)
	for k := 0; k < n-1; k++ {
		c[k] = b[k+1] - b[k]

	}
	for z := 0; z < n-1; z++ {
		fmt.Println(c[z])
	}
	d := make([]int, n-2)
	for l := 0; l < n-2; l++ {
		d[l] = c[l+1] - c[l]

	}
	for w := 0; w < n-2; w++ {
		fmt.Println(d[w])
	}
	e := make([]int, n-3)
	for m := 0; m < n-3; m++ {
		e[m] = d[m+1] - d[m]

	}
	for v := 0; v < n-3; v++ {
		fmt.Println(e[v])
	}

	var h int = a[1] - a[0]
	fmt.Println(h)
	var s int = 60
	fmt.Println(s)
	var p int = (s - a[0]) / h

	fmt.Println(p)
	return
}
