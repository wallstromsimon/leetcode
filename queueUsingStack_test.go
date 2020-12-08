package leetcode

import (
	"fmt"
	"testing"
)

// Solution to https://leetcode.com/problems/two-sum/
func TestQueueUsingStack(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	param_2 := obj.Pop()
	fmt.Println(param_2)
	param_3 := obj.Peek()
	fmt.Println(param_3)
	param_4 := obj.Empty()
	fmt.Println(param_4)

	t.Errorf("Got ; expected ")
}

type MyQueue struct {
}

/** Initialize your data structure here. */
func Constructor() MyQueue {

	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {

}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

	return 0
}

/** Get the front element. */
func (this *MyQueue) Peek() int {

	return 0
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {

	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
