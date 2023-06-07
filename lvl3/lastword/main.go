package main 

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 {
		return
	}
	var flag bool

	args := os.Args[1]
	for i := len(args) - 1; i >= 0; i-- {
		n := rune(args[i])
		if n == ' ' {
			if !flag {
				continue
			}
			return
		} else {
			if !flag {
				defer z01.PrintRune('\n')
			}
			flag = true
			z01.PrintRune(n)
		}
	}
}