package linked_list

import (
	"fmt"
	"leetcode/common"
	"testing"
)

/**
合并两个有序链表

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
递归
*/
func mergeTwoLists(l1 *common.LinkedListNode, l2 *common.LinkedListNode) *common.LinkedListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val >= l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

}

func Test_mergeTwoLists(t *testing.T) {

	l1 := &common.LinkedList{}
	l1.Add(1)
	l1.Add(2)
	l1.Add(4)

	l2 := &common.LinkedList{}
	l2.Add(1)
	l2.Add(3)
	l2.Add(4)

	expect := []int{1, 1, 2, 3, 4, 4}
	out := make([]int, 0)

	head := mergeTwoLists(l1.Head, l2.Head)

	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}

	fmt.Println(out)

	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", out) {
		t.Fail()
	}

}

func mergeTwoLists2(l1 *common.LinkedListNode, l2 *common.LinkedListNode) *common.LinkedListNode {

	//哨兵节点
	preHead := &common.LinkedListNode{}
	result := preHead

	// 指向l1或者l2中较小的一个，直到l1或者l2任一指向null
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			preHead.Next = l1
			l1 = l1.Next
		} else {
			preHead.Next = l2
			l2 = l2.Next
		}
		preHead = preHead.Next
	}

	//余下的这些元素直接合并
	if l1 != nil {
		preHead.Next = l1
	}
	if l2 != nil {
		preHead.Next = l2
	}

	return result.Next

}

func Test_mergeTwoLists2(t *testing.T) {

	l1 := &common.LinkedList{}
	l1.Add(1)
	l1.Add(2)
	l1.Add(4)

	l2 := &common.LinkedList{}
	l2.Add(1)
	l2.Add(3)
	l2.Add(4)

	expect := []int{1, 1, 2, 3, 4, 4}
	out := make([]int, 0)

	head := mergeTwoLists2(l1.Head, l2.Head)

	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}

	fmt.Println(out)

	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", out) {
		t.Fail()
	}

}
