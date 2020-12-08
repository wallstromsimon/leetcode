package leetcode

import (
	"testing"
)

// Solution to https://leetcode.com/problems/climbing-stairs/submissions/
func TestClimbStairs(t *testing.T) {
	result := climbStairsLoop(2)
	expectedResult := 2
	if result != expectedResult {
		t.Errorf("Got %d; expected %d", result, expectedResult)
	}
}

func TestClimbStairs2(t *testing.T) {
	result := climbStairsLoop(3)
	expectedResult := 3
	if result != expectedResult {
		t.Errorf("Got %d; expected %d", result, expectedResult)
	}
}

func climbStairsLoop(n int) int {
	if n <= 1 {
		return n
	}
	pre, cur := 1, 1
	for i := 1; i < n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}

func climbStairs(steps int) int {
	return climbStairsRe(0, steps)
}
// Too slow
func climbStairsRe(currentStep int, steps int) int {
	if currentStep > steps {
		return 0
	}
	if currentStep == steps {
		return 1
	}

	return climbStairsRe(currentStep+1, steps) + climbStairsRe(currentStep+2, steps)
}
