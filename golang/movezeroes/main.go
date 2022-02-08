package main

import "fmt"

func main() {
	array := []int{0, 1, 0, 3, 12}
	moveZeroesSolution2(array)
	fmt.Println(array)
}

func moveZeroesSolution1(nums []int) {
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

func moveZeroesSolution2(nums []int) {

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			j++
		}
	}

}
