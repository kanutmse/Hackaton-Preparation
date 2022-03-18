package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//Brute Force But O(n^2) Limit time exceed feel bad man
func maxArea(height []int) int {

	max := 0
	for i := 0; i < len(height)-1; i++ {

		for j := i + 1; j < len(height); j++ {
			rst := min(height[i], height[j]) * (j - i)

			if max < rst {
				max = rst
			}

		}
	}

	return max
}

//Optimized
func maxAreaOp(height []int) int {

	n := len(height)

	var l, result int
	r := n - 1

	for l < r {
		// Select the minimum of left and right pointer heights for area calculation
		currArea := min(height[l], height[r]) * (r - l)
		// Select the maximum of result so far and current area result
		result = max(result, currArea)

		// If the current l height is less than r height,
		// there is no use selecting l for next iteration because, width decreases, so area cant increase
		// So l++(select r, dont select l for next iteration)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return result
}
