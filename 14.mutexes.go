// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type safeCounter struct {
// 	counts map[string]int
// 	mu     *sync.Mutex
// }

// func (sc safeCounter) inc(key string) {
// 	sc.mu.Lock()
// 	defer sc.mu.Unlock()
// 	sc.counts[key]++
// }

// func (sc safeCounter) val(key string) int {
// 	sc.mu.Lock()
// 	defer sc.mu.Unlock()
// 	return sc.counts[key]
// }

// // don't touch below this line

// func (sc safeCounter) slowIncrement(key string) {
// 	tempCounter := sc.counts[key]
// 	time.Sleep(time.Microsecond)
// 	tempCounter++
// 	sc.counts[key] = tempCounter
// }

// func (sc safeCounter) slowVal(key string) int {
// 	time.Sleep(time.Microsecond)
// 	return sc.counts[key]
// }

// func main() {
// 	m := map[int]int{}

// 	mu := &sync.Mutex{}

// 	go writeLoop(m, mu)
// 	go readLoop(m, mu)

// 	// stop program from exiting, must be killed
// 	block := make(chan struct{})
// 	<-block
// }

// func writeLoop(m map[int]int, mu *sync.Mutex) {
// 	for {
// 		for i := 0; i < 100; i++ {
// 			mu.Lock()
// 			m[i] = i
// 			mu.Unlock()
// 		}
// 	}
// }

// func readLoop(m map[int]int, mu *sync.Mutex) {
// 	for {
// 		mu.Lock()
// 		for k, v := range m {
// 			fmt.Println(k, "-", v)
// 		}
// 		mu.Unlock()
// 	}
// }

// //

// type safeCounterRW struct {
// 	counts map[string]int
// 	mu     *sync.RWMutex
// }

// func (sc safeCounterRW) incRW(key string) {
// 	sc.mu.Lock()
// 	defer sc.mu.Unlock()
// 	sc.slowIncrement(key)
// }

// func (sc safeCounterRW) valRW(key string) int {
// 	sc.mu.RLock()
// 	defer sc.mu.RUnlock()
// 	return sc.counts[key]
// }

// difference between Mutex and MutexRW is that MutexRW allows concurrent reading operation by using RLock() and RUnlock()
// don't touch below this line

// func (sc safeCounterRW) slowIncrement(key string) {
// 	tempCounter := sc.counts[key]
// 	time.Sleep(time.Microsecond)
// 	tempCounter++
// 	sc.counts[key] = tempCounter
// }
