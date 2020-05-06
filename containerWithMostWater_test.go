package leetcode

import (
    "testing"
)

// Solution to https://leetcode.com/problems/container-with-most-water/
func TestContainerWithMostWater(t *testing.T) {
    result := maxArea([]int{1,8,6,2,5,4,8,3,7})
    expectedResult := 49
    if result != expectedResult {
        t.Errorf("Got %d; expected %d", result, expectedResult)
    }
}

func maxArea(height []int) int {
    maxArea := 0
    for i := 0; i < len(height); i++ {
        for j := len(height) - 1; j > i; j-- {
            area := (j - i) *  min(height[i], height[j])
            if area > maxArea {
                maxArea = area
            }
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
