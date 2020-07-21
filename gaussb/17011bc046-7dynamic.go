package main

import (
	"fmt"
)

func main() {

	var ax [10]float32
	var ay [10]float32
	var x float32
	var y float32 = 0
	var h float32
	var p float32
	var n int
	var diff [20][20]float32
	var y1, y2, y3, y4 float32

	fmt.Println(" GAUSS BACKWARD INTERPOLATION ")
	fmt.Println("Enter the no. of terms  : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of x%d : ", i+1)
		fmt.Scan(&ax[i])
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of y%d :  ", i+1)
		fmt.Scan(&ay[i])
	}

	fmt.Println("Enter the value of x for which y is to be solved  :")

	fmt.Scan(&x)

	h = ax[1] - ax[0]
	for i := 0; i < n-1; i++ {
		diff[i][1] = ay[i+1] - ay[i]
	}
	for j := 2; j <= 4; j++ {

		for i := 0; i < n-j; i++ {
			diff[i][j] = diff[i+1][j-1] - diff[i][j-1]
		}
	}

	for i := 0; ax[i] < x; i++ {

		p = (x - ax[i]) / h
		y1 = p * diff[i][1]
		y2 = p * (p + 1) * diff[i][2] / 2
		y3 = (p + 1) * p * (p - 1) * diff[i][3] / 6
		y4 = (p + 1) * p * (p - 1) * (p + 2) * diff[i][4] / 24
		y = ay[i] + y1 + y2 + y3 + y4

		fmt.Println("the value of h : ", h)
		fmt.Println("the value of p : ", p)
		fmt.Printf(" When x= %g", x)
		fmt.Printf("\ny= %g", y)
		break

	}

}
