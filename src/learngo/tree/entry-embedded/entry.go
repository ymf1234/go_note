package main

import (
	"fmt"
	"learngo/tree"
)


// 使用内嵌
type myTreeNode struct {
	*tree.Node //内嵌
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse()  {
	fmt.Println("重载")
}

func main() {

	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Print("In-order traversal: ")
	root.Node.Traverse()

	fmt.Println()

	fmt.Print("My own post-order traversal: ")
	root.postOrder()

}
