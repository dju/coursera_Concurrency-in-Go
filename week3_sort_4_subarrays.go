// https://www.coursera.org/learn/golang-concurrency/home/welcome
// Concurrency in Go
// week 3 : Peer-graded Assignment: Module 3 Activity :
// sorting integers that uses four goroutines to create four sub-arrays and then merge the arrays into a single array

/* Instructions
Write a program to sort an array of integers.
The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)
/*

It I recall correctly I simply divided the array length by 4 and handed 3 slices off,
passing any remainder in the last slice,
sorted each slice then merged them back together by comparing head of each result stream and
popping the lowest into the result array.

*/
/*
 *
 */
func Swap(data []int,index int){
	data[index+1], data[index] = data[index], data[index+1]
}
/*
 *
 */
func BubbleSort(wg *sync.WaitGroup, data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < (len(data) - i); j++ {
			if data[j] < data[j-1] {
				Swap(data, j-1)
			}
		}
	}
	wg.Done()
}
func main() {
	var part int = 4
	var max int = 100
	var dataSlice = make([]int, 0, max)
	var response string
	var numbers []string
	var size, sizepart int
	var wg sync.WaitGroup

	fmt.Println("Enter integers separeted with space")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	response = in.Text()
	numbers = strings.Split(response," ")
	size = len(numbers)
	sizepart = size / part
	if  size < part {
		fmt.Printf("There less than %d integers , ending", part)
		return
	}
	if size > max {
		fmt.Printf("There more than %d integers , ending", max)
		return
	}
	for _, value := range numbers {
		tmp, err := strconv.Atoi(value)
		if err == nil {
			dataSlice = append(dataSlice, tmp)
		} else {
			fmt.Printf(" %s is not an integer , ending", value)
			return
		}
	}
	fmt.Print("Before        :")
	fmt.Println(dataSlice)
	// sort parts without race
	for i := 0; i < (part - 1) ; i++ {
		wg.Add(1)
		go BubbleSort(&wg, dataSlice[(i*sizepart):((i*sizepart)+sizepart)])
	}
	wg.Add(1)
	go BubbleSort(&wg, dataSlice[(sizepart * (part - 1)):size])
	// wait sort go routines
	wg.Wait()
	// show sort parts
	for i := 0; i < (part - 1); i++ {
		fmt.Printf("part %d sorted :", i+1)
		fmt.Println(dataSlice[(i*sizepart):((i*sizepart)+sizepart)])
	}
	fmt.Printf("part %d sorted :", part)
	fmt.Println(dataSlice[(sizepart * (part - 1)):size])
	// show merged parts
	fmt.Print("After merge   :")
	fmt.Println(dataSlice)
	// sort sorted parts
	fmt.Print("After sort    :")
	sort.Ints(dataSlice)
	fmt.Println(dataSlice)
}
