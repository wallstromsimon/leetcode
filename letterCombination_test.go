package leetcode

import (
	"strconv"
	"testing"
)

// Solution to https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func TestLetterComboSum(t *testing.T) {
	//in := "23456789"
	in := "23"
	result := letterCombinations(in)
	expectedResult := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

	if !stringArraysEqual(result, expectedResult) {
		t.Errorf("Got %v; expected %v", result, expectedResult)
	}
}

func TestLetterComboSum2(t *testing.T) {
	in := "2"
	result := letterCombinations(in)
	expectedResult := []string{"a", "b", "c"}

	if !stringArraysEqual(result, expectedResult) {
		t.Errorf("Got %v; expected %v", result, expectedResult)
	}
}

func TestLetterComboSum3(t *testing.T) {
	in := "22"
	result := letterCombinations(in)
	expectedResult := []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"}

	if !stringArraysEqual(result, expectedResult) {
		t.Errorf("Got %v; expected %v", result, expectedResult)
	}
}

var letters = []string{" ", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	var a []string
	for _, r := range digits {
		dig, _ := strconv.Atoi(string(r)) // or - '0', might be faster, or a rune to string map
		a = append(a, letters[dig])
		//fmt.Println(letters[dig] + "\n- - -")
	}
	return []string{"aa"}
}

func letterCombinations2(res, chars []string) []string {

	return []string{"aa"}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func stringArraysEqual(a, b []string) bool {
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
