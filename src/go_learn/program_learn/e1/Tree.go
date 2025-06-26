package main

import "fmt"

// 树
type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

// 先序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Print(tree.Data, " ")
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	MidOrder(tree.Left)
	fmt.Print(tree.Data, " ")
	MidOrder(tree.Right)

}

func AfterOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	AfterOrder(tree.Left)
	AfterOrder(tree.Right)
	fmt.Print(tree.Data, " ")
}

func main() {

	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}

	fmt.Println("先序排序：")
	PreOrder(t)
	fmt.Println("\n中序排序：")
	MidOrder(t)
	fmt.Println("\n后序排序")
	AfterOrder(t)

}
