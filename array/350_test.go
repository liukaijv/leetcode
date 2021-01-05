package array

import (
	"fmt"
	"sort"
	"testing"
)

/*
第350题：两个数组的交集

示例1

输入: nums1 = [1,2,2,1], nums2 = [2,2]

输出: [2,2]

示例2

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]

输出: [4,9]

说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

*/

func intersect(arr1, arr2 []int) []int {
	out := make([]int, 0)
	m := make(map[int]int)

	for _, v := range arr1 {
		m[v] += 1
	}

	for _, v := range arr2 {
		if m[v] > 0 {
			m[v] -= 1
			out = append(out, v)
		}
	}
	return out
}

func TestIntersect1(t *testing.T) {

	arr1 := []int{1, 2, 2, 1}
	arr2 := []int{2, 2}
	expect := []int{2, 2}
	out := intersect(arr1, arr2)

	sort.Ints(out)
	fmt.Printf("%v", out)
	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", out) {
		t.Fail()
	}

}

func TestIntersect2(t *testing.T) {

	arr1 := []int{4, 9, 5}
	arr2 := []int{9, 4, 9, 8, 4}
	expect := []int{4, 9}
	out := intersect(arr1, arr2)

	sort.Ints(out)
	fmt.Printf("%v", out)

	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", out) {
		t.Fail()
	}

}

/**
考虑排序
双指针
*/

func intersectSorted(arr1, arr2 []int) []int {
	out := make([]int, 0)
	i, j := 0, 0

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i < len(arr1) && j < len(arr2) {

		if arr1[i] > arr2[j] {
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			out = append(out, arr1[i])
			i++
			j++
		}

	}

	return out
}

func TestIntersectSorted(t *testing.T) {
	arr1 := []int{4, 9, 5}
	arr2 := []int{9, 4, 9, 8, 4}
	expect := []int{4, 9}
	out := intersectSorted(arr1, arr2)

	fmt.Printf("%v", out)

	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", out) {
		t.Fail()
	}
}
