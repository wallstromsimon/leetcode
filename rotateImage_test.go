package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

// Solution to https://leetcode.com/problems/rotate-image/
func TestRotateImage(t *testing.T) {
	var m = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	rotate(m)

	var expected = [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Nope")
		printM(m)
		fmt.Println(" --- ")
		printM(expected)
	}
}

func TestRotateImage2(t *testing.T) {
	var m = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}

	rotate(m)

	var expected = [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Nope")
		printM(m)
		fmt.Println(" --- ")
		printM(expected)
	}
}

func rotate(matrix [][]int) {
	m := len(matrix)
	start, end := 0, m-1
	for start < end {
		num := end - start + 1
		for i := 0; i < num-1; i++ {
			matrix[start+i][end], matrix[start][start+i], matrix[end-i][start], matrix[end][end-i] = matrix[start][start+i], matrix[end-i][start], matrix[end][end-i], matrix[start+i][end]
		}
		start, end = start+1, end-1
	}
}

func printM(m [][]int) {
	/* output each array element's value */
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			fmt.Printf("m[%d][%d] = %d\n", i, j, m[i][j])
		}
	}
}
