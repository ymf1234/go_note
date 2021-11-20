package main

import (
	"fmt"
	"learngo/tree"
)

/**
  封装：
	  名字一般使用大驼峰 例如：CamelCase
	  首字母大写：public
	  首字母小写：private
  包：
     每个目录一个包
     main包包含可执行入口文件
     为结构定义的方法必须放在同一个包内
     可以是不同的文件

	 如何扩展系统类型或者别人的类型
		定义别名
		使用组合
        使用内嵌
 */

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	fmt.Println(root)
	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println("---------------------")

	root.Traverse()
	fmt.Println("---------------------")
	node := myTreeNode{&root}
	node.postOrder()

	//nodes := []treeNode{
	//	{value: 5},
	//	{},
	//	{6, nil, &root},
	//}
	//
	//fmt.Println(nodes)


}
