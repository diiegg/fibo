package main
 
import "fmt"
 
func main(){
// Declare variables n, fi, f2 and nextFib
    var n int
    f1:=0
    f2:=1
    nextFib:=0
// Print Enter the number
    fmt.Print("Enter the number : ")
// search for the number as user input
    fmt.Scan(&n)
// Print Fibonacci series
    fmt.Print("Fibonacci Series :")
// Iterate until desired position in sequence.
    for i:=1;i<=n;i++ {
        if(i==1){
            fmt.Print(" ",f1)
            continue
        }
        if(i==2){
            fmt.Print(" ",f2)
            continue
        }
        nextFib = f1 + f2
        f1=f2
        f2=nextFib
        fmt.Print(" ",nextFib)
    }
}
