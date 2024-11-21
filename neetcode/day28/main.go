package main

import "fmt"


func search(n []int, target int) int {
	l, r:= 0, len(n)

	for l < r  {
		mid := int((l+r)/2)
		
		if target == n[mid] {
			return mid
		} else if target < n[mid] {
			r = mid
			
		} else if target > n[mid] {
			l = mid+1
			
		}
	}
	return -1

}



func main() {
	nums := []int{-1,0,2,4,6,8}
	nums1 := []int{5}

	fmt.Println(search(nums,4))
	fmt.Println(search(nums,9))
	fmt.Println(search(nums1,5))

}