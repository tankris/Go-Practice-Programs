package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type linkedList struct {
	head *Node
}

func newNode(data int) *Node {
	return &Node{data: data}
}

func (l *linkedList) pushBeg(data int) {
	node := newNode(data)

	tempNode := l.head
	l.head = node
	node.next = tempNode
}

func (l *linkedList) delBeg() {
	if l.head != nil {
		l.head = l.head.next
	} else {
		fmt.Println("Error: List is already Empty")
	}
}

func (l *linkedList) pushEnd(data int) {
	node := newNode(data)

	if l.head == nil {
		l.head = node
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = node
	}
}

func (l *linkedList) delEnd() {
	if l.head == nil {
		fmt.Println("Error: List is already Empty")
	} else {
		curr := l.head
		var prev *Node
		for curr.next != nil {
			prev = curr
			curr = curr.next
		}
		prev.next = nil
	}
}

func (l *linkedList) Traverse() {

	if l.head == nil {
		fmt.Println("Empty LL")
		return
	}

	curr := l.head
	for ; curr != nil; curr = curr.next {
		fmt.Println(curr.data)
	}
}

func (l *linkedList) Reverse() {
	var prev *Node
	forward := l.head

	for forward.next != nil {
		forward = forward.next
		l.head.next = prev
		prev = l.head
		l.head = forward
	}

	l.head.next = prev
}

func main() {
	ll := &linkedList{}

	ll.pushBeg(1)
	ll.pushBeg(2)
	ll.pushBeg(3)

	ll.Traverse()

	ll.Reverse()

	ll.Traverse()
}
