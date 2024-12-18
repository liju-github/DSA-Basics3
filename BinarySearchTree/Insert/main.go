package main

import "fmt"

type Node struct{
	Value int
	Left *Node
	Right *Node
}


func (n *Node)Insert(value int)  {
	newnode := &Node{Value: value}

	if n == nil{
		return
	}

	if n.Value < value{
		if n.Left == nil{
			n.Left = newnode
		}else{
			n.Left.Insert(value)
		}	
	}else{
		if n.Right == nil{
			n.Right = newnode
		}else{
			n.Right.Insert(value)
		}
	}
}

func (n *Node)Inorder()  {
	
	if n == nil{
		return
	}

	n.Left.Inorder()
	fmt.Print("->",n.Value)
	n.Right.Inorder()
}

func main()  {
	root := &Node{Value: 100}

	root.Insert(90)
	root.Insert(80)
	root.Insert(70)
	root.Insert(60)
	root.Insert(50)
	root.Insert(40)
	root.Insert(30)

	root.Inorder()
	fmt.Println()
}