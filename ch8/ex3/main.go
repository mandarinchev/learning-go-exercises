package main

import "fmt"

type LinkedList[T comparable] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
}

type LinkedListNode[T comparable] struct {
	val  T
	next *LinkedListNode[T]
}

func (l *LinkedList[T]) Add(t T) {
	newNode := &LinkedListNode[T]{
		val: t,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = l.head
		return
	}
	l.tail.next = newNode
	l.tail = l.tail.next
}

func (l *LinkedList[T]) Insert(t T, i int) {
	newNode := &LinkedListNode[T]{
		val: t,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	if i <= 0 {
		newNode.next = l.head
		l.head = newNode
		return
	}
	curNode := l.head
	for j := 1; j < i; j++ {
		if curNode.next == nil {
			curNode.next = newNode
			l.tail = curNode.next
			return
		}
		curNode = curNode.next
	}
	newNode.next = curNode.next
	curNode.next = newNode
	if l.tail == curNode {
		l.tail = newNode
	}
}

func (l *LinkedList[T]) Index(t T) int {
	if l == nil {
		return -1
	}
	i := 0
	for curNode := l.head; curNode != nil; curNode = curNode.next {
		if curNode.val == t {
			return i
		}
		i++
	}
	return -1
}

func main() {
	intList := &LinkedList[int]{}
	intList.Add(1)
	intList.Add(2)
	intList.Add(4)
	intList.Insert(3, 2)
	fmt.Println(intList.Index(3))
	fmt.Println(intList.Index(100))

	for curNode := intList.head; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.val)
	}

	stringList := &LinkedList[string]{}
	stringList.Add("One")
	stringList.Add("Two")
	stringList.Add("Three")
	stringList.Add("Five")
	stringList.Insert("Four", 3)
	fmt.Println(stringList.Index("Four"))
	fmt.Println(stringList.Index("Ten"))

	for curNode := stringList.head; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.val)
	}
}
