package main

import (
	"fmt"
	"math"
)

var i uint64 = 1

var m uint64
var fact uint64 = 1

func main() {
	var n, r, nr, factn, factr, factnr, p, result uint64

	fmt.Printf("binomial distribution\n")
	fmt.Printf("enter the value of n:")
	fmt.Scanf("%d", &n)
	fmt.Printf("enter the value of r:")
	fmt.Scanf("%d", &r)
	fmt.Printf("enter the value of p:")
	fmt.Scanf("%lf", &p)

	nr = n - r
	fmt.Printf("Factorial is", facto(2))
	factn = facto(n)
	factr = facto(r)
	factnr = facto(nr)
	result = (factn * math.Pow(p, r) * (math.Pow((1 - p), n-r))) / (factr * factnr)
	fmt.Printf("\n result=%.3lf", result)

}
func facto(m uint64) uint64 {

	for i := 1; i <= m; i++ {
		fact = fact * uint64(i)
	}

	return fact
}
