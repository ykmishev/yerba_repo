package main 

import (
	"strconv"

	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	s := a[0]

	for i := 1; i < len(a); i++ {
		s = f(s, a[i])
	}
	convert := strconv.Itoa(s)
	for _, chr := range convert {
		z01.PrintRune(chr)
	}
	z01.PrintRune('\n')

}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}