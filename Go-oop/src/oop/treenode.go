package main

import "fmt"

type treeNode struct {
	left, right *treeNode
	Value       int
}

func (node *treeNode) print() {
	fmt.Print(node.Value)
}

func CreateNode(value int) *treeNode {
	// go语言支持返回局部变量的地址给外部使用
	return &treeNode{Value: value}
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode
	fmt.Println(root)

	root.left = &treeNode{}
	root.right = &treeNode{nil, nil, 1}
	root.left = new(treeNode)

	nodes := []treeNode{
		{Value: 3},
		{nil, nil, 6},
		{nil, &root, 6},
	}

	fmt.Println(nodes[1].left)
}
