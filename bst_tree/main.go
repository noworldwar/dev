package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	key   int
	left  *Node
	right *Node
}

func main() {
	var t Tree
	t.insert(-10)
	t.insert(-3)
	t.insert(0)
	t.insert(5)
	t.insert(9)
	// fmt.Println("printPreOrder: ")
	// printPreOrder(t.root)
	// fmt.Printf("\n")
	// fmt.Println("printPostOrder: ")
	// printPostOrder(t.root)
	// fmt.Printf("\n")
	// fmt.Println("printInOrder: ")
	// printInOrder(t.root)
	// fmt.Printf("\n")
}

func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

func (n *Node) insert(data int) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func printPreOrder(n *Node) {
	if n == nil {
		fmt.Printf("%s ", "null")
		return
	} else {
		fmt.Printf("%d ", n.key)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printPostOrder(n *Node) {
	if n == nil {
		fmt.Printf("%s ", "null")
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%d ", n.key)
	}
}

func printInOrder(n *Node) {
	if n == nil {
		fmt.Printf("%s ", "null")
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%d ", n.key)
		printInOrder(n.right)
	}
}
