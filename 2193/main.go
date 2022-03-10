package main

import (
	"fmt"
)

func foo(r []rune) int {
	n := len(r)
	if n <= 2 {
		return 0
	}
	leftmost := r[0]
	// fmt.Println("leftmost:", leftmost)
	rightmost := r[n-1]
	// fmt.Println("rightmost:", rightmost)
	sum := 0
	if leftmost == rightmost {
		// recurse
	} else {
		// try to make rightmost character equal to leftmost character
		// i.e. find the rightmost occurence of leftmost character
		i := n-2
		for i > 0 {
			if r[i] == leftmost {
				break
			}
			i--
		}
		d1 := n-1-i

		// try to make leftmost character equal to rightmost character
		// i.e. find the leftmost occurence of rightmost character
		j := 1
		for j < n-1 {
			if r[j] == rightmost {
				break
			}
			j++
		}
		d2 := j

		if d1 <= d2 {
			x := i
			for x < n-1 {
				r[x] = r[x+1]
				x++
			}
			sum = d1
			// recurse
		} else {
			y := j
			for y > 0 {
				r[y] = r[y-1]
				y--
			}
			sum = d2
			// recurse
		}
	}
	// recurse
	return sum + foo(r[1:n-1])
}

func minMovesToMakePalindrome(s string) int {
	runes := []rune(s)
	return foo(runes)
}

func main() {
	fmt.Println("hello, world")
	fmt.Println(minMovesToMakePalindrome("aabb"))
	fmt.Println(minMovesToMakePalindrome("letelt"))
}
