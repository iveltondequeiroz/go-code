package main

import "fmt"

type Stack struct {
	slice []int
}

// adds the item to the top of the Stack (LIFO)
func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}

// returns the top item from the stack
func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

// returns the top item from the stack and REMOVE it
func (s *Stack) Pop() int {
	if s.Size()==0 { return -1 }
	top := s.Peek()
	s.slice = s.slice[0:len(s.slice)-1]
	return top
}

// returns the size of the stack
func (s *Stack) Size() int {
	return len(s.slice)
}

// returns stack's string representation
func (s *Stack) String() string {
	return fmt.Sprint(s.slice) 
}

func main() {
	var s *Stack = new(Stack)
	s.Push(55)
	s.Push(44)
	s.Push(33)
	s.Push(22)
	s.Push(11)
	fmt.Println(s.String())
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Push(99)
	fmt.Println(s.Peek())
	fmt.Println(s.String())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.String())
	
	

}
