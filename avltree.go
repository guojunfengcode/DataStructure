package main

import (
	"fmt"
	"log"
	"errors"
	"strconv"
	"bytes"
)

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

type AVLnode struct {
	data interface{}
	left *AVLnode
	right *AVLnode
	height int
}

func comparator(a, b interface{}) int {
	var A,B int
	var ok bool
	if A, ok = a.(int); !ok{
		return -2
	}
	if B, ok = b.(int); !ok {
		return -2
	}
	if A > B {
		return 1
	} else if A < B {
		return -1
	} else {
		return 0
	}
	
}

func NewNode(data interface{}) *AVLnode{
	node := new(AVLnode)
	node.data = data
	node.left = nil
	node.right = nil
	node.height = 1
	return node
}

func NewAVLTree(data interface{}) (*AVLnode, error) {
	if data == nil {
		return nil,errors.New("AVL Tree not null")
	}

	return NewNode(data), nil
}

func (node *AVLnode) LeftRotate() *AVLnode {
	root := node.right
	node.right = root.left
	root.left = node

	node.height = Max(node.left.GetHeight(), node.right.GetHeight())+1
	root.height = Max(root.left.GetHeight(), root.right.GetHeight())+1

	return root
}

func (node *AVLnode) RightRotate() *AVLnode {
	if node == nil {
		return nil
	}
	root := node.left
	node.left = root.right
	root.right = node
	
	node.height = Max(node.left.GetHeight(), node.right.GetHeight())+1
	root.height = Max(root.left.GetHeight(), root.right.GetHeight())+1

	return root
}

func (node *AVLnode) LeftThenRightRotate() *AVLnode {
	sonnode := node.left.LeftRotate()
	node.left = sonnode
	return node.RightRotate()
}

func (node *AVLnode) RightThenLeftRotate() *AVLnode {
	sonnode := node.right.RightRotate()
	node.right = sonnode
	return node.LeftRotate()
}

func (node *AVLnode) AutoBalance() *AVLnode {
	if node.right.GetHeight() - node.left.GetHeight() == 2 {
		if node.right.right.GetHeight() > node.right.left.GetHeight() {
			node = node.LeftRotate()
		} else {
			node = node.RightThenLeftRotate()
		}
	} else if node.left.GetHeight() - node.right.GetHeight() == 2 {
		if node.left.left.GetHeight() > node.left.right.GetHeight() {
			node = node.RightRotate()
		} else {
			node = node.LeftThenRightRotate()
		}
	}

	return node
}

func (node *AVLnode) Add(value interface{}) *AVLnode{
	if node == nil {
		newnode := &AVLnode{data: value, left: nil, right: nil, height:1}
		return newnode
	}
	switch comparator(value, node.data) {
	case -1:
		node.left = node.left.Add(value)
		node = node.AutoBalance()
	case 1:
		node.right = node.right.Add(value)
		node = node.AutoBalance()
	case 0:
		log.Printf("data:%v is exist", value)
	}
	node.height = Max(node.left.GetHeight(), node.right.GetHeight()) + 1
	return node
}

func (node *AVLnode) Delete(value interface{}) *AVLnode {
	if node == nil {
		return nil
	}
	if replace,_ := node.FindData(value); replace == nil {
		return nil
	} else {
		switch comparator(value, node.data) {
		case -1:
			node.left = node.left.Delete(value)
		case 1:
			node.right = node.right.Delete(value)
		case 0:
			if node.left != nil && node.right != nil { //左右都有节点
				node.data = node.right.FindMin().data
				node.right = node.right.Delete(node.data)
			} else if node.left != nil { //有左节点，右节点可有可无
				node = node.left
			} else { //有右节点或无子节点
				node = node.right
			}
		}
	}
	if node != nil {
		node.height = Max(node.left.GetHeight(), node.right.GetHeight()) + 1
		node = node.AutoBalance()
	}
	return node
}

func (node *AVLnode) Printf() []interface{}{
	if node == nil {
		return nil
	}
	var result []interface{}
	queue := []*AVLnode{node}
	for len(queue) != 0 {
		temp := []*AVLnode{}
		for _, v := range queue {
			result = append(result, v.data)
			if v.left != nil {
				temp = append(temp, v.left)
			}
			if v.right != nil {
				temp = append(temp, v.right)
			}
		}
		queue = temp
	}
	return result	
}

func (node *AVLnode) Find(data interface{}) *AVLnode{
	var find *AVLnode = nil
	switch comparator(data, node.data) {
	case -1:
		find = node.left.Find(data)
	case 1:
		find = node.right.Find(data)
	case 0:
		return node
	}
	return find
}

