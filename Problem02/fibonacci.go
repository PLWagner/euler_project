package main

import "fmt"

const max  = 4000000

func is_even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	var sum = 0
	var seq = 1
	var last = 0
	var actual = 1
	
	for seq <= max {
		fmt.Println(seq)
		if is_even(last + seq) {
			sum = sum + last + seq
		}
		seq = last + actual
		last = actual
		actual = seq
	}
	fmt.Println("The sum of even numbers of Fibonnaci sequence is: ", sum)
}
