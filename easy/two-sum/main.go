package main

import "fmt"

//Use less memory(array and map ) but slower  O(n^2)
func twoSum(nums []int, target int) []int {

	ans := []int{0, 0}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				ans[0] = i
				ans[1] = j
				return ans
			}

		}

	}
	return ans
}

//Use more memory(array and map ) but faster  O(n)
func twoSumOptimized(nums []int, target int) []int {
	tmp := make(map[int]int)

	ans := []int{0, 0}

	for k, v := range nums {
		tmp[v] = k
	}

	for k, v := range nums {

		contain := target - v
		if w, ok := tmp[contain]; ok && tmp[contain] != k {
			ans[0] = k
			ans[1] = w
			return ans
		}

	}

	return ans
}

func main() {

	fmt.Printf("%+v", twoSumOptimized([]int{3, 2, 4}, 6))

}
