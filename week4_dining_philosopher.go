// https://www.coursera.org/learn/golang-concurrency/home/welcome
// Concurrency in Go
// week 4 : Peer-graded Assignment: Module 4 Activity :
// the dining philosopher’s problem

/* Instructions
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks)
	it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks)
	it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
type ChopS struct {
	sync.Mutex
}
type Philo struct {
	number          int
	leftCS, rightCS *ChopS
}
var once sync.Once
var CSticks []*ChopS
var philos []*Philo
var ch chan int
var maxPhilosophers int = 5
var maxChopsticks = 5
var maxDinner int = 3
var maxEater int = 2
/*
 *
 */
func initStructs() {
	fmt.Println("Init")
}
/*
 *
 */
func (p Philo) eat(wg2 *sync.WaitGroup, ch chan int) {
	for i := 0; i < maxDinner; i++ {
		<-ch
		p.leftCS.Lock()
		p.rightCS.Lock()
		if i == maxDinner - 1 {
			fmt.Println("finishing eating : ", p.number)
			wg2.Done()
		} else {
			fmt.Println("starting to eat  : ", p.number)
		}
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		ch <- p.number
	}
}
/*
 *
 */
func host(wg *sync.WaitGroup, ch chan int) {
	once.Do(initStructs)
	var wg2 sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	// fmt.Println("seed : ", time.Now().UnixNano())
	// some random
	for i:=0; i < maxEater; i++ {
		tmp := rand.Intn(maxPhilosophers) + 1
		// fmt.Println("ch :", tmp)
		ch <- tmp
	}
	<- ch
	CSticks = make([]*ChopS, maxChopsticks)
	philos := make([]*Philo, maxPhilosophers)
	for i := 0; i < maxChopsticks; i++ {
		CSticks[i] = new(ChopS)
	}
	for i := 0; i < maxPhilosophers; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}
	for i := 0; i < maxPhilosophers; i++ {
		wg2.Add(1)
		go philos[i].eat(&wg2, ch)
	}
	wg2.Wait()
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	ch = make(chan int, maxEater )
	wg.Add(1)
	go host(&wg, ch)
	wg.Wait()
}
