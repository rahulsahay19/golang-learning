package main

import "fmt"

var a = 27
var b = "Hello Go!"
var c = `"Raw string literals"`

type rahul int

var d rahul

func main() {
	d = 97
	fmt.Println(a)        // will print 27
	fmt.Printf("%T\n", a) // will print int
	fmt.Printf("%T\n", b) // will print string
	fmt.Println(c)        // will print string with quotes
	fmt.Println(d)        // will print 97
	a = int(d)            // type conversion
	fmt.Println(d)        // will print 97
	fmt.Printf("%T\n", a) // will print int
	fmt.Printf("%T\n", d) // will print main.rahul
}
