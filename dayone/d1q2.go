package dayone

import (
	"fmt"
)

// Node represents a node in the expression tree
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// Preorder traversal of the tree
func Preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, " ")
	Preorder(node.Left)
	Preorder(node.Right)
}

// Postorder traversal of the tree
func Postorder(node *Node) {
	if node == nil {
		return
	}
	Postorder(node.Left)
	Postorder(node.Right)
	fmt.Print(node.Value, " ")
}

func Qb() { //main function
	// Manually building the expression tree for "a + b - c"
	root := &Node{Value: "+"}
	root.Left = &Node{Value: "a"}
	root.Right = &Node{Value: "-"}
	root.Right.Left = &Node{Value: "b"}
	root.Right.Right = &Node{Value: "c"}

	fmt.Println("Preorder traversal:")
	Preorder(root)
	fmt.Println()

	fmt.Println("Postorder traversal:")
	Postorder(root)
	fmt.Println()
}
