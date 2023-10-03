package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData []string = []string{"i1", "i2", "i3", "i4", "i5"}
var wg = sync.WaitGroup{}
var results []string // a shared result that is used by multiple go-routines
var mutex = sync.Mutex{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		//dbCall(i) // this runs the function synchronously the program is waiting for the function to get finished
		// to run the function in a new goroutine we must add the go key word
		wg.Add(1)
		go dbCall(i) // ! the program will not wait for it to  be finished without using sync groups

	}

	//time.Sleep(time.Second) // stop the program for 1000
	wg.Wait()
	fmt.Printf("Total execuation time is %v\n", time.Since(t0))
	fmt.Printf("the results slice %v\n", results)
}

// simulate a DB call
func dbCall(i int) {
	var delay float32 = 1999
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Data fetched from db: ", dbData[i])

	//Locking the mutex by the current goroutine
	mutex.Lock()
	results = append(results, dbData[i])
	mutex.Unlock() // unlock the var so it can b accessed by other goroutines
	wg.Done()
}
