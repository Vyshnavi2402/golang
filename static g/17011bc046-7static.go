package main

import (
	"fmt"
)

func main() {
	var ax [10]float32
	var ay [10]float32

	ax[0] = 38
	ax[1] = 30
	ax[2] = 35
	ax[3] = 40
	ax[4] = 45
	ax[5] = 50
	ay[1] = 15.9
	ay[2] = 14.9
	ay[3] = 14.1
	ay[4] = 13.3
	ay[5] = 12.5
	var h float32
	var p float32
	h = ax[2] - ax[1]
	p = (ax[0] - ax[3]) / h
	fmt.Printf("the value of h is : %g", h)
	fmt.Printf("\nthe value of p is : %g", p)
	ay[0] = ay[3] + p*(ay[3]-ay[2]) +
		(p*(p+1)*((ay[4]-ay[3])-(ay[3]-ay[2])))/2 +
		(p*(p+1)*(p-1)*((ay[4]-ay[3])-2*(ay[3]-ay[2])+(ay[2]-ay[1])))/6 +
		(p*(p+1)*(p-1)*(p+2)*((ay[5]-ay[4])-3*(ay[4]-ay[3])+3*(ay[3]-ay[2])-(ay[2]-ay[1])))/24

	fmt.Printf("\nthe value of y when x=38 is : %g", ay[0])

}
