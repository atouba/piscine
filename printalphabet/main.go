package main

import (
	"github.com/01-edu/z01"
)

//func ret() (int) {
//  return 0
//}

func main() {
	var c rune = 'a'
	var n rune = '\n'
	for c <= 'z' {
		z01.PrintRune(c)
		c++
	}
	z01.PrintRune(n)
}
