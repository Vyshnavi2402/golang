package main

import (
	"fmt"
)

func main() {
	var n int

	var px [5]int
	var py [5]int
	var qx [5]int
	var qy [5]int
	var rx [5]int
	var ry [5]int
	fmt.Println("no of sets : n =")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println(" \nfor i=  ", i)

		fmt.Println("px")
		fmt.Scan(&px[i])
		fmt.Println("py")
		fmt.Scan(&py[i])
		fmt.Println("mx")
		fmt.Scan(&qx[i])
		fmt.Println("qy")
		fmt.Scan(&qy[i])

		rx[i] = 2*qx[i] - px[i]
		ry[i] = 2*qy[i] - py[i]

		fmt.Print("rx[", i, "]", rx[i])
		fmt.Print("\nry[", i, "]", ry[i])
	}
}
