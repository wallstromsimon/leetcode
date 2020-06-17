package leetcode

import (
	"testing"
)

// Solution to https://leetcode.com/problems/container-with-most-water/
func TestContainerWithMostWater(t *testing.T) {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	expectedResult := 49
	if result != expectedResult {
		t.Errorf("Got %d; expected %d", result, expectedResult)
	}
}

func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
