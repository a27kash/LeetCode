package main

import (
	"fmt"
)

// assumption: x,y E I+
func areCoPrime(x, y int) (b bool) {
	if x == y {
		return false
	}
	if x > y {
		t := x
		x = y
		y = t
	}
	if x == 1 {
		return true
	}
	r := y%x
	switch r {
	case 0:
		b = false
	case 1:
		b = true
	default:
		b = areCoPrime(r, x)
	}
	return 
}

// assumption: x,y E I+
func gcd(x, y int) (g int) {
	if x == y {
		return x
	}
	if x > y {
		t := x
		x = y
		y = t
	}
	r := y%x
	switch r {
	case 0:
		g = x
	case 1:
		g = 1
	default:
		g = gcd(r, x)
	}
	return
}

// assumption: x,y E I+
func lcm(x, y int) int {
	if x == y {
		return x
	}
	return x * y / gcd(x, y)
}

func replaceNonCoprimes(nums []int) []int {
	output := make([]int, len(nums))
	output[0] = nums[0]
	i := 0 // index on output
	n := 1 // len of output
	j := 1 // index on nums
	for j < len(nums) {
		x := output[i]
		y := nums[j]
		// fmt.Println("x=", x, "; y=", y)
		if !areCoPrime(x, y) {
			output[i] = lcm(x,y)
			// compaction
			for i > 0 {
				xx := output[i-1]
				yy := output[i]
				if !areCoPrime(xx, yy) {
					output[i-1] = lcm(xx,yy)
					i--
					n--
				} else {
					break
				}
			}
		} else {
			i++
			n++
			output[i] = y
		}
		j++
		// fmt.Println(output[:n])
		// fmt.Println(nums[j:])
	}
	return output[:n]
}

func main() {
	fmt.Println("hello, world")
	fmt.Println(areCoPrime(2,6))
	fmt.Println(areCoPrime(8,12))
	fmt.Println(areCoPrime(1,2))
	fmt.Println(areCoPrime(5,6))
	fmt.Println(gcd(2,6))
	fmt.Println(gcd(8,12))
	fmt.Println(gcd(1,2))
	fmt.Println(gcd(5,6))
	fmt.Println(lcm(2,6))
	fmt.Println(lcm(8,12))
	fmt.Println(lcm(1,2))
	fmt.Println(lcm(5,6))

	nums1 := []int {
		6,4,3,2,7,6,2,
	}

	output1 := replaceNonCoprimes(nums1)

	fmt.Println(output1)

	nums2 := []int {
		2,2,1,1,3,3,
	}

	output2 := replaceNonCoprimes(nums2)

	fmt.Println(output2)

	nums3 := []int {
		2,5,10,
	}

	output3 := replaceNonCoprimes(nums3)

	fmt.Println(output3)

	nums4 := []int {
		287,41,49,287,899,23,23,20677,5,825,
	}

	output4 := replaceNonCoprimes(nums4)

	fmt.Println(output4)
}
