package main

import "github.com/01-edu/z01"

func main() {
	chr := "Hello World"
	for _, ch := range chr {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}