package main

import (
	"log"
	"sync"
)

type Queue struct {
	top  *Node
	last *Node
	size int
	sync.Mutex
}

type Node struct {
	next *Node
	data []byte
}

func New() *Queue {
	return new(Queue)
}

func (queue *Queue) Put(data []byte) {
	queue.Lock()
	defer queue.Unlock()
	node := &Node{
		data: data,
	}
	if queue.last == nil {
		queue.top = node
	} else {
		queue.last.next = node
	}
	queue.last = node
	queue.size++
}

func (queue *Queue) Get() (data []byte) {
	queue.Lock()
	defer queue.Unlock()
	if queue.top == nil {
		return nil
	}
	data = queue.top.data
	queue.top = queue.top.next
	queue.size--
	return data
}

func (queue *Queue) Size() int {
	queue.Lock()
	defer queue.Unlock()
	return queue.size
}

func main() {
	queue := New()
	queue.Put([]byte("abc"))
	queue.Put([]byte("edf"))
	queue.Put([]byte("yhv"))
	queue.Put([]byte("oiu"))
	for queue.size != 0 {
		log.Printf("queue size:%v, top data:%v", queue.Size(), string(queue.Get()))
	}
}
