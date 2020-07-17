package main

import (
	"log"
)

type DataNode interface{}

type ListNode struct {
	Data DataNode
	Prev *ListNode
	Next *ListNode
}

type List struct {
	Size uint
	Head *ListNode
	Tail *ListNode
}

func (this *List) ListInit() {
	this.Size = 0
	this.Head = nil
	this.Tail = nil
}

func (this *List) Append(new *ListNode) bool {
	if new == nil {
		return false
	}
	if this.Size == 0 {
		this.Head = new
		this.Tail = new
		new.Next = nil
		new.Prev = nil
	} else {
		new.Prev = this.Tail
		new.Next = nil
		this.Tail.Next = new
		this.Tail = new
	}
	this.Size++
	return true
}

func (this *List) Get(index uint) *ListNode {
	if this.Size == 0 || index > this.Size-1 {
		return nil
	}

	if index == 0 {
		return this.Head
	}

	node := this.Head
	var i uint
	for i = 1; i < index; i++ {
		node = node.Next
	}
	return node
}

func (this *List) Insert(index uint, new *ListNode) bool {
	if index > this.Size || new == nil {
		return false
	}

	if index == this.Size {
		return this.Append(new)
	}

	if index == 0 {
		new.Next = this.Head
		this.Head = new
		this.Head.Prev = nil
		this.Size++
		return true
	}

	getNode := this.Get(index)
	new.Prev = getNode.Prev
	new.Next = getNode
	getNode.Prev.Next = new
	getNode.Prev = new

	this.Size++
	return true
}

func (this *List) Delete(index uint) bool {
	if index > this.Size-1 || this.Size == 0 {
		return false
	}

	if index == 0 {
		if this.Size == 1 {
			this.Head = nil
			this.Tail = nil
		} else {
			this.Head.Next.Prev = nil
			this.Head = this.Head.Next
		}
		this.Size--
		return true
	}

	if index == this.Size-1 {
		this.Tail.Prev.Next = nil
		this.Tail = this.Tail.Prev
		this.Size--
		return true
	}

	getnode := this.Get(index)
	getnode.Prev.Next = getnode.Next
	getnode.Next.Prev = getnode.Prev
	this.Size--
	return true
}

func (this *List) Display() {
	if this == nil || this.Size == 0 {
		log.Printf("this list is null or empty")
		return
	}

	log.Printf("list size is %d\n", this.Size)
	log.Printf("Display...\n")
	node := this.Head
	for node != nil {
		log.Printf("data is %v\n", node.Data)
		node = node.Next
	}

	log.Printf("Reverse Display...\n")

	tailnode := this.Tail
	for tailnode != nil {
		log.Printf("data is %v\n", tailnode.Data)
		tailnode = tailnode.Prev
	}

}

func main() {
	list := new(List)

	list.ListInit()
	node := make([]ListNode, 10)
	for i := 0; i < 10; i++ {
		node[i].Data = i
		list.Append(&node[i])
	}
	list.Delete(3)
	list.Delete(5)

	newnode := &ListNode{
		Data: "hello",
	}
	list.Insert(4, newnode)
	list.Display()
	log.Printf("%v", (list.Get(4)).Data)

}
