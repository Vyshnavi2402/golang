package main

import (
	"fmt"
	"math"
)

func main() {

	var sa float64
	var sb float64
	var sc float64
	var t float64
	var n float64
	var cf float64
	var sqtot float64
	var sstot float64
	var ssc float64
	var ssr float64
	var ssiv float64
	var sscdf float64
	var ssrdf float64
	var msc float64
	var msr float64

	a := [2][3]float64{{14, 10, 11},
		{15, 9, 11}}
	b := [2][3]float64{{12, 7, 10},
		{11, 8, 11}}
	c := [2][3]float64{{10, 11, 8},
		{11, 11, 7}}
	// sum of values of each group
	for x := 0; x < 2; x++ {
		for y := 0; y < 3; y++ {
			sa = sa + (a[x][y])
		}
	}
	for x := 0; x < 2; x++ {
		for y := 0; y < 3; y++ {
			sb = sb + (b[x][y])
		}
	}
	for x := 0; x < 2; x++ {
		for y := 0; y < 3; y++ {
			sc = sc + (c[x][y])
		}
	}
	fmt.Print("\nsum of group a=", sa)
	fmt.Print("\nsum of group b=", sb)
	fmt.Print("\nsum of group c=", sc)

	//calculation of corrrection factor
	t = sa + sb + sc
	n = float64(18)
	cf = math.Pow(t, 2) / n
	fmt.Print("\nt=", t)
	fmt.Print("\nn=", n)
	fmt.Print("\ncorrection factor=", cf)

	//calculation of total SS
	for x := 0; x < 2; x++ {
		for y := 0; y < 3; y++ {
			sqtot = sqtot + (math.Pow((a[x][y]), 2) + math.Pow((b[x][y]), 2) + math.Pow((c[x][y]), 2))
		}
	}
	sstot = sqtot - cf

	fmt.Print("\ntotal SS=", sstot)

	//calculation os SS between column
	ssc = ((math.Pow(a[0][0]+a[1][0]+b[0][0]+b[1][0]+c[0][0]+c[1][0], 2) +
		math.Pow(a[0][1]+a[1][1]+b[0][1]+b[1][1]+c[0][1]+c[1][1], 2) +
		math.Pow(a[0][2]+a[1][2]+b[0][2]+b[1][2]+c[0][2]+c[1][2], 2)) / 6) - cf
	fmt.Print("\n SS between column=", ssc)

	//calculation of SS between rows
	ssr = ((math.Pow(sa, 2) + math.Pow(sb, 2) + math.Pow(sc, 2)) / 6) - cf
	fmt.Print("\n SS between rows=", ssr)

	//SS within sample
	sss := 3.5
	fmt.Print("\n SS within sample =", sss)
	//SS for interaction variation
	ssiv = sstot - (ssc + ssr + sss)
	fmt.Print("\n SS interaction variation =", ssiv)

	//degree of freedom
	sscdf = float64(3 - 1)
	ssrdf = float64(3 - 1)
	//mean sum
	msc = ssc / sscdf
	msr = ssr / ssrdf
	fmt.Print("\n msc =", msc)
	fmt.Print("\n msr =", msr)

}
