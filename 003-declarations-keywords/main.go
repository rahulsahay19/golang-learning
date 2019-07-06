package main

import "fmt"

var a = 20 //var is used, outside the func scope. Its meant for package level scope.

func main() {
	x := 19        // declaration and assignment both
	fmt.Println(x) //will print 19
	x = 77
	fmt.Println(x) //will print 77
	y := 100 + 27  //statement
	fmt.Println(y) // will print 127
	z := "Rahul, Sahay"
	fmt.Println(z) // will print Rahul, Sahay
	fmt.Println(a)
	foo()
}

func foo() {
	fmt.Println("Printing global value", a)
}
