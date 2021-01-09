package main

import (
	"fmt"
	"strconv"
)

var myGlobalvariable int

func AddOne(s string){
	myGlobalvariable++
	fmt.Printf("	AddOne %s - Total %d\n",s,myGlobalvariable)
}
func DelOne(s string){
	myGlobalvariable--
	fmt.Printf("	DelOne %s - Total %d\n",s,myGlobalvariable)
}
func main() {
	fmt.Println("coursera Concurrency in Go")
	fmt.Println("week 2")
	fmt.Println("")
	fmt.Println("A race condition is when two or more goroutines have access to the same resource concurrently,")
	fmt.Println(" 	such as a variable or data structure and")
	fmt.Println("	at least one of them executes a write operation to it")
	fmt.Println("	without any regard to the other goroutines")
	fmt.Println("")
	fmt.Println("add the -race flag to uour build run or test commands")
	fmt.Println("	to detect any potential race conditions")
	fmt.Println("")


	fmt.Println("concurrency acces to a global var myGlobalvariable")
	fmt.Println("you could see sometime data race on myGlobalvariable write access with")
	fmt.Println("go run -race week2_race.go")

	for i:=0;i<=10;i++ {
		myGlobalvariable = 0
		go AddOne("Test" + strconv.Itoa(i))
		go DelOne("Test" + strconv.Itoa(i))
	}
}
