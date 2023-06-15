package main

import "fmt"

// func main() {
// 	n1 := 0
// 	n2 := 1
// 	n3 := 0
// 	count := 10

// 	fmt.Print(n1, n2 )

// 	for i := 2; i<count; i++{
// 		n3=n1+n2
// 		fmt.Print(+ n3 )
// 		n1 = n2
// 		n2 = n3
// 	}

// }

func fibonacci(n int){
	a,b := 0,1
    fmt.Print(a," ",b)
	for i := 2; i<n; i++{
		c := a+b
	fmt.Print(" ", c)
		a=b
		b=c	
	}
}
