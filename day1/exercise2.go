package main

import (
	"fmt"
)

type Node struct{
	Value string
	Left *Node
	Right *Node
}

func preOrderTraversal(node *Node){
	if node == nil{
		return
	}
	fmt.Printf("%s", node.Value)
	preOrderTraversal(node.Left)
	preOrderTraversal(node.Right)
}

func postOrderTraversal(node *Node){
	if node == nil{
		return
	}
	postOrderTraversal(node.Left)
	postOrderTraversal(node.Right)
	fmt.Printf("%s", node.Value)
}

func main(){
	nodeA := &Node{Value:"a"}
	nodeB := &Node{Value:"b"}
	nodeC := &Node{Value:"c"}
	
	nodeMinus := &Node{Value:"-", Left:nodeB, Right:nodeC}
	root := &Node{Value:"+", Left:nodeA, Right:nodeMinus}

	fmt.Println("PreOrderTraversal Root,Left,Right")
	preOrderTraversal(root)
	fmt.Println()

	fmt.Println("PostOrderTraversal Left,Right,Root")
	postOrderTraversal(root)
	fmt.Println()

}