func (node *AVLnode) FindData(data interface{}) (*AVLnode, bool) {
	if node == nil {
		return nil, false
	}
	if data == node.data {
		return node, true
	} else if data.(int) < node.data.(int) {
		return node.left.FindData(data)
	} else {
		return node.right.FindData(data)
	}
}

func (node *AVLnode) FindMax() *AVLnode{
	var find *AVLnode
	if node.right != nil {
		find = node.right.FindMax()
	} else {
		find = node
	}
	return find
}

func (node *AVLnode) FindMin() *AVLnode{
	var find *AVLnode
	if node.left != nil {
		find = node.left.FindMin()
	} else {
		find = node
	}
	return find
}

func (node *AVLnode) GetData() interface{}{
	if node == nil {
		return nil
	}
	return node.data
}

func (node *AVLnode) SetData(data interface{}, setdata interface{}) *AVLnode{
	if node == nil {
		return nil
	}
	if replace,ok := node.FindData(data); replace == nil {
		log.Printf("%v is not exist,%v", data, ok)
	} else {
		node = node.Delete(replace.data)
		node = node.Add(setdata)	
	}
	return node
}

func (node *AVLnode) GetLeft() *AVLnode {
	if node == nil {
		return nil
	}
	return node.left
}

func (node *AVLnode) GetRight() *AVLnode {
	if node == nil {
		return nil
	}
	return node.right
}

func (node *AVLnode) GetHeight() int{
	if node == nil {
		return 0
	}
	return node.height
}

func BinarySearch(arr []interface{}, data int) int {
	left := 0
	right := len(arr) - 1
	if arr[left].(int) == data {
		return left
	}
	if arr[right].(int) == data {
		return right
	}
	for left < right {
		mid := (left + right) / 2
		if arr[mid].(int) > data {
			right = mid - 1
		} else if arr[mid].(int) < data {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func (node *AVLnode) Middle() []interface{} {
	arr := make([]interface{}, 0)
	return middle(node, arr)
}

func middle(node *AVLnode, arr []interface{}) []interface{} {
	if node != nil {
		arr = middle(node.left, arr)
		arr = append(arr, node.data)
		arr = middle(node.right, arr)
	}
	return arr
}

func (node *AVLnode) Prologue() []interface{}{
	arr := make([]interface{}, 0)
	return prologue(node, arr)
}

func prologue(node *AVLnode, arr []interface{}) []interface{}{
	if node != nil {
		arr = append(arr, node.data)
		arr = prologue(node.left, arr)
		arr = prologue(node.right, arr)
	}
	return arr
}

func (node *AVLnode) Post() []interface{}{
	arr := make([]interface{}, 0)
	return post(node, arr)
}

func post(node *AVLnode, arr []interface{}) []interface{}{
	if node != nil {
		arr = post(node.left, arr)
		arr = post(node.right, arr)
		arr = append(arr, node.data)
	}
	return arr
}


func (node *AVLnode) String() string {
	var buffer bytes.Buffer
	node.BstString(&buffer)
	return buffer.String()
}

func (node *AVLnode) BstString(buffer *bytes.Buffer) {
	if node == nil {
		//buffer.WriteString(tree.Depthstring(depth)+"nil\n")
		return 
	}
	node.left.BstString(buffer)
	buffer.WriteString(node.Depthstring(node.height-1) + strconv.Itoa(node.data.(int)) + "\n")
	node.right.BstString(buffer)
}

func (node *AVLnode) Depthstring(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("-")
	}
	return buffer.String()
}

func main() {
	avl,_ := NewAVLTree(1)
	avl = avl.Add(2)
	avl = avl.Add(3)
	avl = avl.Add(4)
	avl = avl.Add(5)
	avl = avl.Add(6)
	avl = avl.Add(7)
	
	log.Println(avl.Printf())
	mid := avl.Middle()
	log.Println("========Prologue==========\n", avl.Prologue())
	log.Println("=========Post=============\n", avl.Post())

	fmt.Println(avl)
	var x interface{} = 14

	
	if node,ok := avl.FindData(x); node == nil {
		log.Printf("%v is not exist,%v", x, ok)
	} else {
		log.Printf("find %v is %v,\nnode status\n%v, height:%v", node.data, ok, node, node.height)		
	}
	log.Println("=========Middle===========\n", mid)
	log.Println(BinarySearch(mid, 7))
	avl = avl.SetData(6, 14)
	log.Println("=========Middle===========\n",avl.Middle())
	fmt.Println(avl)
	
}

