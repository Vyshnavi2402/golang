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
	var sx [20]float64
	var ms float64
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

	var f float64
	fmt.Println("ANOVA")
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

	sx[0] = (a[0] + b[0] + c[0] + d[0]) / 4
	sx[1] = (a[1] + b[1] + c[1] + d[1]) / 4
	sx[2] = (a[2] + b[2] + c[2] + d[2]) / 4

	ms = (sx[0] + sx[1] + sx[2]) / 3
	fmt.Print("ms=", ms)

	for i = 0; i < n; i++ {
		ssc = ssc + (4 * math.Pow(sx[i]-ms, 2))
	}
	fmt.Print("\nss between =", ssc)
	//ssc = (4 * math.Pow(sx1-ms, 2)) + (4 * math.Pow(sx2-ms, 2)) + (4 * math.Pow(sx3-ms, 2))
	//fmt.Print("ssc=", ssc)

	for i = 0; i < n; i++ {
		sse = sse + (math.Pow(a[i]-sx[i], 2) + math.Pow(b[i]-sx[i], 2) + math.Pow(c[i]-sx[i], 2) + math.Pow(d[i]-sx[i], 2))
	}
	fmt.Println("       ss within =", sse)

	sstot = ssc + sse
	fmt.Print("ss total =", sstot)

	sscdf = float64(n - 1)
	fmt.Print("\nss between degree of freedom =", sscdf)
	fmt.Print("\nss within degree of freedom =", ssedf)

	msc = (ssc / sscdf)
	fmt.Print("\n\nms between =", msc)

	ssedf = float64((m * n) - n)

	mse = sse / ssedf
	fmt.Println("         ms within=", mse)

	f = msc / mse
	fmt.Print("f ratio is =", f)
}
