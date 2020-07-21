package main

import (
	"fmt"
)

func main() {
	var n int
	var otot, ef, chis, n1 float64
	var i int
	var o [20]float64
	var e [20]float64
	fmt.Println("Enter no of terms:::")
	fmt.Scan(&n)
	fmt.Println("enter the observed frequency valuees")
	for i = 0; i < n; i++ {
		fmt.Printf("Enter the value of 'o' %d ::", i+1)
		fmt.Scan(&o[i])
	}

	fmt.Println("Observed frequency values are")
	otot = 0

	for i = 0; i < n; i++ {
		otot = otot + o[i]
		fmt.Println(o[i])
	}
	fmt.Print("\n")
	fmt.Println(" (o-e)^2 / e")
	n1 = float64(n)
	ef = otot / n1

	for i = 0; i < n; i++ {
		e[i] = ((o[i] - ef) * (o[i] - ef)) / ef
		fmt.Println(e[i])
	}
	for i = 0; i < n; i++ {
		chis = chis + e[i]
	}

	fmt.Println("otot=", otot)
	fmt.Println("Expected frequency", ef)
	fmt.Println("chis=", chis)
	fmt.Println("chi-square value=", chis)
}
