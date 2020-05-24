# fibo

[![Build Status](https://travis-ci.org/diiegg/fibo.svg?branch=master)](https://travis-ci.org/diiegg/fibo)
[![Go Report Card](https://goreportcard.com/badge/github.com/diiegg/fibo)](https://goreportcard.com/report/github.com/diiegg/fibo)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/diiegg/fibo/blob/master/MIT.txt)
[![Coverage Status](https://coveralls.io/repos/github/diiegg/fibo/badge.svg)](https://coveralls.io/github/diiegg/fibo)
[![Build Status](https://cloud.drone.io/api/badges/diiegg/fibo/status.svg)](https://cloud.drone.io/diiegg/fibo)

# Fibonacci Series Algorithm in GO with a variable declaration

# Introduction!
     
By definition, the first two numbers in the Fibonacci sequence are either 1 and 1, or 0 and 1, depending on the chosen starting point of the sequence, and each subsequent number is the sum of the previous two.

A example of the serie shoulb be like this .....
```sh
1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 …
```

### Example method

In this case Fibonacci series is an integer sequence where every number after the first two is the sum of the two preceding one which receives an int and returns another int, where fibonacci is calculates as as sum of n-1th term and n-2th term.

```sh
F(n)=F(n-1)+F(n-2)
where F{0} = 0 and F{1} = 1
```
### Func
We use a for-loop to iterate until the desired position in the sequence. We continually add up numbers until we reach that position.

We are working with slices because we cannot create arrays with numeric lengths that are not constant. The first thing we do is create the slice which has a length of n+1 and a capacity of n+2. The length needs to be more than n because if we pass in a value of two, we also need to accommodate the zero index which means our slice needs to be of length three, not two.

If our placeholder is less than two, we need to increase the length of our slice to hold at least the first and second Fibonacci numbers. Finally we can loop similarly to how we did with recursion and return the final value of the slice. Each resulting value in the loop is stored in the slice.

```sh
 for i:=1;i<=n;i++ {
        if(i==1){
            fmt.Print(" ",f1)
            continue
        }
        if(i==2){
            fmt.Print(" ",f2)
            continue
        }
        nextTerm = f1 + f2
        f1=f2
        f2=nextFib
        fmt.Print(" ",nextFib)
```



### Variable
It read the input number and calculatete the fibonachi number then print the Fibonacci serie.

```sh
var n int
fmt.Scan(&n)

```

### Installation

to Run  a test :

```sh
go run fib.go
```

Output :

```sh
❯ go run fibo.go
Enter the number : 3
Fibonacci Series : 0 1 1% 
```

License
----

MIT


**Free Software, Hell Yeah!**
