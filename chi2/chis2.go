package main

import (
	"fmt"
)

func main() {
	a := [2][2]float32{{23, 28}, {6, 4}}
	var r1 float32
	var c1 float32
	var r2 float32
	var c2 float32
	var gt float32
	var ef11 float32
	var ef12 float32
	var ef21 float32
	var ef22 float32

	fmt.Println("Observed frequency")
	fmt.Println(a)
	r1 = a[0][0] + a[0][1]
	c1 = a[0][0] + a[1][0]
	r2 = a[1][0] + a[1][1]
	c2 = a[0][1] + a[1][1]
	fmt.Println("sum of row 1 ", r1)
	fmt.Println("sum of row 2 ", r2)
	fmt.Println("sum of col 1 ", c1)
	fmt.Println("sum of col 2 ", c2)

	gt = r1 + r2
	fmt.Println("grand total is ::", gt)
	ef11 = (r1 * c1) / gt
	ef12 = (r1 * c2) / gt
	ef21 = (r2 * c1) / gt
	ef22 = (r2 * c2) / gt
	b := [2][2]float32{{ef11, ef12}, {ef21, ef22}}

	fmt.Println("Expected frequency")
	fmt.Println(b)

}
