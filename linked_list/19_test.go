package linked_list

import (
	"fmt"
	"leetcode/common"
	"testing"
)

/**
删除链表的倒数第N个节点

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的n保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
快慢指针
*/
func removeNthFromEnd(head *common.LinkedListNode, n int) *common.LinkedListNode {
	if head == nil {
		return nil
	}

	var fast = head
	var slow = head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return head
}

func Test_removeNthFromEnd(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5}
	n := 2

	expect := []int{1, 2, 3, 5}

	linkedList := &common.LinkedList{}

	for _, v := range arr {
		linkedList.Add(v)
	}

	out := make([]int, 0)

	head := removeNthFromEnd(linkedList.Head, n)

	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}

	fmt.Println(out)

	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", out) {
		t.Fail()
	}

}
