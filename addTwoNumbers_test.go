package leetcode

import (
	"fmt"
	"testing"
)

// Solution to https://leetcode.com/problems/add-two-numbers/
func TestAddTwoNumbers(t *testing.T) {
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	result := addTwoNumbers(&l1, &l2)
	expectedResult := ListNode{7, &ListNode{0, &ListNode{8, nil}}}

	if !listNodeEqual(result, &expectedResult) {
		t.Errorf("Got %s; expected %s", listNodeToString(result), listNodeToString(&expectedResult))
	}
}

/*
 * Test end with a carry
 */
func TestAddTwoNumbers2(t *testing.T) {
	l1 := ListNode{5, nil}
	l2 := ListNode{5, nil}

	result := addTwoNumbers(&l1, &l2)
	expectedResult := ListNode{0, &ListNode{1, nil}}

	if !listNodeEqual(result, &expectedResult) {
		t.Errorf("Got %s; expected %s", listNodeToString(result), listNodeToString(&expectedResult))
	}
}

/*
 * Test different length
 */
func TestAddTwoNumbers3(t *testing.T) {
	l1 := ListNode{1, &ListNode{8, nil}}
	l2 := ListNode{0, nil}

	result := addTwoNumbers(&l1, &l2)
	expectedResult := ListNode{1, &ListNode{8, nil}}

	if !listNodeEqual(result, &expectedResult) {
		t.Errorf("Got %s; expected %s", listNodeToString(result), listNodeToString(&expectedResult))
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersWithCarry(l1, l2, 0)
}

func addTwoNumbersWithCarry(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry > 0 {
			return &ListNode{carry, nil}
		}
		return nil
	}

	val1 := 0
	var next1 *ListNode = nil
	if l1 != nil {
		val1 = l1.Val
		next1 = l1.Next
	}

	val2 := 0
	var next2 *ListNode = nil
	if l2 != nil {
		val2 = l2.Val
		next2 = l2.Next
	}

	sum := val1 + val2 + carry
	val := sum % 10
	carry = sum / 10
	return &ListNode{val, addTwoNumbersWithCarry(next1, next2, carry)}
}

func listNodeEqual(l1 *ListNode, l2 *ListNode) bool {
	return listNodeToString(l1) == listNodeToString(l2)
}

func listNodeToString(l *ListNode) string {
	if l == nil {
		return "nil"
	}
	return fmt.Sprintf("%d -> %s", l.Val, listNodeToString(l.Next))
}
