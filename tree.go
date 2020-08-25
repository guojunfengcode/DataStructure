package main

import (
	"sync"
	"log"
	"strconv"
	"bytes"
	"fmt"
)

var tree *Tree
var once sync.Once
var mu sync.Mutex
type Node struct {
	data int
	left *Node
	right *Node
}

type Tree struct {
	root *Node
	size int
}

func GetTree() *Tree {
	once.Do(func() {
		tree = &Tree{root: nil, size: 0}
	})
	return tree
	/*
	if tree == nil {
		mu.Lock()
		defer mu.Unlock()
		tree = &Tree{root: nil, size: 0}
	}
	*/
}

func (tree *Tree) GetTreeSize() int {
	return tree.size
}

func (tree *Tree) Add(data int) {
	tree.root = tree.add(tree.root, data)
}

func (tree *Tree) add(n *Node, data int) *Node {
	if n == nil {
		tree.size++
		return &Node{data: data, left: nil, right: nil}
	}
	if data < n.data {
		n.left = tree.add(n.left, data)
	} else if data > n.data {
		n.right = tree.add(n.right, data)
	}
	return n

}

func (tree *Tree) IsExist(data int) bool{
	return tree.isexist(tree.root, data)
}

func (tree *Tree) isexist(n *Node, data int) bool{
	if n == nil {
		return false
	}
	if data == n.data {
		return true
	} else if data < n.data {
		return tree.isexist(n.left, data)
	} else {
		return tree.isexist(n.right, data)
	}
}

func (tree *Tree) FindMax() int{
	if tree.size == 0 {
		log.Printf("tree node is:%v", tree.size)
	}
	return tree.findmax(tree.root).data
}

func (tree *Tree) findmax(n *Node) *Node{
	if n.right == nil {
		return n
	} else {
		return tree.findmax(n.right)
	}
}

func (tree *Tree) FindMin() int{
	if tree.size == 0 {
		log.Printf("tree node is:%v", tree.size)
	}
	return tree.findmin(tree.root).data
}

func (tree *Tree) findmin(n *Node) *Node{
	if n.left == nil {
		return n
	} else {
		return tree.findmin(n.left)
	}
}

func (tree *Tree) Prologue() {
	tree.prologue(tree.root)
}

func (tree *Tree) prologue(n *Node) {
	if n == nil {
		return 
	}
	log.Printf("%v", n.data)
	tree.prologue(n.left)
	tree.prologue(n.right)
}

func (tree *Tree) Middle() {
	tree.middle(tree.root)
}

func (tree *Tree) middle(n *Node) {
	if n == nil {
		return 
	}
	tree.middle(n.left)
	log.Printf("%v", n.data)
	tree.middle(n.right)
}

func (tree *Tree) Post() {
	tree.post(tree.root)
}

func (tree *Tree) post(n *Node) {
	if n == nil {
		return 
	}
	tree.post(n.left)
	tree.post(n.right)
	log.Printf("%v", n.data)
}

func (tree *Tree) String() string {
	var buffer bytes.Buffer
	tree.BstString(tree.root, 0, &buffer)
	return buffer.String()
}

func (tree *Tree) BstString(n *Node, depth int, buffer *bytes.Buffer) {
	if n == nil {
		//buffer.WriteString(tree.Depthstring(depth)+"nil\n")
		return 
	}
	tree.BstString(n.left, depth+1, buffer)
	buffer.WriteString(tree.Depthstring(depth) + strconv.Itoa(n.data) + "\n")
	tree.BstString(n.right, depth+1, buffer)
}

func (tree *Tree) Depthstring(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("-")
	}
	return buffer.String()
}

func (tree *Tree) RemoveMin() int {
	min := tree.FindMin()
	tree.root = tree.removemin(tree.root)
	return min
}

func (tree *Tree) removemin(n *Node) *Node {
	if n.left == nil {
		right := n.right
		tree.size--
		return right
	}
	n.left = tree.removemin(n.left)
	return n
}

func (tree *Tree) RemoveMax() int {
	min := tree.FindMax()
	tree.root = tree.removemax(tree.root)
	return min
}

func (tree *Tree) removemax(n *Node) *Node {
	if n.right == nil {
		left := n.left
		tree.size--
		return left
	}
	n.right = tree.removemax(n.right)
	return n
}

func (tree *Tree) Remove(data int) int {
	tree.remove(tree.root, data)
	return data
}

func (tree *Tree) remove(n *Node, data int) *Node {
	if n == nil {
		return nil
	}
	if tree.root.data == data {
		min := tree.findmin(n.right)
		min.left = n.left
		n.left = nil
		n = n.right
		tree.root = n
		tree.size--
		return n
	}
	if data < n.data {
		n.left = tree.remove(n.left, data)
		return n
	} else if data > n.data {
		n.right = tree.remove(n.right, data)
		return n
	} else {
		if n.left == nil {
			right := n.right
			n.right = nil
			tree.size--
			return right
		}
		if n.right == nil {
			left := n.left
			n.left = nil
			tree.size--
			return left		
		}
		max := tree.findmax(n.right)
		max.right = tree.removemin(n.right)
		max.left = n.left

		n.left = nil
		n.right = nil
		return max
	}
}
func main() {
	root := GetTree();
	//for i := 1; i < 8; i++ {
	//	root.Add(i);
	//}
	root.Add(4)
	root.Add(6)
	root.Add(2)
	root.Add(1)
	root.Add(3)
	root.Add(5)
	root.Add(7)
	log.Printf("----------Prologue----------")
	root.Prologue()
	log.Printf("----------Middle-----------")
	root.Middle()
	log.Printf("----------tail----------")
	root.Post()
	log.Printf("------------------------")
	log.Printf("size:%v", root.GetTreeSize())
	log.Printf("max:%v", root.FindMax())
	log.Printf("min:%v", root.FindMin())
	log.Printf("%v is exist? is %v", 8,root.IsExist(8))
	fmt.Println(tree.String())

	//log.Printf("delete %v", root.RemoveMin())
	//root.Prologue()
	//log.Printf("delete %v", root.RemoveMax())
	//root.Prologue()
	//log.Printf("------------------------")
	//fmt.Println(tree.String())

	tree.Remove(4)
	log.Printf("-remove------------")
	fmt.Println(tree.String())
	
}
