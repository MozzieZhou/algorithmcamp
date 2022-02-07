package main

import "fmt"

func main() {
	array := []int{0, 1, 0, 3, 12}
	moveZeroes(array)
	fmt.Println(array)
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}

}
