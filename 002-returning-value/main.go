package main

import "fmt"

func main() {
	n, err := fmt.Println("Returning Values")
	println(n)   // This will return number of bytes
	println(err) // This will return error incase any
}
