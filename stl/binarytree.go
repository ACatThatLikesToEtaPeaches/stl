package stl

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1.中序遍历，递归方式
func InorderTraverse(root *TreeNode) (res []int) {
	inorderTraverseHelper(root, &res)
	return res
}
func inorderTraverseHelper(root *TreeNode, res *[]int) {
	if root == nil { return }
	inorderTraverseHelper(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraverseHelper(root.Right, res)
}


// 5. 通过前序遍历 中序遍历结果 还原二叉树
func BuildTreeByPreorderAndInorder(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	n := len(preorder)
	mp := make(map[int]int, n)
	// 构造哈希映射，快速定位前序遍历中元素在中序遍历中的下标
	for i, x := range inorder {
		mp[x] = i
	}
	return buildHelper(preorder, inorder, mp,0, n-1, 0, n-1)
}
// preorder: x| x x x x ~x x
// inorder:  x x x x| x| x x
func buildHelper(pre []int, in []int, mp map[int]int, preleft int, preright int, inleft int, inright int) *TreeNode{
	if preleft > preright { return nil }

	preroot := preleft // 前序遍历的第一个节点就是根结点下标
	curRootVal := pre[preroot]
	inroot := mp[curRootVal] // 前序遍历元素在中序遍历中结点下标，也就是当前根结点在中序遍历结果的下标
	sizeLeftTree := inroot - inleft // 注意一定是用inroot-inleft 得到leftTree的数量

	root := &TreeNode{Val: curRootVal}
	root.Left = buildHelper(pre, in, mp, preleft+1, preleft+sizeLeftTree, inleft, inroot-1)
	root.Right = buildHelper(pre, in, mp, preleft+sizeLeftTree+1, preright, inroot+1, inright)

	return root
}
