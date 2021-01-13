package binary_tree

import (
	"fmt"
	"leetcode/common"
	"testing"
)

/**
二叉树的最大深度

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
递归
*/
func maxDepth(root *common.TreeNode) int {

	if root == nil {
		return 0
	}

	return common.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

/**
非递归的DFS
*/
func maxDepth1(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	stack := common.NewStack()
	stack.Push(root)

	deep := 0
	for !stack.IsEmpty() {

		deep++

		size := stack.Size()

		for i := 0; i < size; i++ {
			node := stack.Peek().(*common.TreeNode)

			stack.Pop()

			if node.Right != nil {
				stack.Push(node.Right)
			}

			if node.Left != nil {
				stack.Push(node.Left)
			}
		}

	}

	return deep
}

var rootNode = &common.TreeNode{
	Val:  3,
	Left: &common.TreeNode{Val: 9},
	Right: &common.TreeNode{
		Val:   20,
		Left:  &common.TreeNode{Val: 15},
		Right: &common.TreeNode{Val: 7},
	},
}

func Test_maxDepth(t *testing.T) {
	//[3,9,20,null,null,15,7]，

	expect := 3

	if expect != maxDepth(rootNode) {
		t.Fail()
	}
}

func Benchmark_Test_maxDepth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxDepth(rootNode)
	}
}

func Test_maxDepth1(t *testing.T) {
	//[3,9,20,null,null,15,7]，

	expect := 3
	out := maxDepth1(rootNode)

	fmt.Println(out)

	if expect != out {
		t.Fail()
	}
}
