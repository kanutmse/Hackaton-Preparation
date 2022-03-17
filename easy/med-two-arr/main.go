package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	newArr := append(nums1, nums2...)

	fmt.Printf("%+v", newArr)
	sort.Sort(sort.IntSlice(newArr))
	fmt.Printf("%+v", newArr)

	if len(newArr)%2 == 0 {

		var start int
		var end int
		var ans float64
		start = ((len(newArr)) / 2) - 1

		end = ((len(newArr)) / 2)
		// fmt.Printf("%d %d\n", start, end)
		ans = (float64(newArr[start]) + float64(newArr[end])) / 2
		return ans
	}
	fmt.Printf("%d ", len(newArr)/2)

	return float64(newArr[len(newArr)/2])

}

func main() {
	fmt.Printf("%.2f\n", findMedianSortedArrays([]int{1, 3}, []int{2}))
	// fmt.Printf("%.2f\n", findMedianSortedArrays([]int{1, 2}, []int{3, 4, 5, 6, 7, 8, 9, 10}))
}
