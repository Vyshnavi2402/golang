package main

// Least square curve fitting code
import "fmt"

func main() {
	var n int
	var xtot, ytot, xsqtot, xytot float64
	var i int
	var x [20]float64
	var y [20]float64
	var xsq [20]float64
	var xy [20]float64

	fmt.Println("Enter the no. of terms ->")
	fmt.Scan(&n)

	fmt.Println("Enter the value in the form of x ->")

	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of x%d -> ", i+1)
		fmt.Scan(&x[i])
	}

	fmt.Println("Enter y values")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of y%d -> ", i+1)
		fmt.Scan(&y[i])
	}
	fmt.Println("Entered x values are")
	xtot = 0
	ytot = 0
	xsqtot = 0
	xytot = 0

	for i = 0; i < n; i++ {
		xtot = xtot + x[i]
		fmt.Println(x[i])
	}
	fmt.Println("Entered y values are")
	for i = 0; i < n; i++ {
		ytot = ytot + y[i]
		fmt.Println(y[i])
	}
	fmt.Println("Entered sq of x values are")
	for i = 0; i < n; i++ {
		xsq[i] = x[i] * x[i]
		fmt.Println(xsq[i])
	}
	fmt.Println("Entered x*y values are")
	for i = 0; i < n; i++ {
		xy[i] = x[i] * y[i]
		fmt.Println(xy[i])

	}
	for i = 0; i < n; i++ {
		xytot = xytot + xy[i]
		xsqtot = xsqtot + xsq[i]

	}

	fmt.Println("Xtot=", xtot)
	fmt.Println("Ytot=", ytot)
	fmt.Println("xsqtot=", xsqtot)
	fmt.Println("xytot=", xytot)

	var a1 float64
	var a2 float64
	var det float64
	var a1den float64
	var a2den float64
	var n1 float64
	n1 = float64(n)
	det = n1*xsqtot - xtot*xtot
	a1den = xtot*(-xytot) - xsqtot*(-ytot)
	a1 = a1den / det
	a2den = (-ytot)*(xtot) - (-xytot)*(n1)
	a2 = a2den / det
	fmt.Println(a1)
	fmt.Println(a2)

	y[2] = a1 + (a2 * x[2])
	fmt.Println("Predicted value of y at given x[2] is", y[2])

}
