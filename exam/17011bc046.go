package main
 
import "fmt"
import "math"
 
func main() {
fmt.Println("the given eqation is (x*x)-(5*x)+6")
var a float64 =1 
var b float64=-5 
var c float64=6 
var x1 float64
var x2 float64
var d float64
d=(b*b)-(4*a*c)
 var x float64
 x = math.Sqrt(d)
 
x1=(-b+x)/(2*a)
x2=(-b-x)/(2*a) 
fmt.Println("the zeroes of the above eqation are :")
fmt.Println(x1)
fmt.Println(x2)
}