package main

import (
	"fmt"
	"math"
)

func main() {
	var a [20]float64
	var b [20]float64
	var c [20]float64
	var d [20]float64
	var T float64
	var Tsq float64

	var cf float64
	var ssc float64
	var ssr float64
	var sse float64
	var sscdf float64
	var ssrdf float64
	var ssedf float64
	var msc float64
	var msr float64
	var mse float64
	var sstot float64
	var i = 0
	var n int
	var m int
	var N int
	var fc float64
	var fr float64
	fmt.Println("Enter the no. variety of seed ->")
	fmt.Scan(&n)
	fmt.Println("Enter the no. of varieties of fertilisers ->")
	fmt.Scan(&m)

	fmt.Println("Enter the values of W fertiliser ->")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of production of %d variaty of seed -> ", i+1)
		fmt.Scan(&a[i])
	}
	fmt.Println("Enter the values of X fertilisers ->")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the of production of %d variaty of seed -> ", i+1)
		fmt.Scan(&b[i])
	}
	fmt.Println("Enter the values of Y fertiliser ->")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the of production of %d variaty of seed -> ", i+1)
		fmt.Scan(&c[i])
	}
	fmt.Println("Enter the values of Z fertilisers ->")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of of production of %d variaty of seed -> ", i+1)
		fmt.Scan(&d[i])
	}
	// CALCULATION OF  CORRECTION FACTOR
	for i = 0; i <= n; i++ {
		T = T + (a[i] + b[i] + c[i] + d[i])
	}
	fmt.Print("\nT=", T)

	N = (n * m)

	cf = math.Pow(T, 2) / float64(N)
	fmt.Print("\ncf=", cf)

	//CALCULATION OF SUM OF SQUARES OF DEVIATION FOR TOTAL VARIANCE
	for i = 0; i <= n; i++ {
		Tsq = Tsq + (math.Pow(a[i], 2) + math.Pow(b[i], 2) + math.Pow(c[i], 2) + math.Pow(d[i], 2))
	}
	sstot = Tsq - cf
	fmt.Print("\nss total =", sstot)

	//CALCULATION OF SSC (SUM OF SQUARES BTWN THE COLUMNS)
	ssc = ((math.Pow(a[0]+b[0]+c[0]+d[0], 2) + math.Pow(a[1]+b[1]+c[1]+d[1], 2) + math.Pow(a[2]+b[2]+c[2]+d[2], 2) + math.Pow(a[3]+b[3]+c[3]+d[3], 2)) / float64(m)) - cf
	fmt.Print("\nss between column treatment =", ssc)

	//CALCULATION OF SSR (SUM OF SQUARES BTWN THE ROWS)
	ssr = ((math.Pow(a[0]+a[1]+a[2], 2) + math.Pow(b[0]+b[1]+b[2], 2) + math.Pow(c[0]+c[1]+c[2], 2) + math.Pow(d[0]+d[1]+d[2], 2)) / float64(n)) - cf
	fmt.Print("\nss between row treatment =", ssr)

	//CALCULATION OF SSE (SUM OF SQUARES OF DEVIATION FOR RESIDUAL OR ERROR VARIANCE)
	sse = sstot - (ssc + ssr)
	fmt.Print("\nss residual or error =", sse)

	//degrees of freedom
	sscdf = float64(n - 1)
	ssrdf = float64(m - 1)
	ssedf = float64((n - 1) * (m - 1))
	fmt.Print("\nss between columns degree of freedom =", sscdf)
	fmt.Print("\nss between rows degree of freedom =", ssrdf)
	fmt.Print("\nss of errors degree of freedom =", ssedf)

	//mean squares
	msc = ssc / sscdf
	msr = ssr / ssrdf
	mse = sse / ssedf
	fmt.Print("\nmean square between columns  =", msc)
	fmt.Print("\nmean square between rows  =", msr)
	fmt.Print("\nmean square of residual or error  =", mse)

	//F-Ratios
	fc = msc / mse
	fr = msr / mse
	fmt.Print("\nf ratio of between columns is =", fc)
	fmt.Print("\nf ratio of between rows is =", fr)
}
