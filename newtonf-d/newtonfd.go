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
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for y := 0; y < n; y++ {
		fmt.Println(a[y])
	}

	b := make([]float64, n)
	for j := 0; j < n; j++ {
		fmt.Scan(&b[j])
	}
	for x := 0; x < n; x++ {
		fmt.Println(b[x])
	}
	c := make([]float64, n-1)
	for k := 0; k < n-1; k++ {
		c[k] = b[k+1] - b[k]

	}
	for z := 0; z < n-1; z++ {
		fmt.Println(c[z])
	}
	d := make([]float64, n-2)
	for l := 0; l < n-2; l++ {
		d[l] = c[l+1] - c[l]

	}
	for w := 0; w < n-2; w++ {
		fmt.Println(d[w])
	}
	e := make([]float64, n-3)
	for m := 0; m < n-3; m++ {
		e[m] = d[m+1] - d[m]

	}
	for v := 0; v < n-3; v++ {
		fmt.Println(e[v])
	}

	var h float64 = a[1] - a[0]
	fmt.Println(h)
	var s float64
	fmt.Scan(&s)
	fmt.Println(s)
	var p float64 = (s - a[0]) / h

	fmt.Println(p)
	var f float64 = b[0] + p*c[0] + (p*(p-1)*d[0])/2 + (p*(p-1)*(p-2)*e[0])/6

	fmt.Println(f)
	return
}
