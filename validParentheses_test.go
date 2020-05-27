package leetcode

import (
	"testing"
)

// Solution to https://leetcode.com/problems/valid-parentheses/
func TestValidParentheses(t *testing.T) {
	in := "()"
	result := isValid(in)
	expectedResult := true

	if result != expectedResult {
		t.Errorf("Got %v; expected %v", result, expectedResult)
	}
}

func TestValidParentheses2(t *testing.T) {
	in := "(())((()))()()"
	result := isValid(in)
	expectedResult := true

	if result != expectedResult {
		t.Errorf("Got %v; expected %v", result, expectedResult)
	}
}

func TestValidParentheses3(t *testing.T) {
	in := "()[]{}"
	result := isValid(in)
	expectedResult := true

	if result != expectedResult {
		t.Errorf("Got %v; expected %v", result, expectedResult)
	}
}

func TestValidParentheses4(t *testing.T) {
	in := "(]"
	result := isValid(in)
	expectedResult := false

	if result != expectedResult {
		t.Errorf("Got %v; expected %v", result, expectedResult)
	}
}

func TestValidParentheses5(t *testing.T) {
	in := "([)]"
	result := isValid(in)
	expectedResult := false

	if result != expectedResult {
		t.Errorf("Got %v; expected %v", result, expectedResult)
	}
}

func TestValidParentheses6(t *testing.T) {
	in := "{[]}"
	result := isValid(in)
	expectedResult := true

	if result != expectedResult {
		t.Errorf("Got %v; expected %v", result, expectedResult)
	}
}

func TestValidParentheses7(t *testing.T) {
	in := "["
	result := isValid(in)
	expectedResult := false

	if result != expectedResult {
		t.Errorf("Got %v; expected %v", result, expectedResult)
	}
}

func TestValidParentheses8(t *testing.T) {
	in := "]"
	result := isValid(in)
	expectedResult := false

	if result != expectedResult {
		t.Errorf("Got %v; expected %v", result, expectedResult)
	}
}

func isValid(s string) bool {
	var stack []string

	for _, r := range s {
		c := string(r)
		if c == "(" || c == "[" || c == "{" {
			stack = append(stack, c)
		} else {
			lastIndex := len(stack) - 1
			if lastIndex < 0 {
				return false
			}
			c2 := stack[lastIndex]
			if c2 == "(" && c == ")" || c2 == "[" && c == "]" || c2 == "{" && c == "}" {
				stack = stack[:lastIndex]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
