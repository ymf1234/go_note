package main

import "fmt"

// 结构体
type treeNode struct {
	value int
	left, right *treeNode
}

func (node treeNode) print()  {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}


func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	fmt.Println(root)
	root.print()
	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println("---------------------")

	nodes := []treeNode{
		{value: 5},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)


}
