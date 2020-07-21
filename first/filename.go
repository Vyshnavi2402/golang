package main
 
import "fmt"
 
func main() {
       
    
      var x [20]float
      var y [20]float
      var  j,k,n int
      fmt.Printf("enter the value of the elements :");
      fmt.Scanf(&n);
      fmt.Printf("enter the value of x:\n");
     
           for  i:= 1;i <=n; i++{
                      fmt.Scanf(&x[i]);
                       }

   
          for i:=1;i<=n; i++{
                       fmt.Scanf(&y[i]);
                       }
                       h:=x[2]-x[1];
           fmt.Printf("enter the searching point f:");
scanf(&f);
s:=(f-x[n])/h;
d:=y[n];
p:=1;

for i,k:= n,1 ; i,k>=1,n ; i--,k--{
                 for j:=n;j>=1;j--{
                                 y[j]:=y[j]-y[j-1];
                                 }
                                 p:=p*(s+k-1)/k;
                                 d:=d+p*y[n];
}
fmt.Printf("for f ,ans is=%f",f,d);
getch();
}