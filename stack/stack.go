package main

import (
	"fmt"
)

var stack []int
var top = -1

func main() {
	fmt.Println("Hello, playground")
	myStack(5)
}

func myStack(size int) {
	stack = make([]int, size)
	fmt.Println(isEmpty())
	push(1)
	push(2)
	push(4)
	push(5)
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(peek())
}

func isEmpty() bool {
	return top == -1
}

func size() int {
	return 0
}

func push(item int) {
	if isFull() {
		fmt.Println("Stack is full!!")
	} else {
		top++
		stack[top] = item
	}
}

func pop() int {
	if isEmpty() {
		fmt.Println("Stack is empty !!")
		return -1
	} else {
		res := stack[top]
		top--
		return res
	}
}
func peek() int {
	if isEmpty() {
		fmt.Println("Stack is empty")
		return -1
	} else {
		return stack[top]
	}
}

func isFull() bool {
	return top+1 == len(stack)
}
