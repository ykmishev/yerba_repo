package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	programName := os.Args[0]
	for i, ch := range programName {
		if i >= 2 {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
