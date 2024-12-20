package main

import "fmt"

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	for _, item := range matrix {
		if search(item, target) {
			return true
		}

	}
	return false
}

func main() {
	matrix := [][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}
	fmt.Println(searchMatrix(matrix, 22))

}
