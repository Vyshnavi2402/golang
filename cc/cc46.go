package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var xtot, ytot, rnutot, rxtot, rytot, xmean, ymean, rxy, n1 float64
	var i int
	var x [20]float64
	var y [20]float64
	var rnu [20]float64
	var rx [20]float64
	var ry [20]float64
	var cv = [11]float64{0.9999, 0.990, 0.959, 0.917, 0.874, 0.834, 0.798, 0.798, 0.765, 0.735, 0.708}
	fmt.Println("Enter the no. of terms ->")
	fmt.Scanln(&n)

	fmt.Println("Enter values of x :::")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of x%d -> ", i+1)
		fmt.Scan(&x[i])
	}

	fmt.Println("Enter values of y :::")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of y%d -> ", i+1)
		fmt.Scan(&y[i])
	}

	xtot = 0
	ytot = 0
	rnutot = 0
	rxtot = 0
	rytot = 0
	//fmt.Println("entered x values are :")
	for i = 0; i < n; i++ {
		xtot = xtot + x[i]
		//fmt.Println(x[i])
	}
	//fmt.Println("entered y values are :")
	for i = 0; i < n; i++ {
		ytot = ytot + y[i]
		//fmt.Println(y[i])
	}
	n1 = float64(n)
	xmean = xtot / n1
	ymean = ytot / n1

	for i = 0; i < n; i++ {
		fmt.Println("/n x-xmean")
		fmt.Println(x[i] - xmean)
	}

	for i = 0; i < n; i++ {
		fmt.Println("/n y-ymean")
		fmt.Println(y[i] - ymean)
	}

	for i = 0; i < n; i++ {
		rnu[i] = (x[i] - xmean) * (y[i] - ymean)
		rx[i] = (x[i] - xmean) * (x[i] - xmean)
		ry[i] = (y[i] - ymean) * (y[i] - ymean)
		fmt.Println("x*y")
		fmt.Println(rnu[i])
		fmt.Println("x2")
		fmt.Println(rx[i])
		fmt.Println("y2")
		fmt.Println(ry[i])
	}

	for i = 0; i < n; i++ {
		rnutot = rnutot + rnu[i]
		rxtot = rxtot + rx[i]
		rytot = rytot + ry[i]
	}

	fmt.Println("xtot=", xtot)
	fmt.Println("ytot=", ytot)
	fmt.Println("rnutot=", rnutot)
	fmt.Println("rdextot=", rxtot)
	fmt.Println("rdeytot=", rytot)
	rxy = rnutot / math.Sqrt(rxtot*rytot)
	fmt.Println("correaltion coefficient", rxy)

	fmt.Println("\n Correlation Coefficient Hypothesis Test")
	fmt.Println("level of significance :::  0.001")
	fmt.Println("Sample size is ", n)
	fmt.Println("Degree of freedom is ", n-2)
	fmt.Println("Critical Value is ", cv[n-2])
	fmt.Println("correlation coefficient", rxy)

}
