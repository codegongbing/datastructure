package single

import (
	"errors"
	"fmt"
)

type LinkedLister interface {
}

type Node struct {
	Data interface{}
	Next *Node
}

func NewHead() *Node {
	return new(Node)
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data: data,
	}
}

func (head *Node) Append(data interface{}) error {
	if head == nil {
		return errors.New("链表不存在")
	}
	newNode := NewNode(data)
	lastNode := head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = newNode
	return nil
}

func (head *Node) Add(data interface{}) error {
	if head == nil {
		return errors.New("链表不存在")
	}
	newNode := NewNode(data)
	tempNode := head
	newNode.Next = tempNode.Next
	tempNode.Next = newNode
	return nil
}

func (head *Node) Insert(index int, data interface{}) error {
	if head == nil {
		return errors.New("链表不存在")
	}
	newNode := NewNode(data)
	tempNode := head
	count := 0
	for tempNode.Next != nil {
		count++
		if count == index {
			newNode.Next = tempNode.Next
			tempNode.Next = newNode
		}
		tempNode = tempNode.Next
	}
	return nil
}

func (head *Node) RemoveHead() error {
	if head == nil {
		return errors.New("链表不存在")
	}
	tempNode := head
	tempNode.Next = tempNode.Next.Next
	return nil
}

func (head *Node) RemoveTail() error {
	if head == nil {
		return errors.New("链表不存在")
	}
	tempNode := head
	for tempNode.Next != nil {
		if tempNode.Next.Next == nil {
			tempNode.Next = nil
			return nil
		}
		tempNode = tempNode.Next
	}
	return nil
}

func (head *Node) Remove(index int) error {
	if head == nil {
		return errors.New("链表不存在")
	}
	tempNode := head
	count := 0
	for tempNode.Next != nil {
		count++
		if count == index {
			tempNode.Next = tempNode.Next.Next
		}
		tempNode = tempNode.Next
	}
	return nil
}

func (head *Node) PrintList() error {
	if head.Next == nil {
		return errors.New("链表为空")
	}
	tempNode := head.Next
	info := "["
	for tempNode != nil {
		info += fmt.Sprintf("{Data:%v}", tempNode.Data)
		tempNode = tempNode.Next
	}
	info += "]"
	fmt.Println(info)
	return nil
}
