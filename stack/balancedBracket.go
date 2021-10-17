package stacks

import "fmt"

var mystack []string
var topStack = -1

//[()]

func IsBalanced(str string) bool {
	mystack = make([]string, len(str))
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "[" || string(str[i]) == "(" || string(str[i]) == "{" {
			pushToStack(string(str[i]))
		} else if string(str[i]) == "]" || string(str[i]) == ")" || string(str[i]) == "}" {
			str := topElement()
			if (string(str[i]) == "]" && str == "[") || string(str[i]) == ")" && str == "(" || string(str[i]) == "}" && str == "{" {
				popFromStack()
			} else {
				return false
			}
		}

	}

	return false
}

func isStackEmpty() bool {
	return topStack == -1
}

func stackSize() int {
	return 0
}

func pushToStack(item string) {
	if IsFull() {
		fmt.Println("Stack is full!!")
	} else {
		topStack++
		mystack[topStack] = item
	}
}

func popFromStack() string {
	if IsEmpty() {
		fmt.Println("Stack is empty !!")
		return ""
	} else {
		res := mystack[topStack]
		topStack--
		return res
	}
}
func topElement() string {
	if IsEmpty() {
		fmt.Println("Stack is empty")
		return ""
	} else {
		return mystack[topStack]
	}
}

func isStackFull() bool {
	return topStack+1 == len(stack)
}
