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
	var Tj float64

	var cf float64
	var ssc float64
	var sse float64
	var sscdf float64
	var ssedf float64
	var msc float64
	var mse float64
	var sstot float64
	var i = 0
	var n int
	var m int
	var N int
	var f float64
	fmt.Println("Enter the no. of years ->")
	fmt.Scan(&n)
	fmt.Println("Enter the no. of subjects ->")
	fmt.Scan(&m)

	fmt.Println("Enter the values of subject computer science  average marks ->")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of cs %d year -> ", i+1)
		fmt.Scan(&a[i])
	}
	fmt.Println("Enter the values of subject design average marks ->")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of design %d year -> ", i+1)
		fmt.Scan(&b[i])
	}
	fmt.Println("Enter the values of subject multimedia average marks ->")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of mult %d year -> ", i+1)
		fmt.Scan(&c[i])
	}
	fmt.Println("Enter the values of subject gis average marks ->")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the value of gis %d year -> ", i+1)
		fmt.Scan(&d[i])
	}

	for i = 0; i <= n; i++ {
		T = T + (a[i] + b[i] + c[i] + d[i])
	}
	fmt.Print("\nT=", T)

	N = (n * m)

	cf = math.Pow(T, 2) / float64(N)
	fmt.Print("\ncf=", cf)

	for i = 0; i <= n; i++ {
		Tsq = Tsq + (math.Pow(a[i], 2) + math.Pow(b[i], 2) + math.Pow(c[i], 2) + math.Pow(d[i], 2))
	}
	sstot = Tsq - cf
	fmt.Print("\nss total =", sstot)

	for i = 0; i <= n; i++ {

		Tj = Tj + (math.Pow(a[i]+b[i]+c[i]+d[i], 2))
	}
	fmt.Print("\nTj square=", Tj)

	ssc = (Tj / float64(m)) - cf
	fmt.Print("\nss between =", ssc)

	sse = sstot - ssc
	fmt.Print("\nss within=", sse)

	sscdf = float64(n - 1)
	msc = (ssc / sscdf)
	fmt.Print("\n\nms between =", msc)
	ssedf = float64(12 - n)
	mse = sse / ssedf
	fmt.Println("         ms within =", mse)

	f = msc / mse
	fmt.Print("f ratio is =", f)

}
