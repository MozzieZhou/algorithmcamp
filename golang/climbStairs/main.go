package main

import "fmt"

func main() {
	fmt.Println(addSolution2(4))
}

func addSolution1(n int) int {

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return addSolution1(n-1) + addSolution1(n-2)
}

func addSolution2(n int) int {

	if n <= 2 {
		return n
	}
	a, b, c := 1, 2, 3

	for i := 3; i < n; i++ {
		c = b + c // c = 5 a=1 b=2
		a = b     // a=2 b=2 c=5
		b = c - a
	}
	return c
}
