package main

// A race condition occurs when two or more threads can access shared data and they try to change it at the same time. Because the thread scheduling algorithm can swap between threads at any time, you don't know the order in which the threads will attempt to access the shared data. Therefore, the result of the change in data is dependent on the thread scheduling algorithm, i.e. both threads are "racing" to access/change the data.

import (
	"fmt"
	"time"
)

// Race condition occurs because

func add(s *int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		*s = *s + 2
		fmt.Println(*s)
	}
}

func multiply(s *int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		*s = *s * 2
		fmt.Println(*s)
	}
}

func main() {
	count := 0
	go add(&count)
	multiply(&count)
}
