package main

import (
	"fmt"
	"math"
)

func main() {
	var a [10]float64
	var b [10]float64
	var c [10]float64
	var d [10]float64
	var s [10]float64

	var ms float64

	var ssb float64
	var ssw float64

	var ssbdf float64
	var sswdf float64

	var msb float64
	var msw float64

	var ss float64
	var f float64
	var i int
	var n int

	a[0] = 56
	a[1] = 43
	a[2] = 69
	b[0] = 52
	b[1] = 47
	b[2] = 70
	c[0] = 58
	c[1] = 69
	c[2] = 72
	d[0] = 54
	d[1] = 65
	d[2] = 68
	n = 3

	s[0] = (a[0] + b[0] + c[0] + d[0]) / 4
	s[1] = (a[1] + b[1] + c[1] + d[1]) / 4
	s[2] = (a[2] + b[2] + c[2] + d[2]) / 4

	ms = (s[0] + s[1] + s[2]) / 3
	fmt.Print("ms=", ms)

	for i = 0; i < n; i++ {
		ssb = ssb + (4 * math.Pow(s[i]-ms, 2))
	}
	fmt.Print("ssb =", ssb)

	for i = 0; i < n; i++ {
		ssw = ssw + (math.Pow(a[i]-s[i], 2) + math.Pow(b[i]-s[i], 2) + math.Pow(c[i]-s[i], 2) + math.Pow(d[i]-s[i], 2))
	}
	fmt.Println("ssw =", ssw)

	ss = ssb + ssw
	fmt.Print("ss =", ss)

	ssbdf = float64(n - 1)
	sswdf = float64(12 - n)
	fmt.Print("ssbdf =", ssbdf)
	fmt.Print("sswdf =", sswdf)

	msb = (ssb / ssbdf)
	msw = ssw / sswdf
	fmt.Print("msb =", msb)
	fmt.Println("msw=", msw)

	f = msb / msw
	fmt.Print("f ratio is =", f)

}
