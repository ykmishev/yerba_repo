package main 


import (
	"github.com/01-edu/z01"
)

func main() {
	cd := '.'
	for i := '0'; i <= '6'; i++ {
		for j := '0'; j <= '5'; j++ {
			z01.PrintRune(cd)
		} 
	}
	z01.PrintRune('\n')
}
