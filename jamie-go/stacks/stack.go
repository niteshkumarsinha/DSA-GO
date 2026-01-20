package main


import "fmt"


// Stack represent a stack that hold a slice
type Stack struct {
	items []int
}


// Push
func (s *Stack) Push(i int){
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() int {
	length := len(s.items)
	itemToRemove := s.items[length - 1]
	s.items = s.items[:length - 1]
	return itemToRemove
}

func main(){
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
