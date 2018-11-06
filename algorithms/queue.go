package main

import "fmt"

type Queue struct {
	slice []int
}

// Adds the item to the back of the queue
func (q *Queue) Enqueue(i int) {
	q.slice = append(q.slice, i)
}

// returns the first item in the queue and remove it from the queue (FIFO)
// or panics if there isn't one
func (q *Queue) Dequeue() int {
	first := q.slice[0]
	q.slice = q.slice[1 : len(q.slice)]
	return first
}

func (q *Queue) String() string {
	return fmt.Sprint(q.slice) 
}

func main() {
	var q *Queue = new(Queue)
	q.Enqueue(11)
	q.Enqueue(22)
	q.Enqueue(33)
	fmt.Println(q.String())
	fmt.Println(q.Dequeue())
	fmt.Println(q.String())
	q.Enqueue(44)
	fmt.Println(q.String())
	q.Dequeue()
	fmt.Println(q.String())
}


