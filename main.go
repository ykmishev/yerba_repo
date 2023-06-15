package main

import "fmt"

func main() {

	var n int

	fmt.Scanln(&n)

	slice := make([]string, n)
	for i := 0; i < n; i++ {
		slice[i] = string('a' + i)
	}

	fmt.Println(slice)
}
