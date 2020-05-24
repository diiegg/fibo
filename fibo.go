package main
 
import "fmt"
 
func main(){
    var n int
    f1:=0
    f2:=1
    nextFib:=0
     
    fmt.Print("Enter the number : ")
    fmt.Scan(&n)
    fmt.Print("Fibonacci Series :")
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
