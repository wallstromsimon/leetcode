package leetcode

import (
    "testing"
)

// Solution to https://leetcode.com/problems/two-sum/
func TestTwoSum(t *testing.T) {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    expectedResult := []int{0, 1}

    if !equal(result, expectedResult) {
        t.Errorf("Got %v; expected %v", result, expectedResult)
    }
}

func TestTwoSum2(t *testing.T) {
    nums := []int{3, 2, 4}
    target := 6
    result := twoSum(nums, target)
    expectedResult := []int{1, 2}

    if !equal(result, expectedResult) {
        t.Errorf("Got %v; expected %v", result, expectedResult)
    }
}

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if i == j {
                continue
            }
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
