//A race condition occurs when two or more thread access the same shared resource at the same 
//time and the outcome of the program is due to the sequence or timing of execution and could be bugged

package main

import (
	"fmt"
	"time"
)

var count int

func myFunc() {
	count++
	fmt.Print(count)
}

//the count variable is shared between the two goroutine and this might causes a race condition
//to check it you can run:
//go run -race race.go
func main() {
	go myFunc()
	go myFunc()
	time.Sleep(1000 * time.Millisecond)
}
