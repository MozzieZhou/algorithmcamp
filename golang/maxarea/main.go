package main

import "fmt"

func main() {
	array := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxAreaSolution2(array))
}

func maxAreaSolution2(height []int) int {

	max := 0
	j := len(height) - 1
	i := 0

	for i < j {
		minHeight := 0
		if height[i] > height[j] {
			minHeight = height[j]
			j--
		} else {
			minHeight = height[i]
			i++
		}

		weight := j - i + 1
		max = maxInt(max, minHeight*weight)
	}

	return max
}

func maxAreaSolution1(height []int) int {

	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			minHeight := minInt(height[i], height[j])
			weight := j - i
			max = maxInt(max, minHeight*weight)
		}
	}

	return max
}

func minInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
