package cuncurrencyFunctions

import (
	"fmt"
	"runtime"
	"sync"
)

func CuncurencyInAction() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 1000
	functionsToRun := 2

	const goRoutins = 10000
	var waitGroup1 sync.WaitGroup // create  wait group example.
	// var waitGroup2 sync.WaitGroup                                     - no need for this , i think this was a mistake creating 2 waiting groups.
	waitGroup1.Add(goRoutins * functionsToRun) //  set the amount of wait you need.
	// waitGroup2.Add(goRoutins)                                             - no need for this , i think this was a mistake creating 2 waiting groups.
	var mutex1 sync.Mutex // create a mutex

	go spendy(goRoutins, &counter, &waitGroup1, &mutex1)
	go spengy(goRoutins, &counter, &waitGroup1, &mutex1)

	waitGroup1.Wait()
	// waitGroup2.Wait()                             - no need for this , i think this was a mistake creating 2 waiting groups.
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count after cuncurrency happend:", counter)
}

func spendy(goRoutins int, counter *int, waitGroup *sync.WaitGroup, mutex1 *sync.Mutex) {
	for i := 0; i < goRoutins; i++ {
		mutex1.Lock() // i do not know why i dont need to put * before mutex1 ...
		*counter = *counter - 10
		mutex1.Unlock()
		//fmt.Printf("spendy's counter: %v\n", *counter)
		runtime.Gosched() // time.Sleep(time.Second)
		waitGroup.Done()  // each loop we will decrees the amount of wait groups that the "waitGroup1.Wait()" needs to wait.
	}
}

func spengy(goRoutins int, counter *int, waitGroup *sync.WaitGroup, mutex1 *sync.Mutex) {
	for i := 0; i < goRoutins; i++ {
		mutex1.Lock() // i do not know why i dont need to put * before mutex1 ...
		*counter = *counter + 10
		mutex1.Unlock()
		//fmt.Printf("spengy's counter: %v\n", *counter)
		runtime.Gosched() // time.Sleep(time.Second)
		waitGroup.Done()  // each loop we will decrees the amount of wait groups that the "waitGroup1.Wait()" needs to wait.
	}
}
