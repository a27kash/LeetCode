package main

import (
	"fmt"
)

func mergesort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	mid := n/2
	a := mergesort(nums[:mid])
	b := mergesort(nums[mid:])
	output := make([]int, n)
	i := 0 // index on a
	j := 0 // index on b
	k := 0 // index on output
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			output[k] = a[i]
			i++
		} else {
			output[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		output[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		output[k] = b[j]
		j++
		k++
	}
	return output
}

func minimalKSum(nums []int, k int) int64 {
	n := len(nums)
	sorted := mergesort(nums)
	i := 0 // index on sorted
	v := 1 // value
	var sum int64
	sum = 0
	for i < n && k > 0 {
		if sorted[i] == v {
			i++
			for i < n && sorted[i] == v {
				i++
			}
			v++
		} else {
			sum += int64(v)
			v++
			k--
		}
	}
	max := sorted[n-1]
	d := 1
	for k > 0 {
		sum += int64(max + d)
		k--
		d++
	}
	return sum
}

func main() {
	fmt.Println("hello, world")
	fmt.Println(mergesort([]int{1,}))
	fmt.Println(mergesort([]int{1,4,25,10,25,}))

	fmt.Println(minimalKSum([]int{1,4,25,10,25,}, 2))
	fmt.Println(minimalKSum([]int{5,6,}, 6))
}
