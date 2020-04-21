package leetcode

import (
    "sort"
    "testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
    result := findMedianSortedArrays([]int{1, 3}, []int{2})
    expectedResult := 2.0
    if result != expectedResult {
        t.Errorf("Got %f; expected %f", result, expectedResult)
    }
}

func TestFindMedianSortedArrays2(t *testing.T) {
    result := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
    expectedResult := 2.5
    if result != expectedResult {
        t.Errorf("Got %f; expected %f", result, expectedResult)
    }
}

func TestFindMedianSortedArrays3(t *testing.T) {
    result := findMedianSortedArrays([]int{1, 2}, []int{})
    expectedResult := 1.5
    if result != expectedResult {
        t.Errorf("Got %f; expected %f", result, expectedResult)
    }
}

// Solution to https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    full := append(nums1, nums2...)
    sort.Ints(full)
    //fmt.Println(full)
    length := len(full)
    mid := length/2
    //fmt.Println(mid)
    //fmt.Println(len(full))
    if length % 2 == 0 {
        //fmt.Printf("%d %d\n", full[mid -1 ], full[mid])
        return float64(full[mid-1] + full[mid]) / 2.0
    }
    return float64(full[mid])
}
