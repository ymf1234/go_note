package tree

import "fmt"

// 结构体
type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print()  {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}


