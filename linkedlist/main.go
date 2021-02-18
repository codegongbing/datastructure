package main

import "datastructure/linkedlist/single"

func main() {
	head := single.NewHead()
	for i := 0; i < 10; i++ {
		head.Append(i + 1)
	}
	head.Insert(1, 100)
	head.Insert(2, 200)
	head.RemoveHead()
	head.RemoveTail()
	head.Remove(1)
	head.PrintList()
}
